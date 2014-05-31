package lekh

type LocalProxy struct {
    intentionReceiver chan Intention
    eventDispatcher   chan Event
}

func NewLocalProxy() *LocalProxy {
    p := LocalProxy{}

    p.intentionReceiver = make(chan Intention)
    p.eventDispatcher = make(chan Event)
    return &p
}

func (lp *LocalProxy) DispatchEvent(event Event) {
    go func() {
        lp.eventDispatcher <- event
    }()
}

func (lp *LocalProxy) ReceiveIntentions(intentDispatcher chan Intention) {
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

func (lp *LocalProxy) SendIntention(intent Intention) {
    go func() {
        lp.intentionReceiver <- intent
    }()
}

func (lp *LocalProxy) ReceiveEvents(eventReceiver chan Event) {
    go func() {
        for {
            select {
            case event := <-lp.eventDispatcher:
                eventReceiver <- event
                break
            }
        }
    }()
}
