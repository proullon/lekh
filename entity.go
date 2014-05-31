package lekh

type Entity interface {
    Die()
    ID() (id int)
    String() string
}
