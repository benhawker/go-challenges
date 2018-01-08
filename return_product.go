// Given an array of n integers where n > 1, nums,
// return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
// Solve it without division and in O(n).
//
// For example, given [1,2,3,4], return [24,12,8,6].
// Follow up:
// Could you solve it with constant space complexity?
// (Note: The output array does not count as extra space for the purpose of space complexity analysis.)

package main
import "fmt"

func main() {
	array := [4]int{1, 2, 3, 4}
	var productArray [4]int

	for i := 0; i < len(array); i++ {
		product := calculateProduct(array, i)
		productArray[i] = product
	}
	
	fmt.Println(productArray)
}


func calculateProduct(array [4]int, i int) int {
  var sum int = 1 

  for i := 0; i < len(array); i++ {
    sum *= array[i]
  }

  return sum / array[i]
}