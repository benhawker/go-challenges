// Write a function that compresses a string from aaabbcc => a3b2c2

package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	input := os.Args[1]
	sliceResult := []string{}

	slice := strings.Split(input, "")

	index := 0
	lastValue := 0
	lastIndex := 0

	for i := index; i < (len(slice)-1); i++ {
		if slice[i] != slice[i+1] {
			sliceResult = append(sliceResult, slice[i])
			sliceResult = append(sliceResult, (strconv.Itoa((i+1) - lastIndex)))

			lastValue += i
			lastIndex = i+1
		}
	}

	if slice[len(slice)-2] == slice[len(slice)-1] {
		// It matches
		sliceResult = append(sliceResult, slice[len(slice)-1])
		sliceResult = append(sliceResult, (strconv.Itoa(len(slice) - lastIndex)))

	} else {
		// Does not match It's a single value at the end.
		sliceResult = append(sliceResult, slice[len(slice)-1])
		sliceResult = append(sliceResult, "1")
	}

	fmt.Println(sliceResult)
}

