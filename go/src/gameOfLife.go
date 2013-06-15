package main

import (
	"fmt"
	"lib"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	rows := 10
	cols := 10
	sleepMs := 500
	iterations := 10

	g := createRandomGrid(cols, rows)
	round := 1

	for iterations == 0 || round <= iterations {
		clearScreen()

		fmt.Printf("Round: %d\n", round)
		renderGrid(&g)
		fmt.Println()

		g = lib.GetNextGrid(&g)
		round += 1

		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	}
}

func renderGrid(g *lib.Grid) {
	for _, cells := range g.Cells {
		for _, v := range cells {
			if v == 1 {
				fmt.Print("O")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func createRandomGrid(rows, cols int) lib.Grid {
	g := lib.NewGrid(rows, cols)
	rand.Seed(time.Now().UTC().UnixNano())
	for y, cells := range g.Cells {
		for x, _ := range cells {
			if rand.Int()%2 == 0 {
				g.Set(x, y, 1)
			}
		}
	}

	return g
}
