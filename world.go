package lekh

type Place struct {
}

type World struct {
    terrain [][]Place
}

func (w *World) Init(x int, y int) {
    w.terrain = make([][]Place, x)
    for i := range w.terrain {
        w.terrain[i] = make([]Place, y)
    }
}
