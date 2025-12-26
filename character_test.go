package tinyrogue

import (
	"testing"

	"github.com/firefly-zero/firefly-go/firefly"
)

func TestCharacterName(t *testing.T) {
	c := &character{name: "Hero"}
	if c.Name() != "Hero" {
		t.Errorf("expected name 'Hero', got '%s'", c.Name())
	}
}

func TestCharacterKind(t *testing.T) {
	c := &character{kind: "warrior"}
	if c.Kind() != "warrior" {
		t.Errorf("expected kind 'warrior', got '%s'", c.Kind())
	}
}

func TestCharacterSetImage(t *testing.T) {
	c := &character{}
	img := &firefly.Image{}
	c.SetImage(img)
	if c.Image != img {
		t.Error("SetImage did not set the image correctly")
	}
}

func TestCharacterGetSpeed(t *testing.T) {
	c := &character{speed: 5}
	if c.GetSpeed() != 5 {
		t.Errorf("expected speed 5, got %d", c.GetSpeed())
	}
}

func TestCharacterSetSpeed(t *testing.T) {
	c := &character{}
	c.SetSpeed(10)
	if c.GetSpeed() != 10 {
		t.Errorf("expected speed 10, got %d", c.GetSpeed())
	}
}

func TestCharacterGetPosition(t *testing.T) {
	c := &character{pos: Position{X: 3, Y: 7}}
	pos := c.GetPosition()
	if pos.X != 3 || pos.Y != 7 {
		t.Errorf("expected position (3, 7), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestCharacterMove(t *testing.T) {
	c := &character{pos: Position{X: 5, Y: 5}}
	c.Move(2, -3)
	pos := c.GetPosition()
	if pos.X != 7 || pos.Y != 2 {
		t.Errorf("expected position (7, 2), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestCharacterMoveNegative(t *testing.T) {
	c := &character{pos: Position{X: 10, Y: 10}}
	c.Move(-5, -5)
	pos := c.GetPosition()
	if pos.X != 5 || pos.Y != 5 {
		t.Errorf("expected position (5, 5), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestCharacterMoveTo(t *testing.T) {
	c := &character{pos: Position{X: 0, Y: 0}}
	c.MoveTo(Position{X: 10, Y: 20})
	pos := c.GetPosition()
	if pos.X != 10 || pos.Y != 20 {
		t.Errorf("expected position (10, 20), got (%d, %d)", pos.X, pos.Y)
	}
}

func TestCharacterMoveToOverwrite(t *testing.T) {
	c := &character{pos: Position{X: 100, Y: 200}}
	c.MoveTo(Position{X: 5, Y: 5})
	pos := c.GetPosition()
	if pos.X != 5 || pos.Y != 5 {
		t.Errorf("expected position (5, 5), got (%d, %d)", pos.X, pos.Y)
	}
}
