//Given a single array of real values, each of which represents the stock value of a company
//after an arbitrary period of time, find the best buy price and its corresponding best sell
//price (buy low, sell high)//
//To illustrate with an example, let's take the stock ticker of Company Z
//
//55.39, 109.23, 48.29, 81.59, 105.53, 94.45, 12.24
//Important to note is the fact that the array is "sorted" temporally -
//i.e. as time passes, values are appended to the right end of the array.
//Thus, our buy value will be (has to be) to the left of our sell value
//
//in the above example, the ideal solution is to buy at 48.29 and sell at 105.53


package main
import "fmt"

func main() {
	var input = []float32{55.39, 109.23, 48.29, 81.59, 105.53, 94.45, 12.24}

	var bestBuy float32
	var bestSell float32
	var maxDiff  float32 = 0.0

	for i := 0; i <= (len(input)-1); i++ {
		var lowestPreviousValue float32 = getLowestPrevious(input,i)
		var maxDifferencePerIteration float32 = input[i] - lowestPreviousValue

		if (maxDifferencePerIteration > maxDiff) {
			bestBuy  = lowestPreviousValue
			bestSell = input[i]
		}
	}

	fmt.Println("Buy at:", bestBuy)
	fmt.Println("Sell at:", bestSell)
}


func getLowestPrevious(inputArray []float32, i int) float32 {
	sliceToConsider := inputArray[0:i]
	min := inputArray[0]

	for i := 0; i < len(sliceToConsider); i++ {
		if (min > sliceToConsider[i]) {
			min = sliceToConsider[i]
		}
	}
	return min
}