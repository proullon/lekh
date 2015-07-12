package lekh

import (
    "log"
    "reflect"
    "time"
)

type Lekh struct {
    cycle             time.Duration
    stop              chan bool
    engine            Engine
    actions           []Action
    input             chan Event
    laws              []lawyer
    intentionReceiver chan Intention
    proxies           []Proxy

    idIndex int
}

func (l *Lekh) Init(engine Engine, worldSizeX int, worldSizeY int, cycle time.Duration) {

    l.stop = make(chan bool)
    l.input = make(chan Event)
    l.intentionReceiver = make(chan Intention)
    l.cycle = cycle

    l.engine = engine
    l.engine.Init(l.input, worldSizeX, worldSizeY)
}

func (l *Lekh) Stop() {
    l.stop <- true
}

func (lekh *Lekh) SpawnEntity(e Entity, p Proxy) (entityID int) {
    log.Printf("[LEKH] Spawing entity %s with id %d\n", e, lekh.idIndex)
    // Add entity in engine EntityManager
    em := lekh.engine.Entities()
    em.Add(e)

    // Add proxy in proxy slice
    lekh.AddProxy(p)

    // Get id
    entityID = lekh.idIndex
    lekh.idIndex++

    return
}

func (l *Lekh) AddProxy(p Proxy) {
    l.proxies = append(l.proxies, p)
    p.ReceiveIntentions(l.intentionReceiver)
}

func (l *Lekh) AddLaw(law Law) {
    log.Printf("Adding new law %v\n", law)
    handler := lawyer{CurrentDelay: law.Delay, L: law}
    l.laws = append(l.laws, handler)
}

func (l *Lekh) Bind(i Intention, a Action) {

}

func (l *Lekh) Run() {
    log.Printf("Lekh engine starting...\n")

    // Main lekh loop
    for {
        select {
        case _ = <-l.stop:
            log.Println("Lekh main loop stoped")
            return
        case _ = <-time.After(l.cycle):
            l.timeUpdate()
            break
        case event := <-l.input:
            l.engine.HandleEvent(event)
            for i := range l.proxies {
                l.proxies[i].DispatchEvent(event)
            }
            break
        case intent := <-l.intentionReceiver:
            log.Printf("New intent : %v\n", intent)
            break
        }
    }
}

func (l *Lekh) timeUpdate() {
    // Go through each actions an reduce delay
    // If delay is 0, activate it
    for i := range l.actions {
        l.actions[i].Delay--
        if l.actions[i].Delay == 0 {
            l.actions[i].Actioner(l.actions[i].Target)
        }
    }

    // Check laws
    for i := range l.laws {
        l.laws[i].CurrentDelay--
        if l.laws[i].CurrentDelay == 0 {
            // Reset Delay
            l.laws[i].CurrentDelay = l.laws[i].L.Delay
            // Go Through all entities
            em := l.engine.Entities()
            entities := em.Entities()

            for x := range entities {
                if l.laws[i].L.Targets == nil || reflect.TypeOf(entities[x]) == l.laws[i].L.Targets || reflect.TypeOf(entities[x]) == reflect.PtrTo(l.laws[i].L.Targets) {
                    go l.laws[i].L.Law(entities[x], l.engine.Terrain(), l.engine.Entities())
                }
            }
        }
    }

}
