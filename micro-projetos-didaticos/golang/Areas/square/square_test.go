package square

import "testing"

func TestSquareArea(t *testing.T) {
	testCases := []struct {
		side     float64
		expected float64
	}{
		{5, 25},
		{0, 0},
		{10.5, 110.25},
	}

	for _, tc := range testCases {
		sq := Square{Side: tc.side}
		actual := sq.Area()
		if actual != tc.expected {
			t.Errorf("For side %.2f, expected area %.2f but got %.2f", tc.side, tc.expected, actual)
		}
	}
}
