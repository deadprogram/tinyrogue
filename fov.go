package tinyrogue

import (
	"math"
)

func clearVisibility(tiles []*MapTile) {
	for _, tile := range tiles {
		tile.IsRevealed = false
	}
}

func getIndexFromXY(x int, y int) int {
	gd := CurrentGame().Data
	return (y * gd.Cols) + x
}

// adapted from https://www.albertford.com/shadowcasting/

func computeFoV(origin Position, tiles []*MapTile) {
	tiles[getIndexFromXY(origin.X, origin.Y)].IsRevealed = true
	for i := 0; i < 4; i++ {
		quadrant := Quadrant{cardinal: i, origin: origin}
		first_row := Row{depth: 1.0, start_slope: -1.0, end_slope: 1.0}
		scan(first_row, tiles, &quadrant)
	}
}

func reveal(pos Position, tiles []*MapTile, quadrant *Quadrant) {
	q := transform(quadrant, pos)
	if inBounds(q.X, q.Y) {
		tiles[getIndexFromXY(pos.X, pos.Y)].IsRevealed = true
	}
}

func isWall(pos Position, tiles []*MapTile, quadrant *Quadrant) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}

	var w bool
	q := transform(quadrant, pos)
	if inBounds(q.X, q.Y) {
		w = tiles[getIndexFromXY(pos.X, pos.X)].Blocked
	} else {
		w = true
	}

	return w
}

func inBounds(x int, y int) bool {
	gd := CurrentGame().Data

	if x < gd.Cols && y < gd.Rows && x >= 0 && y >= 0 {
		return true
	} else {
		return false
	}
}

func isFloor(tile Position, tiles []*MapTile, quadrant *Quadrant) bool {
	if tile.X < 0 || tile.Y < 0 {
		return false
	}

	var f bool
	q := transform(quadrant, tile)
	if inBounds(q.X, q.Y) {
		f = !tiles[getIndexFromXY(tile.X, tile.Y)].Blocked
	}

	return f
}

func scan(row Row, m []*MapTile, quadrant *Quadrant) {
	rows := []Row{row}
	for len(rows) > 0 {
		row = rows[len(rows)-1]
		rows = rows[:len(rows)-1]
		prev_tile := Position{X: -1, Y: -1}
		for _, tile := range tiles(row) {
			if isWall(tile, m, quadrant) || isSymmetric(row, tile) {
				reveal(tile, m, quadrant)
			}
			if isWall(prev_tile, m, quadrant) && isFloor(tile, m, quadrant) {
				row.start_slope = slope(tile)
			}
			if isFloor(prev_tile, m, quadrant) && isWall(tile, m, quadrant) {
				next_row := next(&row)
				next_row.end_slope = slope(tile)
				rows = append(rows, next_row)
			}
			prev_tile = tile
		}
		if isFloor(prev_tile, m, quadrant) {
			rows = append(rows, next(&row))
		}
	}
}

const (
	north = iota
	south
	east
	west
)

type Quadrant struct {
	cardinal int
	origin   Position
}

func transform(self *Quadrant, tile Position) Position {
	row, col := tile.X, tile.Y
	var v Position
	switch self.cardinal {
	case north:
		v = Position{X: self.origin.X + col, Y: self.origin.Y - row}
	case south:
		v = Position{X: self.origin.X + col, Y: self.origin.Y + row}
	case east:
		v = Position{X: self.origin.X + row, Y: self.origin.Y + col}
	case west:
		v = Position{X: self.origin.X - row, Y: self.origin.Y + col}
	}
	return v
}

type Row struct {
	depth       float64
	start_slope float64
	end_slope   float64
}

func tiles(self Row) []Position {
	min_col := roundTiesUp(self.depth * self.start_slope)
	max_col := roundTiesDown(self.depth * self.end_slope)

	var tiles []Position
	for col := min_col; col < max_col+1; col++ {
		tiles = append(tiles, Position{X: int(self.depth), Y: col})
	}
	return tiles
}

func next(self *Row) Row {
	return Row{depth: self.depth + 1.0, start_slope: self.start_slope, end_slope: self.end_slope}
}

func slope(tile Position) float64 {
	row_depth, col := tile.X, tile.Y
	return float64(2*col-1) / float64(2*row_depth)
}

func isSymmetric(row Row, tile Position) bool {
	if tile.X < 0 || tile.Y < 0 {
		return false
	}

	col := tile.Y
	return float64(col) >= row.depth*row.start_slope && float64(col) <= row.depth*row.end_slope
}

func roundTiesUp(n float64) int {
	return int(math.Floor(n + 0.5))
}

func roundTiesDown(n float64) int {
	return int(math.Ceil(n - 0.5))
}
