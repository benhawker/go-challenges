// For the given array [1,4,3,6,7,5] return the elements that have the peak difference
// vs the previous element
//
// array = [1,4,3,6,7,5]
// Should return 4 and 6.

package main
import (
	"fmt"
	"sort"
)

func main() {
	var array = []int{1,4,3,6,7,5}

	// Diffs
	diffMap := make(map[int][]int)
	
	// fmt.Println(array)

	for i := 1; i < len(array); i++ {
		difference := (array[i] - array[i-1])

		diffMap[difference] = append(diffMap[difference], i)
	} 

 	// To store the keys in slice in sorted order
  var keys []int
  for k := range diffMap {
    keys = append(keys, k)
  }
  
  sort.Ints(keys)

  fmt.Println("The peak difference elements are:")

  var results = diffMap[(keys[(len(keys)-1)])]

  for i := 0; i < (len(results)); i++ {
  	var value = results[i]
  	fmt.Println(array[value])
  }
}