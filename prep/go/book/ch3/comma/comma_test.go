package comma

import "testing"

func TestComma(t *testing.T) {
	t.Run("it inserts a comma every three digits", func(t *testing.T) {
		got := Comma("12345")
		want := "12,345"

		if want != got {
			t.Errorf("wanted: %s - got: %s", want, got)
		}
	})
	t.Run("it does not insert comma for 3 digits", func(t *testing.T) {
		got := Comma("12")
		want := "12"

		if want != got {
			t.Errorf("wanted: %s - got: %s", want, got)
		}
	})
	t.Run("it works on large numbers too", func(t *testing.T) {
		got := Comma("12345678999")
		want := "12,345,678,999"

		if want != got {
			t.Errorf("wanted: %s - got: %s", want, got)
		}
	})
}
