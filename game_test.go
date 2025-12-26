package tinyrogue

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	if game == nil {
		t.Fatal("expected NewGame() to return a game")
	}
	if game.Images == nil {
		t.Error("expected Images map to be initialized")
	}
	if game.Creatures == nil {
		t.Error("expected Creatures slice to be initialized")
	}
	if game.Turn != PlayerTurn {
		t.Errorf("expected Turn to be PlayerTurn, got %d", game.Turn)
	}
	if game.TurnCounter != 0 {
		t.Errorf("expected TurnCounter to be 0, got %d", game.TurnCounter)
	}
	if !game.Debug {
		t.Error("expected Debug to be true by default")
	}
}

func TestCurrentGame(t *testing.T) {
	game := NewGame()

	current := CurrentGame()
	if current != game {
		t.Error("expected CurrentGame() to return the same game instance")
	}
}

func TestGameSetData(t *testing.T) {
	game := NewGame()
	gd := NewGameData(16, 10, 16, 16)

	game.SetData(gd)

	if game.Data.Cols != 16 {
		t.Errorf("expected Cols 16, got %d", game.Data.Cols)
	}
	if game.Data.Rows != 10 {
		t.Errorf("expected Rows 10, got %d", game.Data.Rows)
	}
	if game.Data.TileWidth != 16 {
		t.Errorf("expected TileWidth 16, got %d", game.Data.TileWidth)
	}
	if game.Data.TileHeight != 16 {
		t.Errorf("expected TileHeight 16, got %d", game.Data.TileHeight)
	}
}

func TestGameSetMap(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewSingleLevelGameMap()
	game.SetMap(gm)

	if game.Map == nil {
		t.Error("expected SetMap to set the game map")
	}
	if game.Map.Name != "Dungeon" {
		t.Errorf("expected map name 'Dungeon', got '%s'", game.Map.Name)
	}
}

func TestGameSetPlayer(t *testing.T) {
	game := NewGame()
	player := NewPlayer("Hero", "player", nil, 5)

	game.SetPlayer(player)

	if game.Player != player {
		t.Error("expected SetPlayer to set the player")
	}
	if game.Player.Name() != "Hero" {
		t.Errorf("expected player name 'Hero', got '%s'", game.Player.Name())
	}
}

func TestGameSetActionSystem(t *testing.T) {
	game := NewGame()
	action := &DebugAction{}

	game.SetActionSystem(action)

	if game.ActionSystem == nil {
		t.Error("expected SetActionSystem to set the action system")
	}
}

func TestGameAddCreature(t *testing.T) {
	game := NewGame()
	creature := NewCreature("Goblin", "monster", nil, 3)

	game.AddCreature(creature)

	if len(game.Creatures) != 1 {
		t.Errorf("expected 1 creature, got %d", len(game.Creatures))
	}
	if game.Creatures[0] != creature {
		t.Error("expected creature to be added")
	}
}

func TestGameAddMultipleCreatures(t *testing.T) {
	game := NewGame()
	creature1 := NewCreature("Goblin", "monster", nil, 3)
	creature2 := NewCreature("Orc", "monster", nil, 4)

	game.AddCreature(creature1)
	game.AddCreature(creature2)

	if len(game.Creatures) != 2 {
		t.Errorf("expected 2 creatures, got %d", len(game.Creatures))
	}
}

func TestGameRemoveCreature(t *testing.T) {
	game := NewGame()
	creature1 := NewCreature("Goblin", "monster", nil, 3)
	creature2 := NewCreature("Orc", "monster", nil, 4)

	game.AddCreature(creature1)
	game.AddCreature(creature2)
	game.RemoveCreature(creature1)

	if len(game.Creatures) != 1 {
		t.Errorf("expected 1 creature after removal, got %d", len(game.Creatures))
	}
	if game.Creatures[0] != creature2 {
		t.Error("expected remaining creature to be Orc")
	}
}

func TestGameRemoveNonExistentCreature(t *testing.T) {
	game := NewGame()
	creature1 := NewCreature("Goblin", "monster", nil, 3)
	creature2 := NewCreature("Orc", "monster", nil, 4)

	game.AddCreature(creature1)
	game.RemoveCreature(creature2)

	if len(game.Creatures) != 1 {
		t.Errorf("expected 1 creature, got %d", len(game.Creatures))
	}
}

func TestGameGetCreatureByName(t *testing.T) {
	game := NewGame()
	creature1 := NewCreature("Goblin", "monster", nil, 3)
	creature2 := NewCreature("Orc", "monster", nil, 4)

	game.AddCreature(creature1)
	game.AddCreature(creature2)

	found := game.GetCreatureByName("Orc")
	if found == nil {
		t.Error("expected to find creature 'Orc'")
	}
	if found.Name() != "Orc" {
		t.Errorf("expected creature name 'Orc', got '%s'", found.Name())
	}
}

