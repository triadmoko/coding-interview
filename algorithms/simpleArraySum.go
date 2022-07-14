package algorithms

import "fmt"

// sample value
// val := []int{1, 2, 43, 45, 34, 3, 534}
func SimpleArraySum(sum []int) int {
	num := 0
	for _, v := range sum {
		num += v
	}
	fmt.Println("Result from = ", sum, num)
	return num
}
