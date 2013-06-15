package lib

import "testing"

// TODO This smells
func TestGetNextBoard(t *testing.T) {
	g := NewGrid(5, 5)
	g.Set(0, 0, 1)
	g.Set(1, 0, 1)
	g.Set(2, 0, 1)

	expected := NewGrid(5, 5)
	expected.Set(1, 0, 1)
	expected.Set(1, 1, 1)

	next := GetNextGrid(&g)

	for y, cells := range expected.Cells {
		for x, _ := range cells {
			if expected.Get(x, y) != next.Get(x, y) {
				t.Errorf("Expected x:%d y:%d to match", x, y)
			}
		}
	}
}

// TODO Dislike testing more than one thing
func TestGetNextStatus(t *testing.T) {
	if 1 == getNextStatus(1, 0) || 1 == getNextStatus(1, 1) {
		t.Errorf("Any live cell with fewer than two live neighbours dies, as if caused by under-population.")
	}

	if 0 == getNextStatus(1, 2) || 0 == getNextStatus(1, 3) {
		t.Errorf("Any live cell with two or three live neighbours lives on to the next generation.")
	}

	if 1 == getNextStatus(1, 4) || 1 == getNextStatus(1, 5) || 1 == getNextStatus(1, 6) || 1 == getNextStatus(1, 7) || 1 == getNextStatus(1, 8) {
		t.Errorf("Any live cell with more than three live neighbours dies, as if by overcrowding.")
	}

	if 0 == getNextStatus(0, 3) {
		t.Errorf("Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.")
	}

}
