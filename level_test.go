package tinyrogue

import (
	"testing"
)

func TestLevel(t *testing.T) {
	currentGameData = NewGameData(15, 15)

	level := NewLevel()
	if len(level.Rooms) == 0 {
		t.Error("Failed to create level")
	}
	if len(level.Tiles) == 0 {
		t.Error("Failed to create level")
	}
}
