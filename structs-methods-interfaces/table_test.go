package smi

import "testing"

func TestAreaTable(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// adding the field names makes the intent clearer for others
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
