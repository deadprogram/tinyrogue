package tinyrogue

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

func TestAStar(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = &firefly.Image{}
	game.Images["wall"] = &firefly.Image{}

	game.SetData(NewGameData(15, 15))
	game.SetMap(NewGameMap())

	level := game.Map.CurrentLevel

	room1 := level.Rooms[0]
	room2 := level.Rooms[len(level.Rooms)-1]

	// Set the start and end points
	x1, y1 := room1.Center()
	start := &Position{x1, y1}
	x2, y2 := room2.Center()
	end := &Position{x2, y2}
	// Create the AStar object
	as := AStar{}
	// Get the path
	path := as.GetPath(level, start, end)
	// Check the path
	if len(path) == 0 {
		t.Errorf("Expected path length of > 0, got %d", len(path))
	}
}
