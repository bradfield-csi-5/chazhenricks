package arrays

func Sum(numbers []int) int {
  var total int

  for _, val := range numbers {
    total += val
  }
	return total
}
