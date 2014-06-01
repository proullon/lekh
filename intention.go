package lekh

import (
    "encoding/json"
    "log"
)

/* A course of action that an entity intends to follow
** Send from Entity to Lekh engine
 */
type Intention struct {
    EventID  int
    EntityID int
    Payload  interface{}
}

func NewIntention(data []byte) (intent Intention, err error) {
    err = json.Unmarshal(data, &intent)

    if err != nil {
        log.Printf("Intention: Cannot unmarshall %s: %s\n", string(data), err)
        return
    }

    return
}

func (intent Intention) Json() []byte {
    b, err := json.Marshal(intent)
    log.Printf("Intent.JSON : %s\n", string(b))

    if err != nil {
        log.Printf("Intention: Cannot marshall intent: %s\n", err)
        return nil
    }

    return b
}
