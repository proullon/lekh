package lekh

import (
    "reflect"
)

/* Rule usually refers to standards for activities
 */
type Rule struct {
    Name    string
    Delay   int
    Targets reflect.Type
    Ruler   func(Entity, *World, EntityManager)
}

type ruleHandler struct {
    CurrentDelay int
    R            Rule
}
