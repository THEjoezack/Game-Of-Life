package lib

import "testing"

func TestNewGridSize(t *testing.T) {
	cols, rows := 11, 35
	g := NewGrid(cols, rows)
	if len(g.Cells) != cols || len(g.Cells[0]) != rows {
		t.Errorf("NewGrid created grid of incorrect size")
	}
}

func TestNewGridInitialized(t *testing.T) {
	cols, rows := 14, 5
	g := NewGrid(cols, rows)
	if g.Get(1, 1) != 0 {
		t.Errorf("NewGrid should be initialized to 0")
	}
}

func TestGetInvalidIndex(t *testing.T) {
	g := new(Grid)
	if g.Get(-1, 13) != 0 {
		t.Errorf("Invalid index should return 0")
	}
}

func TestGetValidIndex(t *testing.T) {
	g := NewGrid(10, 10)
	x, y := 1, 3
	var expected byte = 2

	g.Cells[x][y] = expected
	if expected != g.Cells[x][y] {
		t.Errorf("Get should retrieve expected index")
	}
}

func TestSetValidIndex(t *testing.T) {
	g := NewGrid(10, 10)
	x, y := 1, 3
	var expected byte = 7

	g.Set(x, y, expected)
	if g.Get(x, y) != expected {
		t.Errorf("Set should set correct index")
	}
}

func TestSetInvalidIndex(t *testing.T) {
	g := NewGrid(10, 10)
	x, y := -1, 3
	var expected byte = 7

	err := g.Set(x, y, expected)
	if err == nil {
		t.Errorf("Should have returned an error")
	}
}

func TestIsInvalidRange(t *testing.T) {
	g := NewGrid(10, 10)

	if g.isInvalidRange(0, 0) {
		t.Errorf("Invalid range")
	}

	if !g.isInvalidRange(10, 2) {
		t.Errorf("Invalid range")
	}

	if !g.isInvalidRange(0, -1) {
		t.Errorf("Invalid range")
	}
}

func TestLiveCount(t *testing.T) {
	g := NewGrid(10, 10)
	g.Set(0, 0, 1)
	g.Set(1, 1, 1)

	if g.liveCount(0, 0) != 1 {
		t.Errorf("Incorrect neighbor count")
	}

	if g.liveCount(1, 0) != 2 {
		t.Errorf("Incorrect neighbor count")
	}
}
