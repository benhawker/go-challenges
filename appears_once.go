// Find the element that occurs only once in a given set of integers
// while all the other numbers occur thrice.
// Example 1:
// Input : 3 3 3 4
// Output: 4
// Example 2:
// Input : 5 5 4 8 4 5 8 9 4 8
// Output: 9
// array = %w(5 5 4 8 4 5 8 9 4 8).map(&:to_i)

package main
import "fmt"

func main() {
	//Input array hardcoded.
	array := [10]int{5,5,4,8,4,5,8,9,4,8}

	// Map with int as key, int as value.
	seen := make(map[int]int)

	for i := 0; i < len(array); i++ {
		var value = array[i]
		if _, ok := seen[value]; ok {
			seen[value] += 1
		} else {
			seen[value] = 1
		}
	}

	for k, v := range seen {
		if v == 1 {
			fmt.Printf("[%s] appears once\n", k)
		}
	}
}