package smi

import "testing"

func TestAreaTable(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// adding the field names makes the intent clearer for others
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

// func TestAreaTable(t *testing.T) {
// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		// adding the field names makes the intent clearer for others
// 		{shape: Rectangle{12, 6}, want: 72.0},
// 		{shape: Circle{10}, want: 314.1592653589793},
// 		{shape: Triangle{12, 6}, want: 36.0},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %g want %g", got, tt.want)
// 		}
// 	}
// }
