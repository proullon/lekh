package main

import (
    "github.com/proullon/lekh"
    "log"
)

func NewDebugLaw() lekh.Law {
    law := lekh.Law{}

    law.Name = "Debug"
    law.Delay = 5
    law.Targets = nil
    law.Law = func(entity lekh.Entity, w *lekh.World, em lekh.EntityManager) {
        log.Printf("[DEBUG] Entity %d : %s\n", entity.ID(), entity)
    }
    return law
}
