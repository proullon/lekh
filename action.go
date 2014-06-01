package lekh

type Action struct {
    Delay    int
    Target   Entity
    Actioner func(Entity)
}
