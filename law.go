package lekh

import (
    "reflect"
)

/* Laws usually refers to standards mechanism happening in time
 */
type Law struct {
    Name    string
    Delay   int
    Targets reflect.Type
    Law     func(Entity, *World, EntityManager)
}

/* Interal struct for engine
 */
type lawyer struct {
    CurrentDelay int
    L            Law
}
