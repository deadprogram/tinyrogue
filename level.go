package tinyrogue

import (
	"strconv"
	"strings"

	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
)

// Each of the map tiles will be represented  by one of these structures
type MapTile struct {
	PixelX   int // Upper left corner of the tile
	PixelY   int
	Blocked  bool           // The tile should block the player or monster ?
	Image    *firefly.Image // Pointer to an Image
	Visible  bool
	Explored bool
	TileType TileType
}

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles      []*MapTile
	Rooms      []Rect
	playerFoV  *FieldOfVision
	ViewRadius int
}

// NewLevel creates a new game level in a dungeon.
func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	l.playerFoV = &FieldOfVision{}
	l.playerFoV.InitializeFOV()

	return l
}

// Tiles will be stored in one slice. We will use GetIndexFromXY to
// determine which tile to return.
// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	return CurrentGame().GetIndexFromXY(x, y)
}

// createTiles creates a map of all walls as a baseline for carving out a level.
func (level *Level) createTiles() []*MapTile {
	gd := CurrentGame().Data
	tiles := make([]*MapTile, gd.Rows*gd.Cols)
	index := 0
	for x := 0; x < gd.Cols; x++ {
		for y := 0; y < gd.Rows; y++ {
			index = level.GetIndexFromXY(x, y)
			tile := MapTile{
				PixelX:   x * gd.TileWidth,
				PixelY:   y * gd.TileHeight,
				Blocked:  true,
				Image:    level.getWallImage(),
				TileType: WALL,
			}
			tiles[index] = &tile
		}
	}

	logDebug("Total tiles created: " + strconv.Itoa(len(tiles)))
	return tiles
}

// createRoom creates a room in the level.
func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			if index > 0 && index < len(level.Tiles) {
				level.Tiles[index].Blocked = false
				level.Tiles[index].TileType = FLOOR
				level.Tiles[index].Image = level.getFloorImage()
			}
		}
	}
}

// getWallImage returns a random wall image from the list of wall images.
func (level *Level) getWallImage() *firefly.Image {
	walls := strings.Split(CurrentGame().Data.WallTypes, ",")
	if len(walls) < 2 {
		return CurrentGame().Images[walls[0]]
	}
	wall := walls[GetDiceRoll(len(walls))-1]
	return CurrentGame().Images[wall]
}

// getFloorImage returns a random floor image from the list of floor images.
func (level *Level) getFloorImage() *firefly.Image {
	floors := strings.Split(CurrentGame().Data.FloorTypes, ",")
	if len(floors) < 2 {
		return CurrentGame().Images[floors[0]]
	}

	floor := floors[GetDiceRoll(len(floors))-1]
	return CurrentGame().Images[floor]
}

// GenerateLevelTiles creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	gd := CurrentGame().Data
	tiles := level.createTiles()
	level.Tiles = tiles
	contains_rooms := false

	for idx := 0; idx < gd.MaxRooms; idx++ {
		w := GetRandomBetween(gd.MinSize, gd.MaxSize)
		h := GetRandomBetween(gd.MinSize, gd.MaxSize)
		x := GetDiceRoll(gd.Cols - w - 1)
		y := GetDiceRoll(gd.Rows - h - 1)

		new_room := NewRect(x, y, w, h)
		okToAdd := true
		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			if contains_rooms {
				newX, newY := new_room.Center()
				prevX, prevY := level.Rooms[len(level.Rooms)-1].Center()

				coinflip := GetDiceRoll(2)

				if coinflip == 2 {
					level.createHorizontalTunnel(prevX, newX, prevY)
					level.createVerticalTunnel(prevY, newY, newX)
				} else {
					level.createHorizontalTunnel(prevX, newX, newY)
					level.createVerticalTunnel(prevY, newY, prevX)
				}
			}
			level.Rooms = append(level.Rooms, new_room)
			contains_rooms = true
		}
	}
	logDebug("Total rooms created: " + strconv.Itoa(len(level.Rooms)))
}

// createHorizontalTunnel creates a horizontal tunnel between two points.
func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	gd := CurrentGame().Data
	for x := tinymath.Min(float32(x1), float32(x2)); x < tinymath.Max(float32(x1), float32(x2))+1; x++ {
		index := level.GetIndexFromXY(int(x), y)
		if index > 0 && index < gd.Rows*gd.Cols {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Image = level.getFloorImage()
		}
	}
}

// createVerticalTunnel creates a vertical tunnel between two points.
func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := CurrentGame().Data
	for y := tinymath.Min(float32(y1), float32(y2)); y < tinymath.Max(float32(y1), float32(y2))+1; y++ {
		index := level.GetIndexFromXY(x, int(y))
		if index > 0 && index < gd.Rows*gd.Cols {
			level.Tiles[index].TileType = FLOOR
			level.Tiles[index].Blocked = false
			level.Tiles[index].Image = level.getFloorImage()
		}
	}
}

// InBounds checks if the given x and y coordinates are within the level bounds.
func (level *Level) InBounds(x, y int) bool {
	gd := CurrentGame().Data
	if x < 0 || x > gd.Cols-1 || y < 0 || y > gd.Rows-1 {
		return false
	}
	return true
}

// IsOpaque checks if the given x and y coordinates are within the level bounds.
func (level *Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].TileType == WALL
}

// Block sets the blocked property of a tile at the given x and y coordinates.
func (level *Level) Block(x, y int, block bool) {
	level.Tiles[level.GetIndexFromXY(x, y)].Blocked = block
}

// Draw the level.
func (level *Level) Draw() {
	gd := CurrentGame().Data
	for x := 0; x < gd.Cols; x++ {
		for y := 0; y < gd.Rows; y++ {
			idx := level.GetIndexFromXY(x, y)
			tile := level.Tiles[idx]
			if tile.Visible || tile.Explored || !CurrentGame().UseFOV {
				firefly.DrawImage(*tile.Image, firefly.Point{X: tile.PixelX, Y: tile.PixelY})
			}
		}
	}
}

// RayCast casts out rays each degree in a 360 circle from the player, to help determine what the player can see.
func (level *Level) RayCast(playerX, playerY int) {
	level.playerFoV.RayCast(playerX, playerY, level)
}

// SetViewRadius sets the view radius for the player.
func (level *Level) SetViewRadius(radius int) {
	level.playerFoV.SetTorchRadius(radius)
}

// RandomLocation returns a random location in the level.
func (level *Level) RandomLocation() (Position, bool) {
	if len(level.Rooms) == 0 {
		return Position{}, false
	}
	randomRoom := level.Rooms[GetDiceRoll(len(level.Rooms))-1]
	x, y := randomRoom.Center()

	idx := level.GetIndexFromXY(x, y)
	tile := level.Tiles[idx]

	return Position{x, y}, !tile.Blocked
}

// Dump prints the level to the console.
func (level *Level) Dump() {
	gd := CurrentGame().Data
	hdr := "  "
	for x := 0; x < gd.Cols; x++ {
		hdr += strconv.Itoa(x % 10)
	}
	logDebug(hdr)

	for y := 0; y < gd.Rows; y++ {
		data := strconv.Itoa(y%10) + " "
		for x := 0; x < gd.Cols; x++ {
			idx := level.GetIndexFromXY(x, y)
			tile := level.Tiles[idx]
			if tile.TileType == WALL {
				data += "#"
			} else {
				data += "."
			}
		}
		logDebug(data)
	}
}
