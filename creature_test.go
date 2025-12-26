package tinyrogue

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

func TestNewCreature(t *testing.T) {
	img := &firefly.Image{}
	c := NewCreature("Goblin", "monster", img, 3)

	if c.Name() != "Goblin" {
		t.Errorf("expected name 'Goblin', got '%s'", c.Name())
	}
	if c.Kind() != "monster" {
		t.Errorf("expected kind 'monster', got '%s'", c.Kind())
	}
	if c.GetSpeed() != 3 {
		t.Errorf("expected speed 3, got %d", c.GetSpeed())
	}
	if c.Image != img {
		t.Error("expected image to be set")
	}
}

func TestCreatureVisibility(t *testing.T) {
	c := NewCreature("Goblin", "monster", nil, 3)

	if c.IsVisible() {
		t.Error("expected creature to not be visible by default")
	}

	c.SetVisible(true)
	if !c.IsVisible() {
		t.Error("expected creature to be visible after SetVisible(true)")
	}

	c.SetVisible(false)
	if c.IsVisible() {
		t.Error("expected creature to not be visible after SetVisible(false)")
	}
}

func TestCreatureBehavior(t *testing.T) {
	c := NewCreature("Goblin", "monster", nil, 3)

	if c.CurrentBehavior != CreatureIgnore {
		t.Errorf("expected default behavior CreatureIgnore, got %d", c.CurrentBehavior)
	}

	c.SetBehavior(CreatureApproach)
	if c.CurrentBehavior != CreatureApproach {
		t.Errorf("expected behavior CreatureApproach, got %d", c.CurrentBehavior)
	}

	c.SetBehavior(CreatureAvoid)
	if c.CurrentBehavior != CreatureAvoid {
		t.Errorf("expected behavior CreatureAvoid, got %d", c.CurrentBehavior)
	}
}

func TestCreatureMove(t *testing.T) {
	c := NewCreature("Goblin", "monster", nil, 3)
	c.MoveTo(Position{X: 5, Y: 5})

	pos := c.GetPosition()
	if pos.X != 5 || pos.Y != 5 {
		t.Errorf("expected position (5, 5), got (%d, %d)", pos.X, pos.Y)
	}

	c.Move(2, -1)
	pos = c.GetPosition()
	if pos.X != 7 || pos.Y != 4 {
		t.Errorf("expected position (7, 4), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestCreatureBehaviorConstants(t *testing.T) {
	if CreatureIgnore != 0 {
		t.Errorf("expected CreatureIgnore to be 0, got %d", CreatureIgnore)
	}
	if CreatureApproach != 1 {
		t.Errorf("expected CreatureApproach to be 1, got %d", CreatureApproach)
	}
	if CreatureAvoid != 2 {
		t.Errorf("expected CreatureAvoid to be 2, got %d", CreatureAvoid)
	}
}

func TestCreatureUpdateWithIgnoreBehavior(t *testing.T) {
	game := NewGame()
	game.Images["floor"] = firefly.Image{}
	game.Images["wall"] = firefly.Image{}
	game.SetData(NewGameData(16, 10, 16, 16))
	game.SetMap(NewSingleLevelGameMap())

	player := NewPlayer("Player", "player", nil, 5)
	player.MoveTo(Position{X: 5, Y: 5})
	game.SetPlayer(player)

	c := NewCreature("Goblin", "monster", nil, 3)
	c.MoveTo(Position{X: 8, Y: 8})
	c.SetBehavior(CreatureIgnore)

	initialPos := c.GetPosition()

	// Update should do nothing with CreatureIgnore behavior
	c.Update()

	newPos := c.GetPosition()
	if initialPos.X != newPos.X || initialPos.Y != newPos.Y {
		t.Errorf("expected position to remain (%d, %d), got (%d, %d)",
			initialPos.X, initialPos.Y, newPos.X, newPos.Y)
	}
}
