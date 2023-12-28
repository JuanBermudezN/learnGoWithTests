package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 35.0},
	}
	// The only new syntax here is creating an "anonymous struct", areaTests. We are declaring a slice of structs by using []struct with two fields, the shape and the want.
	// Then we fill the slice with cases.
	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
				// #v format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.
			}
		})
	}
}

// checkArea := func(t testing.TB, shape Shape, want float64) {
// 	t.Helper()
// 	got := shape.Area()
// 	if got != want {
// 		t.Errorf("got %g want %g", got, want)
// 	}
// }
// t.Run("rectangles", func(t *testing.T) {
// 	rectangle := Rectangle{6.0, 12.0}
// 	checkArea(t, rectangle, 72.0)
// })

// t.Run("circles", func(t *testing.T) {
// 	circle := Circle{10}
// 	checkArea(t, circle, 314.1592653589793)

// })
