package tinyrogue

import (
	"github.com/orsinium-labs/tinymath"
)

// Based on Gogue https://github.com/gogue-framework/gogue

// FieldOfVision represents an area that an entity can see, defined by the torch radius. The cos and sin tables are
// generated once on instantiation, so we don't have to build them each time we want to calculate visible distances.
type FieldOfVision struct {
	cosTable    map[int]float32
	sinTable    map[int]float32
	torchRadius int
}

// InitializeFOV generates the cos and sin tables, for 360 degrees, for use when raycasting to determine line of sight
func (f *FieldOfVision) InitializeFOV() {
	f.cosTable = make(map[int]float32)
	f.sinTable = make(map[int]float32)

	for i := 0; i < 360; i++ {
		ax := tinymath.Sin(float32(i) / (float32(180) / tinymath.Pi))
		ay := tinymath.Cos(float32(i) / (float32(180) / tinymath.Pi))

		f.sinTable[i] = ax
		f.cosTable[i] = ay
	}
}

// SetTorchRadius sets the radius of the FOVs torch, or how far the entity can see
func (f *FieldOfVision) SetTorchRadius(radius int) {
	if radius > 1 {
		f.torchRadius = radius
	}
}

// Equal to the code that used to live in main's game:clearFOV()
// SetAllInvisible makes all tiles on the gamemap invisible to the player.
func (f *FieldOfVision) SetAllInvisible(m *Level) {
	for _, tile := range m.Tiles {
		tile.Visible = false
	}
}

// RayCast casts out rays each degree in a 360 circle from the player. If a ray passes over a floor (does not block sight)
// tile, keep going, up to the maximum torch radius (view radius) of the player. If the ray intersects a wall
// (blocks sight), stop, as the player will not be able to see past that. Every visible tile will get the Visible
// and Explored properties set to true.
func (f *FieldOfVision) RayCast(playerX, playerY int, m *Level) {
	for i := 0; i < 360; i++ {

		ax := f.sinTable[i]
		ay := f.cosTable[i]

		x := float32(playerX)
		y := float32(playerY)

		// Mark the players current position as explored
		tile := m.Tiles[m.GetIndexFromXY(playerX, playerY)]
		tile.Explored = true
		tile.Visible = true

		for j := 0; j < f.torchRadius; j++ {
			x -= ax
			y -= ay

			roundedX := int(round(x))
			roundedY := int(round(y))

			if !inBounds(roundedX, roundedY) {
				break
			}

			tile := m.Tiles[m.GetIndexFromXY(roundedX, roundedY)]
			tile.Explored = true
			tile.Visible = true

			// any creature on this tile should be visible
			creature := CurrentGame().GetCreatureForTile(m.GetIndexFromXY(roundedX, roundedY))
			if creature != nil {
				creature.SetVisible(true)
			}

			if tile.TileType == WALL {
				// The ray hit a wall, go no further
				break
			}
		}
	}
}

func round(f float32) float32 {
	return tinymath.Floor(f + .5)
}

func inBounds(x, y int) bool {
	gd := CurrentGame().Data

	if x < gd.Cols && y < gd.Rows && x >= 0 && y >= 0 {
		return true
	} else {
		return false
	}
}
