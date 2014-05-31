package lekh

import (
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

func (lp *NetworkProxy) DispatchEvent(event Event) {
    go func() {
        lp.eventDispatcher <- event
    }()
}

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

func (lp *LocalProxy) Listen(port int) {
    ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
    if err != nil {
        panic(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            panic(err)
            continue
        }
        go lp.handleConnection(conn)
    }
}

func (lp *LocalProxy) handleConnection(conn net.Conn) {

    // Read intentions from TCP
    intentionTcp := make(chan Intention)
    go func() {
        for {
            data := make([]byte, 512)
            // Let's say everything is going well...
            _, err := conn.Read(data)
            if err != nil {
                panic(err)
            }
            intent := NewIntention(data)
            intentionTcp <- intent
        }
    }()

    for {
        select {
        // Lekh engine want to dispatch an event. Send it through conn
        case event := <-lp.eventDispatcher:
            conn.Write(event.Json())
            break
        // New intent from TCP conn. Send it to Lekh through channel
        case intent := <-intentionTcp:
            lp.intentionReceiver <- intent
            break
        }
    }
}
