package tinyrogue

// GameData holds the values for the size of elements within the game
type GameData struct {
	Cols       int
	Rows       int
	TileWidth  int
	TileHeight int
	MinSize    int
	MaxSize    int
	MaxRooms   int
	FloorTypes string
	WallTypes  string
}

// NewGameData creates a fully populated GameData Struct.
func NewGameData(cols, rows, tilewidth, tileheight int) GameData {
	g := GameData{
		Cols:       cols,
		Rows:       rows,
		TileWidth:  tilewidth,
		TileHeight: tileheight,
		MinSize:    4,
		MaxSize:    12,
		MaxRooms:   36,
		FloorTypes: "floor",
		WallTypes:  "wall",
	}
	return g
}

// GameWidth returns the width of the game in pixels.
func (gd *GameData) GameWidth() int {
	return gd.TileWidth * gd.Cols
}

// GameHeight returns the height of the game in pixels.
func (gd *GameData) GameHeight() int {
	return gd.TileHeight * gd.Rows
}
