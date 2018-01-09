// Find the length of streaks within the following array:
//
// states = [:read, :read, :read, :unread, :unread, :read, :read, :read, :read]
// The array has either :read or :unread, and I want the length of each streak of :read/:unread. For states, this would become:
// streak_lengths = [3, 3, 3, 2, 2, 4, 4, 4, 4]
// The array opens with three :read elements, and we label each of those as being part of a 3 streak, it then has two :unread elements, so they're each labelled with a 2 streak, and then finally we have a streak of four read messages, so they're each labelled with a 4.


package main
import "fmt"

func main() {
	var array = [9]string{"read","read","read","unread","unread","read","read","read","read"}
	var result = [9]int{0,0,0,0,0,0,0,0,0}
	var counter = 1
	var lastIndex = 0

	for i := 0; i < (len(array)-1); i++ {		
		if (array[i] == array[i+1]) {
			counter += 1
		}

		if array[i] != array[i+1] {
			for lastIndex := lastIndex; lastIndex <= i; lastIndex++ {
				result[lastIndex] = counter
			}
			counter = 1
			lastIndex = i+1
		}
	}

	if (array[len(array)-2] == array[len(array)-1]) {
		result[len(array)-1] = counter

		for lastIndex := lastIndex; lastIndex <= (len(array)-1); lastIndex++ {
			result[lastIndex] = counter
		}
	} else {
		result[len(array)-1] = 1
	}

	fmt.Println(result)
}