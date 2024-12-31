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
	ENTRANCE
	EXIT
)

// Each of the map tiles will be represented  by one of these structures
type MapTile struct {
	PixelX   int // Upper left corner of the tile
	PixelY   int
	Blocked  bool           // tile should block the player or creatures?
	Image    *firefly.Image // image for this tile
	Visible  bool
	Explored bool
	TileType TileType
}

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Name       string
	Generated  bool
	Tiles      []*MapTile
	Rooms      []Rect
	FloorTypes string
	WallTypes  string
	ViewRadius int
	Entrance   *Portal
	Exit       *Portal
}

// NewLevel creates a new game level in a dungeon.
func NewLevel(name, floors, walls string) *Level {
	l := &Level{
		Name:       name,
		Rooms:      make([]Rect, 0),
		FloorTypes: floors,
		WallTypes:  walls,
	}

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
				level.setFloorTile(level.Tiles[index])
			}
		}
	}
}

// getWallImage returns a random wall image from the list of wall images.
func (level *Level) getWallImage() *firefly.Image {
	walls := strings.Split(level.WallTypes, ",")
	if len(walls) < 2 {
		img := CurrentGame().Images[walls[0]]
		return &img
	}
	wall := walls[GetDiceRoll(len(walls))-1]
	img := CurrentGame().Images[wall]
	return &img
}

// getFloorImage returns a random floor image from the list of floor images.
func (level *Level) getFloorImage() *firefly.Image {
	floors := strings.Split(level.FloorTypes, ",")
	if len(floors) < 2 {
		img := CurrentGame().Images[floors[0]]
		return &img
	}

	floor := floors[GetDiceRoll(len(floors))-1]
	img := CurrentGame().Images[floor]
	return &img
}

// Generate creates a new Dungeon Level Map.
func (level *Level) Generate() {
	if level.Generated {
		return
	}

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

	level.Generated = true
	logDebug("Total rooms created: " + strconv.Itoa(len(level.Rooms)))
}

func ConnectExits(startDungeon *Dungeon, startLevel *Level, destinationDungeon *Dungeon, destinationLevel *Level) {
	portalImg := CurrentGame().Images["portal"]
	p := NewPortal("portal", &portalImg, startDungeon, startLevel)
	destinationLevel.SetEntrance(p, destinationLevel.OpenLocation())

	// what is the destination level exit?
	nextLevel := destinationDungeon.NextLevel(destinationLevel)
	if nextLevel == nil {
		nextDungeon := CurrentGame().Map.NextDungeon()
		if nextDungeon == nil {
			// end of the game?
			logDebug("End of the game reached.")
			return
		}

		// to another dungeon, first level
		p = NewPortal("portal", &portalImg, nextDungeon, nextDungeon.Levels[0])
		destinationLevel.SetExit(p, destinationLevel.OpenLocation())
		return
	}

	p = NewPortal("portal", &portalImg, destinationDungeon, nextLevel)
	destinationLevel.SetExit(p, destinationLevel.OpenLocation())
}

// createHorizontalTunnel creates a horizontal tunnel between two points.
func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	gd := CurrentGame().Data
	for x := tinymath.Min(float32(x1), float32(x2)); x < tinymath.Max(float32(x1), float32(x2))+1; x++ {
		index := level.GetIndexFromXY(int(x), y)
		if index > 0 && index < gd.Rows*gd.Cols {
			level.setFloorTile(level.Tiles[index])
		}
	}
}

// createVerticalTunnel creates a vertical tunnel between two points.
func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := CurrentGame().Data
	for y := tinymath.Min(float32(y1), float32(y2)); y < tinymath.Max(float32(y1), float32(y2))+1; y++ {
		index := level.GetIndexFromXY(x, int(y))
		if index > 0 && index < gd.Rows*gd.Cols {
			level.setFloorTile(level.Tiles[index])
		}
	}
}

func (level *Level) setFloorTile(tile *MapTile) {
	tile.TileType = FLOOR
	tile.Image = level.getFloorImage()
	tile.Blocked = false
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

// Block sets the blocked property of a tile at the given [Position].
func (level *Level) Block(pos Position, block bool) {
	level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = block
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

// RandomLocation returns a random location in the level.
func (level *Level) RandomLocation() (Position, bool) {
	if len(level.Rooms) == 0 {
		return Position{}, false
	}
	randomRoom := level.Rooms[GetDiceRoll(len(level.Rooms))-1]
	x, y := randomRoom.Center()

	idx := level.GetIndexFromXY(x, y)
	tile := level.Tiles[idx]

	return Position{x, y}, !tile.Blocked && tile.TileType == FLOOR
}

// OpenLocation returns an open location in the level. Used for "spawning".
func (level *Level) OpenLocation() Position {
	for {
		pos, free := level.RandomLocation()
		if free {
			return pos
		}
	}
}

// SetEntrance sets the entrance to the level.
func (level *Level) SetEntrance(p *Portal, pos Position) {
	level.Entrance = p
	tile := level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)]
	tile.TileType = ENTRANCE
	tile.Image = p.Image
}

// GetEntrancePosition returns the position of the entrance for this level.
func (level *Level) GetEntrancePosition() Position {
	for i, tile := range level.Tiles {
		if tile.TileType == ENTRANCE {
			return Position{X: i % CurrentGame().Data.Cols, Y: i / CurrentGame().Data.Cols}
		}
	}
	logDebug("No entrance found")
	return Position{}
}

// SetExit sets the exit to the level.
func (level *Level) SetExit(p *Portal, pos Position) {
	level.Exit = p
	tile := level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)]
	tile.TileType = EXIT
	tile.Image = p.Image
}

// GetExitPosition returns the position of the exit for this level.
func (level *Level) GetExitPosition() Position {
	for i, tile := range level.Tiles {
		if tile.TileType == EXIT {
			return Position{X: i % CurrentGame().Data.Cols, Y: i / CurrentGame().Data.Cols}
		}
	}
	logDebug("No exit found")
	return Position{}
}

func (level *Level) GetRoom(x, y int) int {
	for i, r := range level.Rooms {
		if r.Contains(x, y) {
			return i
		}
	}
	return -1
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
			switch tile.TileType {
			case WALL:
				data += "#"
			case ENTRANCE, EXIT:
				data += "E"
			default:
				rm := level.GetRoom(x, y)
				if rm == -1 {
					data += "."
				} else {
					data += strconv.Itoa(rm % 10)
				}
			}
		}
		logDebug(data)
	}

	for i, r := range level.Rooms {
		logDebug("Room: " + strconv.Itoa(i) + ": " + strconv.Itoa(r.X1) + "," + strconv.Itoa(r.Y1) + " " + strconv.Itoa(r.X2) + "," + strconv.Itoa(r.Y2))
	}
}
