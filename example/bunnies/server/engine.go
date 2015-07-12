package main

import (
    "github.com/proullon/lekh"
    "log"
)

const (
    _   = iota
    Event_SpawnBunny
)

type Engine struct {
    lekh.BaseEngine
}

func (engine *Engine) HandleEvent(event lekh.Event) {

    switch {
    case event.EventID == Event_SpawnBunny:
        bunny := &Bunny{name: "Tic"}
        em := engine.Entities()
        em.Add(bunny)
        break
    default:
        log.Printf("HandleEvent: Unknown event : %s\n", event)
    }
}
