package arrays

func Sum(numbers []int) int {
	var total int

	for _, val := range numbers {
		total += val
	}
	return total
}

// func SumAll(numsToSum ...[]int) []int {
// 	lenOfNums := len(numsToSum)
//
// 	// make will create a slice of ints with an initial length of the
// 	// length of our inputs
// 	// this reads "make me a slice with a size of lenOfNums and initialize
// 	// everything to 0
// 	sums := make([]int, lenOfNums)
//
//   // numbers here will be each individual slice
// 	for i, numbers := range numsToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

func SumAll(numsToSum ...[]int) []int {
	var sums []int

	// numbers here will be each individual slice
	for _, numbers := range numsToSum {
		//append will take in a slice and return a new slice allotted correctly
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var tailSums []int

	for _, numbers := range numsToSum {
		tail := numbers[1:]
		tailSums = append(tailSums, Sum(tail))
	}
	return tailSums
}
