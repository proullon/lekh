package lekh

type Actioner interface {
    Action(Entity)
}

type Action struct {
    Delay    int
    Target   Entity
    Actioner func(Entity)
}
