package main

type rules struct { }

func (r *rules) getNextStatus(g *Grid, x, y int) byte {
        status := g.Get(x, y)
        count := g.liveCount(x, y)
        if status == 0 && count == 3 {
                return 1
        }
        if status == 1 && (count == 2 || count == 3) {
                return 1
        }
        return 0
}

func (r *rules) getNextGrid(g *Grid) Grid {
        next := newGrid(len(g.cells), len(g.cells[0]))
        for x, cells := range g.cells {
                for y, _ := range cells {
                        next.set(x, y, r.getNextStatus(g, x, y))
                }
        }
        return next
}

