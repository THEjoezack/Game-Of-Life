package main

import (
	"flag"
	"fmt"
	"lib"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var cols = flag.Int("columns", 50, "How many columns")
	var rows = flag.Int("rows", 20, "How many rows")
	var sleepMs = flag.Int("sleepMs", 250, "How many milliseconds to sleep between frames")
	var iterations = flag.Int("iterations", 50, "How many iterations to run, 0 for infinite")
	flag.Parse()

	g := createRandomGrid(*cols, *rows)
	round := 1

	for *iterations == 0 || round <= *iterations {
		renderGame(g, round)
		g, round = lib.GetNextGrid(&g), round+1
		time.Sleep(time.Duration(*sleepMs) * time.Millisecond)
	}
}

func renderGame(g lib.Grid, round int) {
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
	g := lib.NewGrid(cols, rows)
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
