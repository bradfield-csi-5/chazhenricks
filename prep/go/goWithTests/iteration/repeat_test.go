package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats characters", func(t *testing.T) {
		repeated := Repeat("a", 20)
		expected := "aaaaaaaaaaaaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected: '%s' - received: '%s'", expected, repeated)
		}
	})

	t.Run("tests that target is containted in base", func(t *testing.T) {
		got := ContainsAny("chaz", "az")
		want := true

    assertCorrectMessage(t, got, want)

	})
	t.Run("tests that target is not containted in base", func(t *testing.T) {
		got := ContainsAny("chaz", "sx")
		want := false

    assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t - want %t", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}

}
func ExampleRepeat() {
	repeated := Repeat("chaz", 3)
	fmt.Println(repeated)
	//Output: chazchazchaz
}
