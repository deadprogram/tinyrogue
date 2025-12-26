package tinyrogue

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

func TestLevel(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	level := game.CurrentLevel()
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

func TestMultiLevel(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewGeneratedGameMap("Big Forest", 1, 5, "floor", "wall")
	game.SetMap(gm)

	level := game.CurrentLevel()
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

func TestExitIsReachable(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.Images["portal"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewGeneratedGameMap("TestDungeon", 1, 2, "floor", "wall")
	game.SetMap(gm)

	level := game.CurrentLevel()

	// Get player start position and exit position
	startPos := level.OpenLocation()
	exitPos := level.GetExitPosition()

	// Verify exit tile is not blocked
	exitTile := level.Tiles[level.GetIndexFromXY(exitPos.X, exitPos.Y)]
	if exitTile.Blocked {
		t.Error("exit tile should not be blocked")
	}

	// Verify exit tile type
	if exitTile.TileType != EXIT {
		t.Errorf("expected EXIT tile type, got %d", exitTile.TileType)
	}

	// Verify path exists from start to exit
	as := AStar{}
	path := as.GetPath(level, startPos, exitPos)

	if len(path) == 0 {
		t.Errorf("no path found from start %v to exit %v", startPos, exitPos)
		level.Dump()
	}
}

func TestEntranceIsNotBlocked(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.Images["portal"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	level := NewLevel("TestLevel", "floor", "wall")
	level.Generate()

	portalImg := game.Images["portal"]
	dungeon := NewDungeon("TestDungeon", "floor", "wall")
	destLevel := NewLevel("DestLevel", "floor", "wall")

	portal := NewPortal("portal", &portalImg, &dungeon, destLevel)
	entrancePos := level.OpenLocation()
	level.SetEntrance(portal, entrancePos)

	entranceTile := level.Tiles[level.GetIndexFromXY(entrancePos.X, entrancePos.Y)]
	if entranceTile.Blocked {
		t.Error("entrance tile should not be blocked")
	}
	if entranceTile.TileType != ENTRANCE {
		t.Errorf("expected ENTRANCE tile type, got %d", entranceTile.TileType)
	}
}

func TestExitIsNotBlocked(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.Images["portal"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	level := NewLevel("TestLevel", "floor", "wall")
	level.Generate()

	portalImg := game.Images["portal"]
	dungeon := NewDungeon("TestDungeon", "floor", "wall")
	destLevel := NewLevel("DestLevel", "floor", "wall")

	portal := NewPortal("portal", &portalImg, &dungeon, destLevel)
	exitPos := level.OpenLocation()
	level.SetExit(portal, exitPos)

	exitTile := level.Tiles[level.GetIndexFromXY(exitPos.X, exitPos.Y)]
	if exitTile.Blocked {
		t.Error("exit tile should not be blocked")
	}
	if exitTile.TileType != EXIT {
		t.Errorf("expected EXIT tile type, got %d", exitTile.TileType)
	}
}

func TestExitIsReachableFromEntrance(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.Images["portal"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewGeneratedGameMap("TestDungeon", 1, 2, "floor", "wall")
	game.SetMap(gm)

	level := game.CurrentLevel()

	// Get any open location (simulating player start)
	startPos := level.OpenLocation()
	exitPos := level.GetExitPosition()

	// Verify path exists from any start position to exit
	as := AStar{}
	path := as.GetPath(level, startPos, exitPos)

	if len(path) == 0 {
		t.Errorf("no path found from start %v to exit %v", startPos, exitPos)
		level.Dump()
	}
}

func TestOpenLocationReachableFrom(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	level := NewLevel("TestLevel", "floor", "wall")
	level.Generate()

	startPos := level.OpenLocation()
	reachablePos := level.OpenLocationReachableFrom(startPos)

	// Verify the returned position is reachable
	as := AStar{}
	path := as.GetPath(level, startPos, reachablePos)

	if len(path) == 0 {
		t.Errorf("OpenLocationReachableFrom returned unreachable position: start %v, target %v", startPos, reachablePos)
		level.Dump()
	}
}

func TestMultipleLevelsHaveReachableExits(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.Images["portal"] = firefly.Image{}
	game.SetData(NewGameData(20, 15, 16, 16))

	// Test with multiple levels
	gm := NewGeneratedGameMap("TestDungeon", 1, 3, "floor", "wall")
	game.SetMap(gm)

	// Verify first level
	level := game.CurrentLevel()
	startPos := level.OpenLocation()
	exitPos := level.GetExitPosition()

	as := AStar{}
	path := as.GetPath(level, startPos, exitPos)

	if len(path) == 0 {
		t.Errorf("Level 1: no path found from start %v to exit %v", startPos, exitPos)
		level.Dump()
	}
}
