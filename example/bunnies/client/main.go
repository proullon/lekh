package main

import (
    "github.com/proullon/lekh"
    "log"
)

func main() {
    log.Printf("Starting bunnies client\n")

    // Create new proxy
    proxy := lekh.NewNetworkProxy()

    err := proxy.DialAndServe("127.0.0.1:8080")
    if err != nil {
        log.Printf("Cannot connect to server: %s\n", err)
        return
    }

    intent := lekh.Intention{}
    intent.EntityID = 1
    intent.EventID = 2
    proxy.DispatchIntent(intent)

    eventReceiver := make(chan lekh.Event)
    proxy.ReceiveEvents(eventReceiver)
    for {
        select {
        case event := <-eventReceiver:
            log.Printf("Event ! %s\n", event.Json())
            break
        }
    }
}
