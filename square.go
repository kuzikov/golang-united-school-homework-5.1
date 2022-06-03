package square

// Point ...
type Point struct {
	x, y int
}

// Square ...
type Square struct {
	start Point
	a     uint
}

// End ...
func (r Square) End() Point {
	return Point{
		x: r.start.x + int(r.a),
		y: r.start.y + int(r.a),
	}
}

// Area ...
func (r Square) Area() uint {
	a := r.End().x - r.start.x
	b := r.End().y - r.start.y
	return uint(a * b)
}

// Perimeter ...
func (r Square) Perimeter() uint {
	a := r.End().x - r.start.x
	b := r.End().y - r.start.y
	return uint(a * b)
}
