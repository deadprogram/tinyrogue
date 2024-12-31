package tinyrogue

import "strconv"

// GameMap holds all the level and aggregate information for the entire world.
type GameMap struct {
	Name           string
	Dungeons       []Dungeon
	CurrentDungeon string //*Dungeon
	CurrentLevel   string //*Level
}

// NewGameMap creates a new set of maps for the entire game
// Using the predefined levels and dungeons.
func NewGameMap(name string, dungeons []Dungeon, startDungeon string, startLevel string) *GameMap {
	return &GameMap{
		Name:           name,
		Dungeons:       dungeons,
		CurrentDungeon: startDungeon,
		CurrentLevel:   startLevel}
}

// NewGeneratedGameMap generated a new set of dungeons and levels for the entire game.
func NewGeneratedGameMap(name string, dungeonCount int, levelCount int, floors, walls string) *GameMap {
	dungeons := make([]Dungeon, 0)
	for i := 0; i < dungeonCount; i++ {
		d := NewDungeon(name+"-"+strconv.Itoa(i), floors, walls)
		for j := 0; j < levelCount; j++ {
			nextLevel := NewLevel(d.Name+"-"+strconv.Itoa(j), d.FloorTypes, d.WallTypes)
			d.Levels = append(d.Levels, nextLevel)
		}
		dungeons = append(dungeons, d)
	}

	logDebug("Total dungeons created: " + strconv.Itoa(len(dungeons)))

	// generate the first level of the first dungeon
	dungeons[0].Levels[0].Generate()

	// put exit on first level to second level
	if levelCount > 1 {
		portalImg := CurrentGame().Images["portal"]
		p := NewPortal("portal", &portalImg, &dungeons[0], dungeons[0].Levels[1])
		dungeons[0].Levels[0].SetExit(p, dungeons[0].Levels[0].OpenLocation())
	}

	return &GameMap{Name: name, Dungeons: dungeons, CurrentDungeon: dungeons[0].Name, CurrentLevel: dungeons[0].Levels[0].Name}
}

// NewSingleGameMap creates a single level generated game map.
func NewSingleLevelGameMap() *GameMap {
	return NewGeneratedGameMap("Dungeon", 1, 1, "floor", "wall")
}

// NewSingleGameMapWithTerrain creates a single level generated game map.
func NewSingleGameMapWithTerrain(floors, walls string) *GameMap {
	return NewGeneratedGameMap("Dungeon", 1, 1, floors, walls)
}

func (gm *GameMap) Dungeon(name string) *Dungeon {
	for _, d := range gm.Dungeons {
		if d.Name == name {
			return &d
		}
	}
	return nil
}

func (gm *GameMap) NextDungeon() *Dungeon {
	for i, d := range gm.Dungeons {
		if d.Name == gm.CurrentDungeon {
			if i+1 < len(gm.Dungeons) {
				return &gm.Dungeons[i+1]
			}
		}
	}
	return nil
}

func (gm *GameMap) SetCurrentLevel(d *Dungeon, l *Level) {
	gm.CurrentDungeon = d.Name
	gm.CurrentLevel = l.Name
}
