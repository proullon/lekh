package lekh

type Proxy interface {
    /* Lekh side */
    DispatchEvent(Event)
    ReceiveIntentions(chan Intention)
}
