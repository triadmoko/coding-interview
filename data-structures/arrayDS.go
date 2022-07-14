package datastructures

import "fmt"

func ReverseArray(a []int) []int {
	newNumbers := make([]int, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, a[i])
	}
	fmt.Println("Result from = ", a ,newNumbers)
	return newNumbers
}
