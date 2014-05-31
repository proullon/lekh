package lekh

type Engine interface {
    HandleEvent(event Event)
    TimeUpdate()
    Init(input chan Event, worldSizeX int, worldSizeY int)
    Entities() EntityManager
    Terrain() *World
    RemoveEntity(entity Entity)
}

type BaseEngine struct {
    entities      EntityManager
    terrain       World
    input         chan Event
    entityRemover chan Entity

    idIndex int
}

func (e *BaseEngine) Terrain() *World {
    return &e.terrain
}

func (e *BaseEngine) Entities() EntityManager {
    return e.entities
}

func (e *BaseEngine) Init(input chan Event, worldSizeX int, worldSizeY int) {
    e.input = input
    e.entityRemover = make(chan Entity)
    e.terrain.Init(worldSizeX, worldSizeY)
    e.entities = NewEntityManager()
}

func (e *BaseEngine) RemoveEntity(entity Entity) {
    e.entityRemover <- entity
}
