package lib

func GetNextGrid(g *Grid) Grid {
	next := NewGrid(len(g.Cells), len(g.Cells[0]))
	for x, cells := range g.Cells {
		for y, _ := range cells {
			next.Set(x, y, getNextStatus(g, x, y))
		}
	}
	return next
}

func getNextStatus(g *Grid, x, y int) byte {
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
