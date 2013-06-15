package lib

type Grid struct {
	Cells [][]byte
}

func NewGrid(x, y int) Grid {
	g := new(Grid)
	g.Cells = make([][]byte, y)
	for k, _ := range g.Cells {
		g.Cells[k] = make([]byte, x)
	}
	return *g
}

func (g *Grid) Get(x, y int) byte {
	if x >= len(g.Cells) || x < 0 || y >= len(g.Cells[0]) || y < 0 {
		return 0
	}
	return g.Cells[y][x]
}

func (g *Grid) Set(x, y int, value byte) {
	if x > len(g.Cells) || x < 0 || y > len(g.Cells[0]) || y < 0 {
		panic("Invalid index!")
	}
	g.Cells[y][x] = value
}

func (g *Grid) liveCount(x, y int) byte {
	return g.Get(x-1, y-1) + g.Get(x, y-1) + g.Get(x+1, y-1) +
		g.Get(x-1, y) + g.Get(x+1, y) +
		g.Get(x-1, y+1) + g.Get(x, y+1) + g.Get(x+1, y+1)
}
