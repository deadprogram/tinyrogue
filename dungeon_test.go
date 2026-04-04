package tinyrogue

import (
	"testing"
)

func TestNewDungeon(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor,floor2", "wall,wall2")

	if d.Name != "TestDungeon" {
		t.Errorf("expected name 'TestDungeon', got '%s'", d.Name)
	}
	if d.FloorTypes != "floor,floor2" {
		t.Errorf("expected FloorTypes 'floor,floor2', got '%s'", d.FloorTypes)
	}
	if d.WallTypes != "wall,wall2" {
		t.Errorf("expected WallTypes 'wall,wall2', got '%s'", d.WallTypes)
	}
	if len(d.Levels) != 0 {
		t.Errorf("expected 0 levels, got %d", len(d.Levels))
	}
}

func TestDungeonCreateLevels(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor", "wall")

	d.CreateLevels(3)

	if len(d.Levels) != 3 {
		t.Errorf("expected 3 levels, got %d", len(d.Levels))
	}

	// Check level names
	if d.Levels[0].Name != "TestDungeon-0" {
		t.Errorf("expected level name 'TestDungeon-0', got '%s'", d.Levels[0].Name)
	}
	if d.Levels[1].Name != "TestDungeon-1" {
		t.Errorf("expected level name 'TestDungeon-1', got '%s'", d.Levels[1].Name)
	}
	if d.Levels[2].Name != "TestDungeon-2" {
		t.Errorf("expected level name 'TestDungeon-2', got '%s'", d.Levels[2].Name)
	}
}

func TestDungeonCreateLevelsAppends(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor", "wall")

	d.CreateLevels(2)
	d.CreateLevels(2)

	if len(d.Levels) != 4 {
		t.Errorf("expected 4 levels, got %d", len(d.Levels))
	}

	// Check that names are sequential
	if d.Levels[2].Name != "TestDungeon-2" {
		t.Errorf("expected level name 'TestDungeon-2', got '%s'", d.Levels[2].Name)
	}
	if d.Levels[3].Name != "TestDungeon-3" {
		t.Errorf("expected level name 'TestDungeon-3', got '%s'", d.Levels[3].Name)
	}
}

func TestDungeonLevel(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor", "wall")
	d.CreateLevels(3)

	// Find existing level
	level := d.Level("TestDungeon-1")
	if level == nil {
		t.Fatal("expected to find level 'TestDungeon-1'")
	}
	if level.Name != "TestDungeon-1" {
		t.Errorf("expected level name 'TestDungeon-1', got '%s'", level.Name)
	}

	// Find non-existing level
	nonExistent := d.Level("NonExistent")
	if nonExistent != nil {
		t.Error("expected nil for non-existent level")
	}
}

func TestDungeonNextLevel(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor", "wall")
	d.CreateLevels(3)

	// Get next level from first
	nextLevel := d.NextLevel(d.Levels[0])
	if nextLevel == nil {
		t.Fatal("expected next level to exist")
	}
	if nextLevel.Name != "TestDungeon-1" {
		t.Errorf("expected next level 'TestDungeon-1', got '%s'", nextLevel.Name)
	}

	// Get next level from second
	nextLevel = d.NextLevel(d.Levels[1])
	if nextLevel == nil {
		t.Fatal("expected next level to exist")
	}
	if nextLevel.Name != "TestDungeon-2" {
		t.Errorf("expected next level 'TestDungeon-2', got '%s'", nextLevel.Name)
	}

	// Get next level from last (should be nil)
	nextLevel = d.NextLevel(d.Levels[2])
	if nextLevel != nil {
		t.Error("expected nil for next level from last level")
	}
}

func TestDungeonNextLevelNonExistent(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor", "wall")
	d.CreateLevels(2)

	// Create a level that's not in the dungeon
	otherLevel := NewLevel("OtherLevel", "floor", "wall")

	nextLevel := d.NextLevel(otherLevel)
	if nextLevel != nil {
		t.Error("expected nil for next level of non-existent level")
	}
}

func TestDungeonLevelFloorAndWallTypes(t *testing.T) {
	d := NewDungeon("TestDungeon", "grass,dirt", "tree,rock")
	d.CreateLevels(1)

	level := d.Levels[0]
	if level.FloorTypes != "grass,dirt" {
		t.Errorf("expected FloorTypes 'grass,dirt', got '%s'", level.FloorTypes)
	}
	if level.WallTypes != "tree,rock" {
		t.Errorf("expected WallTypes 'tree,rock', got '%s'", level.WallTypes)
	}
}

func TestDungeonCreateZeroLevels(t *testing.T) {
	d := NewDungeon("TestDungeon", "floor", "wall")
	d.CreateLevels(0)

	if len(d.Levels) != 0 {
		t.Errorf("expected 0 levels, got %d", len(d.Levels))
	}
}
