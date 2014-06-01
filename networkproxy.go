package lekh

import (
    "log"
    "net"
    "strconv"
)

type NetworkProxy struct {
    intentionReceiver chan Intention
    eventDispatcher   chan Event
}

func NewNetworkProxy() *NetworkProxy {
    p := NetworkProxy{}

    p.intentionReceiver = make(chan Intention)
    p.eventDispatcher = make(chan Event)
    return &p
}

/* Server side
 */
func (lp *NetworkProxy) DispatchEvent(event Event) {
    go func() {
        lp.eventDispatcher <- event
    }()
}

/* Server side
 */
func (lp *NetworkProxy) ReceiveIntentions(intentDispatcher chan Intention) {
    go func() {
        for {
            select {
            case intent := <-lp.intentionReceiver:
                intentDispatcher <- intent
                break
            }
        }
    }()
}

/* Client side
 */
func (lp *NetworkProxy) DispatchIntent(intent Intention) {
    go func() {
        lp.intentionReceiver <- intent
    }()
}

/* Client side
 */
func (lp *NetworkProxy) ReceiveEvents(eventDispatcher chan Event) {
    go func() {
        for {
            select {
            case event := <-lp.eventDispatcher:
                eventDispatcher <- event
                break
            }
        }
    }()
}

/* Client side
 */
func (lp *NetworkProxy) DialAndServe(addr string) error {

    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return err
    }

    go func() {
        // Read event from TCP
        eventTcp := make(chan Event)
        go func() {
            for {
                data := make([]byte, 512)
                // Let's say everything is going well...
                n, err := conn.Read(data)
                if err != nil {
                    if n != 0 {
                        log.Printf("NetworkProxy: %s\n", err)
                    }
                    close(eventTcp)
                    return
                }

                event := NewEvent(data)
                eventTcp <- event
            }
        }()

        for {
            select {
            // New event from TCP conn. Send it to entity through channel
            case event, ok := <-eventTcp:
                if ok == false {
                    log.Printf("Closing proxy\n")
                    return
                }
                lp.eventDispatcher <- event
                break
            // Entity wants to dispatch an intent. Send it through conn to lekh engine
            case intent := <-lp.intentionReceiver:
                conn.Write(intent.Json())
                break
            }
        }
    }()

    return nil
}

/* Server side
 */
func (lp *NetworkProxy) ListenAndServe(port int) error {
    ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
    if err != nil {
        return err
    }

    go func() {
        for {
            conn, err := ln.Accept()
            if err != nil {
                log.Printf("NetworkProxy: Cannot accept new connection: %s\n", err)
                continue
            }

            log.Printf("NetworkProxy: New client\n")
            go lp.handleConnection(conn)
        }
    }()

    return nil
}

/* Server side
 */
func (lp *NetworkProxy) handleConnection(conn net.Conn) {

    // Read intentions from TCP
    intentionTcp := make(chan Intention)
    go func() {
        for {
            data := make([]byte, 512)
            // Let's say everything is going well...
            n, err := conn.Read(data)
            if err != nil {
                if n != 0 {
                    log.Printf("NetworkProxy: %s\n", err)
                }
                close(intentionTcp)
                return
            }

            intent, err := NewIntention(data[:n])
            if err != nil {
                continue
            }

            intentionTcp <- intent
        }
    }()

    for {
        select {
        // Lekh engine wants to dispatch an event. Send it through conn to entity
        case event := <-lp.eventDispatcher:
            conn.Write(event.Json())
            break
        // New intent from TCP conn. Send it to Lekh through channel
        case intent, ok := <-intentionTcp:
            if ok == false {
                log.Printf("Closing proxy\n")
                return
            }
            lp.intentionReceiver <- intent
            break
        }
    }
}
