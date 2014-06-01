package main

import (
    "github.com/proullon/lekh"
    "log"
    "time"
)

func main() {
    log.Printf("Starting bunnies server\n")

    // Initialise Engine
    l := lekh.Lekh{}
    engine := &Engine{}
    l.Init(engine, 10, 10, time.Millisecond*800)

    // Add engine laws
    l.AddLaw(NewDebugLaw())

    // Add network proxy
    proxy := lekh.NewNetworkProxy()
    proxy.ListenAndServe(8080)
    l.AddProxy(proxy)

    // Run simulation
    l.Run()
}
