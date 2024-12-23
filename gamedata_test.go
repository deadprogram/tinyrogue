package tinyrogue

import (
	"testing"
)

func TestGameData(t *testing.T) {
	gameData := NewGameData(42, 24)
	if gameData.GameHeight() != 24*16 {
		t.Error("Failed to create game data")
	}
	if gameData.GameWidth() != 42*16 {
		t.Error("Failed to create game data")
	}
}
