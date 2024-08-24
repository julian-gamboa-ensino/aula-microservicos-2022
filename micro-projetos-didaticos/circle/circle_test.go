package circle

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	testCases := []struct {
		radius   float64
		expected float64
	}{
		{1, math.Pi},
		{0, 0},
		{5.2, math.Pi * 5.2 * 5.2},
	}

	for _, tc := range testCases {
		c := Circle{Radius: tc.radius}
		actual := c.Area()
		if actual != tc.expected {
			t.Errorf("For radius %.2f, expected area %.2f but got %.2f", tc.radius, tc.expected, actual)
		}
	}
}
