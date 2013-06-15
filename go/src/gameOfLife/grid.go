package main

type Grid struct {
        cells [][]byte
}

func (g *Grid) Get(x, y int) byte {
        if x >= len(g.cells) || x < 0 || y >= len(g.cells[0]) || y < 0 {
                return 0
        }
        return g.cells[y][x]
}

func (g *Grid) set(x, y int, value byte) {
        if x > len(g.cells) || x < 0 || y > len(g.cells[0]) || y < 0 {
                panic("Invalid index!")
        }
        g.cells[y][x] = value
}

func (g *Grid) liveCount(x, y int) byte {
        return g.Get(x-1, y-1) + g.Get(x, y-1) + g.Get(x+1, y-1) +
                g.Get(x-1, y) + g.Get(x+1, y) +
                g.Get(x-1, y+1) + g.Get(x, y+1) + g.Get(x+1, y+1)
}

func newGrid(x, y int) Grid {
        g := new(Grid)
        g.cells = make([][]byte, y)
        for k, _ := range g.cells {
                g.cells[k] = make([]byte, x)
        }
        return *g
}

