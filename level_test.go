package tinyrogue

import (
	"testing"
)

func TestLevel(t *testing.T) {
	game := NewGame()
	game.SetData(NewGameData(42, 24))
	game.SetMap(NewGameMap())

	level := game.Map.CurrentLevel
	if len(level.Rooms) == 0 {
		t.Error("Failed to create level")
	}
	if len(level.Tiles) == 0 {
		t.Error("Failed to create level")
	}
}
