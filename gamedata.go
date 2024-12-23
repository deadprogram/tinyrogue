package tinyrogue

// GameData holds the values for the size of elements within the game
type GameData struct {
	Cols       int
	Rows       int
	TileWidth  int
	TileHeight int
}

// NewGameData creates a fully populated GameData Struct.
func NewGameData(cols, rows int) GameData {
	g := GameData{
		Cols:       cols,
		Rows:       rows,
		TileWidth:  16,
		TileHeight: 16,
	}
	return g
}

func (gd *GameData) GameWidth() int {
	return gd.TileWidth * gd.Cols
}

func (gd *GameData) GameHeight() int {
	return gd.TileHeight * gd.Rows
}

// func CurrentGameData() GameData {
// 	return CurrentGame().Data
// }
