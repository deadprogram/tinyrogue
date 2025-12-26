package tinyrogue

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

func TestNewPlayer(t *testing.T) {
	img := &firefly.Image{}
	p := NewPlayer("Hero", "player", img, 5)

	if p.Name() != "Hero" {
		t.Errorf("expected name 'Hero', got '%s'", p.Name())
	}
	if p.Kind() != "player" {
		t.Errorf("expected kind 'player', got '%s'", p.Kind())
	}
	if p.GetSpeed() != 5 {
		t.Errorf("expected speed 5, got %d", p.GetSpeed())
	}
	if p.Image != img {
		t.Error("expected image to be set")
	}
}

func TestPlayerDefaultViewRadius(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)

	if p.ViewRadius != 3 {
		t.Errorf("expected default ViewRadius 3, got %d", p.ViewRadius)
	}
}

func TestPlayerSetViewRadius(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)
	p.ViewRadius = 5

	if p.ViewRadius != 5 {
		t.Errorf("expected ViewRadius 5, got %d", p.ViewRadius)
	}
}

func TestPlayerIsVisibleAlwaysTrue(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)

	if !p.IsVisible() {
		t.Error("expected player to always be visible")
	}
}

func TestPlayerSetVisibleHasNoEffect(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)

	p.SetVisible(false)
	if !p.IsVisible() {
		t.Error("expected player to still be visible after SetVisible(false)")
	}

	p.SetVisible(true)
	if !p.IsVisible() {
		t.Error("expected player to still be visible after SetVisible(true)")
	}
}

func TestPlayerMove(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)
	p.MoveTo(Position{X: 5, Y: 5})

	pos := p.GetPosition()
	if pos.X != 5 || pos.Y != 5 {
		t.Errorf("expected position (5, 5), got (%d, %d)", pos.X, pos.Y)
	}

	p.Move(2, -1)
	pos = p.GetPosition()
	if pos.X != 7 || pos.Y != 4 {
		t.Errorf("expected position (7, 4), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestPlayerMoveTo(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)
	p.MoveTo(Position{X: 10, Y: 20})

	pos := p.GetPosition()
	if pos.X != 10 || pos.Y != 20 {
		t.Errorf("expected position (10, 20), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestPlayerFOVInitialized(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)

	if p.fov == nil {
		t.Error("expected FOV to be initialized")
	}
}

func TestPlayerGetSpeed(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 10)

	if p.GetSpeed() != 10 {
		t.Errorf("expected speed 10, got %d", p.GetSpeed())
	}
}

func TestPlayerSetSpeed(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)
	p.SetSpeed(15)

	if p.GetSpeed() != 15 {
		t.Errorf("expected speed 15, got %d", p.GetSpeed())
	}
}

func TestPlayerSetImage(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)
	img := &firefly.Image{}

	p.SetImage(img)

	if p.Image != img {
		t.Error("expected image to be set")
	}
}

func TestPlayerLevelSwitchDelayInitialValue(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)

	if p.levelSwitchDelay != 0 {
		t.Errorf("expected levelSwitchDelay 0, got %d", p.levelSwitchDelay)
	}
}

func TestPlayerInGameContext(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	player := NewPlayer("Hero", "player", nil, 5)
	game.SetPlayer(player)

	level := game.CurrentLevel()
	pos := level.OpenLocation()
	player.MoveTo(pos)

	if game.Player != player {
		t.Error("expected player to be set in game")
	}

	playerPos := game.Player.GetPosition()
	if playerPos.X != pos.X || playerPos.Y != pos.Y {
		t.Errorf("expected player position (%d, %d), got (%d, %d)",
			pos.X, pos.Y, playerPos.X, playerPos.Y)
	}
}

func TestPlayerCharacterInterface(t *testing.T) {
	var c Character = NewPlayer("Hero", "player", nil, 5)

	if c.Name() != "Hero" {
		t.Errorf("expected name 'Hero', got '%s'", c.Name())
	}
	if c.Kind() != "player" {
		t.Errorf("expected kind 'player', got '%s'", c.Kind())
	}
	if !c.IsVisible() {
		t.Error("expected player to be visible")
	}
}

func TestPlayerMultipleViewRadiusChanges(t *testing.T) {
	p := NewPlayer("Hero", "player", nil, 5)

	p.ViewRadius = 1
	if p.ViewRadius != 1 {
		t.Errorf("expected ViewRadius 1, got %d", p.ViewRadius)
	}

	p.ViewRadius = 10
	if p.ViewRadius != 10 {
		t.Errorf("expected ViewRadius 10, got %d", p.ViewRadius)
	}

	p.ViewRadius = 5
	if p.ViewRadius != 5 {
		t.Errorf("expected ViewRadius 5, got %d", p.ViewRadius)
	}
}
