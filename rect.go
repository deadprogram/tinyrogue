package tinyrogue

// Rect represents a rectangle in 2D space.
type Rect struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// NewRect creates a new Rect given the top-left corner, width, and height.
func NewRect(x int, y int, width int, height int) Rect {
	return Rect{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

// Center returns the center point of the rectangle.
func (r *Rect) Center() (int, int) {
	centerX := (r.X1 + r.X2) / 2
	centerY := (r.Y1 + r.Y2) / 2
	return centerX, centerY
}

// Intersect returns true if this rectangle intersects with another.
func (r *Rect) Intersect(other Rect) bool {
	return (r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y1 && r.Y2 >= other.Y1)
}

// Contains returns true if the rectangle contains the given point.
func (r *Rect) Contains(x, y int) bool {
	return (x >= r.X1 && x <= r.X2 && y >= r.Y1 && y <= r.Y2)
}
