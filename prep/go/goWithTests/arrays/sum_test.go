package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum numbers in an array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want: '%d' - got:'%d'", want, got)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: '%v' - got:'%v'", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{5, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: '%v' - got:'%v'", want, got)
	}
}
