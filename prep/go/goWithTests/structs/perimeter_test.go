package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if want != got {
		t.Errorf("got: '%.2f' - want:'%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 20.0}, want: 200.0},
    {shape: Circle{Radius: 20.0}, want: 1256.64},
    {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		gotRounded := math.Round(got*100) / 100
		if tt.want != gotRounded {
      //%#v will read a struct and print out all its fields
			t.Errorf("%#v got: '%g' - want:'%g'",tt.shape, gotRounded, tt.want)
		}
	}

}
