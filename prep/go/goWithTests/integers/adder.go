package integers

import "fmt"

func main() {

	result := Add(2, 2)
  fmt.Printf("Result: %d", result)
}

// Takes two integers and returns a sum of them
func Add(a, b int) int {
  return a + b
}
