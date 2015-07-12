package lekh

import (
    "encoding/json"
)

/* Message indicating that something has happened
** Send from Lekh engine to Entity
 */
type Event struct {
    EventID  int
    EntityID int
    Payload  interface{}
}

func NewEvent(data []byte) (event Event) {
    err := json.Unmarshal(data, &event)

    if err != nil {
        panic(err)
    }

    return
}

func (event Event) Json() []byte {
    b, err := json.Marshal(event)

    if err != nil {
        panic(err)
    }

    return b
}
