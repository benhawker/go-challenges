// Given a 2-D grid of size MxN, make every row and column that has a 0 in it become all 0s

package main
import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

// Pass length and width args to the program.
func main() {
	m := os.Args[1]
	n := os.Args[2]

	// string to int
  width, err := strconv.Atoi(m)
  if err != nil {
    // handle error
    fmt.Println(err)
    os.Exit(2)
  }

  length, err := strconv.Atoi(n)
  if err != nil {
    // handle error
    fmt.Println(err)
    os.Exit(2)
  }

  size := width*length
  grid := make([]int, size)

  // Insert some random ints between 0-9 into the grid.
	for i := 0; i < len(grid); i++ {
		grid[i] = rand.Intn(9)
	}

	duplicatedGrid := make([]int, len(grid))
	copy(duplicatedGrid, grid)

	fmt.Println("Original")
	printGrid(grid, width)
	fmt.Println("--------")

	// Next deal with the zeros
	for i := 0; i < len(grid); i++ {
		if (grid[i] == 0) {

			// Write zeros downwards in a col
			for j := i; j < size; j += length {
				duplicatedGrid[j] = 0
			}

			// Write zeros upwards in a col
			for j := i; j >= 0; j -= length {
				duplicatedGrid[j] = 0
			}

			// Write zeros to the right in a row
			// for j := i; ((j+1) % (width) == 0); j++ {
			for j := i; (j) % 5 != 0; j++ {
				duplicatedGrid[j] = 0
			}

			// Write zeros to the left in a row
			for j := i; (j+1) % 5 != 0; j-- {
				duplicatedGrid[j] = 0
			}
		}
	}
	fmt.Println("--------")
	printGrid(duplicatedGrid, width)
}


func printGrid(grid []int, width int) {
	for i, _ := range grid {
  	fmt.Print(grid[i])
  	if (i != 0) && ((i+1) % width == 0) {
			fmt.Println()
		}
	}
}