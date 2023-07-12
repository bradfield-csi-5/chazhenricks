package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "chaz", "chaz")
		AssertNotEqual(t, "alex", "chaz")
	})

  // AssertEqual(t, 1, "1")
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v - want %v", got, want)
	}
}
func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didnt want %v", got)
	}
}


func AssertTrue(t *testing.T, got bool){
  t.Helper()
  if !got{
    t.Errorf("got %v, want true", got)
  }
}
func AssertFalse(t *testing.T, got bool){
  t.Helper()
  if got{
    t.Errorf("got %v, want false", got)
  }
}
