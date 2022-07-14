package algorithms

// var intBig = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

func AVeryBigSum(ar []int64) int64 {
	var counter int64 = 0
	for i := 0; i < len(ar); i++ {
		counter += ar[i]
	}
	return counter
}
