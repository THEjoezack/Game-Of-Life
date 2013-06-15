package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"math/rand"
)

func renderGrid(g *Grid) {
	for _, cells := range g.cells {
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

func createRandomGrid(rows, cols int) Grid {
	g := newGrid(rows,cols)
	rand.Seed(time.Now().UTC().UnixNano())
	for y, cells := range g.cells {
		for x, _ := range cells {
			if rand.Int() % 2 == 0 {
				g.set(x,y,1)
			}
		}
	}

	return g
}

func main() {
	g := createRandomGrid(45,45)
	r := new(rules)
	round := 0

	for {
		clearScreen()
		renderGrid(&g)
		fmt.Println()
		fmt.Printf("Round: %d", round)

		g = r.getNextGrid(&g)
		round += 1
		time.Sleep(100 * time.Millisecond)
		
	}
}
