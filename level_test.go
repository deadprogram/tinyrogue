package tinyrogue

import (
	"testing"
)

func TestLevel(t *testing.T) {
	game := NewGame()
	game.SetData(NewGameData(16, 10))
	game.SetMap(NewGameMap())

	level := game.Map.CurrentLevel
	if len(level.Rooms) == 0 {
		t.Error("failed to create rooms for level")
	}
	if len(level.Tiles) != 160 {
		t.Errorf("incorrect number of tiles for level, wanted 160, got %d", len(level.Tiles))
	}
}
