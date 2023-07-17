package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	nums := []int{4, 3, 8, 12, 5, 99, 76, 2, 1, -999}
	Sort(nums)
	fmt.Printf("%d\n", nums)

}

func Sort(values []int) {
	// create an empty node for root
	var root *tree
	for _, v := range values {
		//for each item in values, call add
		root = add(root, v)
	}
	//values comes in as a slice, by passing values[:0] to
	//append values we are passing the reference to the memory block of the underlying array without passing any values
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	//if the root is empty
	if t == nil {
		t = new(tree)   //allocate a new tree node
		t.value = value // set value
		return t
	}
	//if root has a value
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// using append() here on a 0 length slice of the original slice
// lets us use the space allocated to the original slice but not have to pass any references to its values over.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)  // pass in left side (everything smaller than root) until we hit an edge
		values = append(values, t.value)       //append the root value
		values = appendValues(values, t.right) // pass in everything larger than root value
	}
	return values
}
