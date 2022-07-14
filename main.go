package main

import (
	"fmt"

	"github.com/triadmoko/coding-interview/algorithms"
)

func main() {
	// simpel array sum
	val := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	data := algorithms.AVeryBigSum(val)
	fmt.Println(data)
}
