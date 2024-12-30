package tinyrogue

import (
	"testing"
)

func TestLevel(t *testing.T) {
	game := NewGame()
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	level := game.Map.CurrentLevel
	if len(level.Rooms) == 0 {
		t.Error("failed to create rooms for level")
	}
	if len(level.Tiles) != 160 {
		t.Errorf("incorrect number of tiles for level, wanted 160, got %d", len(level.Tiles))
	}

	pos := level.OpenLocation()
	if pos.X < 0 || pos.X >= 16 || pos.Y < 0 || pos.Y >= 10 {
		t.Errorf("invalid open location %v", pos)
	}
}
