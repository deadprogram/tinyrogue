package tinyrogue

import "github.com/orsinium-labs/tinymath"

type Position struct {
	X int
	Y int
}

func (p *Position) GetManhattanDistance(other Position) int {
	xDist := tinymath.Abs(float32(p.X - other.X))
	yDist := tinymath.Abs(float32(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

func (p *Position) IsEqual(other Position) bool {
	return (p.X == other.X && p.Y == other.Y)
}
