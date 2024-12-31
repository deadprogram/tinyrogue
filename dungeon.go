package tinyrogue

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
		if level == l {
			if i+1 < len(d.Levels) {
				return d.Levels[i+1]
			}
		}
	}
	return nil
}
