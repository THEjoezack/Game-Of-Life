package lib

import "testing"

func TestNewGrid(t *testing.T) {
	cols, rows := 11, 35
	g := NewGrid(cols, rows)
	if len(g.Cells) != Cols || len(g.Cells[0]) != rows {
		t.Errorf("NewGrid created grid of incorrect size")
	}
}
