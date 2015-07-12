package lekh

import (
    "log"
)

type EntityManager struct {
    entities []Entity

    stoper        chan bool
    entityRemover chan Entity
    entityAdder   chan Entity
    entityGetter  chan chan []Entity
}

func NewEntityManager() (em EntityManager) {

    em.entityRemover = make(chan Entity)
    em.entityAdder = make(chan Entity)
    em.stoper = make(chan bool)
    em.entityGetter = make(chan chan []Entity)

    go em.run()
    return em
}

func (em *EntityManager) Remove(entity Entity) {
    em.entityRemover <- entity
}

func (em *EntityManager) Add(entity Entity) {
    em.entityAdder <- entity
}

func (em *EntityManager) Entities() (entities []Entity) {
    channel := make(chan []Entity)
    em.entityGetter <- channel
    entities = <-channel
    return entities
}

func (em *EntityManager) run() {
    for {
        select {
        case _ = <-em.stoper:
            return
        case e := <-em.entityAdder:
            em.entities = append(em.entities, e)
            break
        case e := <-em.entityRemover:
            em.removeEntity(e)
            break
        case channel := <-em.entityGetter:
            channel <- em.entities
            break
        }
    }
}

func (em *EntityManager) removeEntity(entity Entity) {
    log.Printf("Entities.Remove (len=%d)\n", len(em.entities))

    // Find entity in slice
    for i := range em.entities {
        log.Printf("Removing entity %d\n", em.entities[i].ID())
        if em.entities[i].ID() == entity.ID() {
            em.entities = append(em.entities[:i], em.entities[i+1:]...)
            return
        }
    }
}