func TestGameGetCreatureByNameNotFound(t *testing.T) {
	game := NewGame()
	creature := NewCreature("Goblin", "monster", nil, 3)

	game.AddCreature(creature)

	found := game.GetCreatureByName("Dragon")
	if found != nil {
		t.Error("expected nil for non-existent creature")
	}
}

func TestGameCurrentLevel(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	level := game.CurrentLevel()

	if level == nil {
		t.Error("expected CurrentLevel() to return a level")
	}
}

func TestGameCurrentDungeon(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	dungeon := game.CurrentDungeon()

	if dungeon == nil {
		t.Error("expected CurrentDungeon() to return a dungeon")
	}
}

func TestGameUseFOV(t *testing.T) {
	game := NewGame()

	if game.UseFOV {
		t.Error("expected UseFOV to be false by default")
	}

	game.UseFOV = true
	if !game.UseFOV {
		t.Error("expected UseFOV to be true after setting")
	}
}

func TestGameTurnBased(t *testing.T) {
	game := NewGame()

	if game.TurnBased {
		t.Error("expected TurnBased to be false by default")
	}

	game.TurnBased = true
	if !game.TurnBased {
		t.Error("expected TurnBased to be true after setting")
	}
}

func TestGameGetIndexFromXY(t *testing.T) {
	game := NewGame()
	game.SetData(NewGameData(16, 10, 16, 16))

	index := game.GetIndexFromXY(5, 3)
	expected := (3 * 16) + 5

	if index != expected {
		t.Errorf("expected index %d, got %d", expected, index)
	}
}

func TestGameGetIndexFromXYOrigin(t *testing.T) {
	game := NewGame()
	game.SetData(NewGameData(16, 10, 16, 16))

	index := game.GetIndexFromXY(0, 0)

	if index != 0 {
		t.Errorf("expected index 0, got %d", index)
	}
}

func TestGameGetCreatureForTile(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	creature := NewCreature("Goblin", "monster", nil, 3)
	creature.MoveTo(Position{X: 5, Y: 5})
	game.AddCreature(creature)

	index := game.GetIndexFromXY(5, 5)
	found := game.GetCreatureForTile(index)

	if found == nil {
		t.Error("expected to find creature at tile")
	}
	if found.Name() != "Goblin" {
		t.Errorf("expected creature name 'Goblin', got '%s'", found.Name())
	}
}

func TestGameGetCreatureForTileEmpty(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	index := game.GetIndexFromXY(5, 5)
	found := game.GetCreatureForTile(index)

	if found != nil {
		t.Error("expected nil for empty tile")
	}
}

func TestGameLayout(t *testing.T) {
	game := NewGame()
	game.SetData(NewGameData(16, 10, 16, 16))

	w, h := game.Layout(800, 600)

	if w != 256 {
		t.Errorf("expected width 256, got %d", w)
	}
	if h != 160 {
		t.Errorf("expected height 160, got %d", h)
	}
}

func TestGameDialogShowing(t *testing.T) {
	game := NewGame()

	if game.DialogShowing {
		t.Error("expected DialogShowing to be false by default")
	}
}

func TestGameImagesMap(t *testing.T) {
	game := NewGame()

	if game.Images == nil {
		t.Fatal("expected Images map to be initialized")
	}

	img := firefly.Image{}
	game.Images["test"] = img

	if _, ok := game.Images["test"]; !ok {
		t.Error("expected to store image in Images map")
	}
}

func TestGameTurnStateTransitions(t *testing.T) {
	game := NewGame()

	if game.Turn != PlayerTurn {
		t.Errorf("expected initial Turn to be PlayerTurn, got %d", game.Turn)
	}

	game.Turn = CreatureTurn
	if game.Turn != CreatureTurn {
		t.Errorf("expected Turn to be CreatureTurn, got %d", game.Turn)
	}

	game.Turn = GameOver
	if game.Turn != GameOver {
		t.Errorf("expected Turn to be GameOver, got %d", game.Turn)
	}
}

func TestGameMultipleDungeons(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewGeneratedGameMap("World", 3, 2, "floor", "wall")
	game.SetMap(gm)

	if len(game.Map.Dungeons) != 3 {
		t.Errorf("expected 3 dungeons, got %d", len(game.Map.Dungeons))
	}
}

func TestGameNextDungeon(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewGeneratedGameMap("World", 3, 2, "floor", "wall")
	game.SetMap(gm)

	next := game.NextDungeon()
	if next == nil {
		t.Error("expected NextDungeon() to return a dungeon")
	}
	if next.Name != "World-1" {
		t.Errorf("expected next dungeon name 'World-1', got '%s'", next.Name)
	}
}

func TestGameNextDungeonFromLast(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))

	gm := NewGeneratedGameMap("World", 1, 1, "floor", "wall")
	game.SetMap(gm)

	next := game.NextDungeon()
	if next != nil {
		t.Error("expected nil for next dungeon from last dungeon")
	}
}
