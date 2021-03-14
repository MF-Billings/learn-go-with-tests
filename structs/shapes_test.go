package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// v2: refactored v1 to use Table driven tests
func TestArea(t *testing.T) {

	// if you want to test different implementations of an interface, or it the data being passed
	// in to a function has lots of different requirements that need testing then Table driven tests are a great fit
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}

// v1
// func TestArea(t *testing.T) {

// 	assertCorrectMessage := func(t testing.TB, got, want float64) {
// 		t.Helper()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		assertCorrectMessage(t, got, want)
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		r := Rectangle{12.0, 6.0}
// 		checkArea(t, r, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		c := Circle{10}
// 		checkArea(t, c, 314.1592653589793)
// 	})
// }
