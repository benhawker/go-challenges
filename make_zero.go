// Given a 2-D grid of size MxN, make every row and column that has a 0 in it become all 0s

package main
import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

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

	// Print the original grid optimised for the terminal.
	for i := 0; i < len(grid); i++ {
		fmt.Print(grid[i])
		if (i != 0) && ((i+1) % width == 0) {
			fmt.Println()
		}
	}

	// Next deal with the zeros

}