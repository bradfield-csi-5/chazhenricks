package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		//check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		//add a thing, then check if not empty
		myStackOfInts.Push(32)
		AssertFalse(t, myStackOfInts.IsEmpty())

		//add another thing, pop it back
		myStackOfInts.Push(42)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 42)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 32)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("sting stack", func(t *testing.T) {

		myStackOfStrings := new(Stack[string])

		//check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		//add a thing, then check if not empty
		myStackOfStrings.Push("chaz")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		//add another thing, pop it back
		myStackOfStrings.Push("alex")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "alex")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "chaz")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
