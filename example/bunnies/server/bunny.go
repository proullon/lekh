package main

type Bunny struct {
    name string
    id   int
}

func (b *Bunny) String() string {
    return b.name
}

func (b *Bunny) Die() {

}

func (b *Bunny) ID() int {
    return b.id
}
