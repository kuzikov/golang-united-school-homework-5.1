package square

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		sq   Square
		area uint
	}{
		{sq: Square{start: Point{1, 1}, a: 5}, area: 25},
		// {sq: Square{start: Point{0, 0},a: }, area: 0},
	}

	for _, test := range tests {
		if ar := test.sq.Area(); ar != test.area {
			t.Errorf("Have: <%v>, Want: <%v>\n", ar, test.area)
		}
	}
}
