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
	rows, cols, sleepMs, iterations, round := 50, 20, 250, 25, 1
	g := createRandomGrid(cols, rows)

	for iterations == 0 || round <= iterations {
		drawGame(g, round)
		g, round = lib.GetNextGrid(&g), round+1
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	}
}

func drawGame(g lib.Grid, round int) {
	clearScreen()
	fmt.Printf("Round: %d\n", round)
	renderGrid(&g)
	fmt.Println()
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

// TODO Doesn't work on windows
func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func createRandomGrid(rows, cols int) lib.Grid {
	g := lib.NewGrid(rows, cols)
	rand.Seed(time.Now().UTC().UnixNano())

	for x, cells := range g.Cells {
		for y, _ := range cells {
			if rand.Int()%2 == 0 {
				if err := g.Set(x, y, 1); err != nil {
					panic(err)
				}
			}
		}
	}

	return g
}
