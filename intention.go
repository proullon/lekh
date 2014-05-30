package lekh

import (
    "encoding/json"
)

/* A course of action that an entity intends to follow
** Send from Entity to Lekh engine
 */
type Intention struct {
    EventID  int
    EntityID int
    Payload  interface{}
}

func NewIntention(data []byte) (intent Intention) {
    err := json.Unmarshal(data, &intent)

    if err != nil {
        panic(err)
    }

    return
}

func (intent Intention) Json() []byte {
    b, err := json.Marshal(intent)

    if err != nil {
        panic(err)
    }

    return b
}
