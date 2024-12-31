package tinyrogue

import "strconv"

// Dungeon is a container for all the levels that make up
// a particular dungeon in the world.
type Dungeon struct {
	Name       string
	Levels     []*Level
	FloorTypes string
	WallTypes  string
}

func NewDungeon(name, floors, walls string) Dungeon {
	return Dungeon{
		Name:       name,
		Levels:     make([]*Level, 0),
		FloorTypes: floors,
		WallTypes:  walls,
	}
}

// NextLevel returns the next level in the dungeon after the given level.
func (d *Dungeon) NextLevel(l *Level) *Level {
	for i, level := range d.Levels {
		if level.Name == l.Name {
			if i+1 < len(d.Levels) {
				return d.Levels[i+1]
			}
		}
	}
	return nil
}

// CreateLevels creates a number of empty levels in the dungeon.
func (d *Dungeon) CreateLevels(n int) {
	start := len(d.Levels)
	for i := 0; i < n; i++ {
		nextLevel := NewLevel(d.Name+"-"+strconv.Itoa(start+i), d.FloorTypes, d.WallTypes)
		d.Levels = append(d.Levels, nextLevel)
	}
}

func (d *Dungeon) Level(name string) *Level {
	for _, level := range d.Levels {
		if level.Name == name {
			return level
		}
	}
	return nil
}
