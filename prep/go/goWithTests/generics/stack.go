package generics

// non-generic stack implementations of ints and strings
type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

//now using generics

// stack can be of ANY type, but the type must be specified when the stack is created
type Stack[T any] struct {
  //underlying data structure is a slice under the hood
	values []T
}


// Add method to the Stack struct that adds a new value to the _end_ of the slice
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		//need a zero value for whatever T is going to be
		var zero T
		return zero, false
	}

	index := len(s.values) - 1 //need the las index value 
	el := s.values[index] //grab the value at the last index location 
	s.values = s.values[:index] //make the stack equal to the stack-minus-last-element
	return el, true //return last element
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
