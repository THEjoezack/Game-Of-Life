package lib

import "errors"

type Grid struct {
	Cells [][]byte
}

func NewGrid(cols, rows int) Grid {
	g := new(Grid)
	g.Cells = make([][]byte, cols)
	for k, _ := range g.Cells {
		g.Cells[k] = make([]byte, rows)
	}
	return *g
}

func (g *Grid) Get(x, y int) byte {
	if g.isInvalidRange(x, y) {
		return 0
	}
	return g.Cells[x][y]
}

func (g *Grid) Set(x, y int, value byte) error {
	if g.isInvalidRange(x, y) {
		return errors.New("Unable to set invalid index")
	}
	g.Cells[x][y] = value
	return nil
}

func (g *Grid) isInvalidRange(x, y int) bool {
	return x >= len(g.Cells) || x < 0 || y >= len(g.Cells[0]) || y < 0
}

func (g *Grid) liveCount(x, y int) byte {
	return g.Get(x-1, y-1) + g.Get(x, y-1) + g.Get(x+1, y-1) +
		g.Get(x-1, y) + g.Get(x+1, y) +
		g.Get(x-1, y+1) + g.Get(x, y+1) + g.Get(x+1, y+1)
}
