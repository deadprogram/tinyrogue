package tinyrogue

import "strconv"

// GameMap holds all the level and aggregate information for the entire world.
type GameMap struct {
	Name         string
	Dungeons     []Dungeon
	CurrentLevel *Level
}

// NewGameMap creates a new set of maps for the entire game
// Using the predefined levels and dungeons.
func NewGameMap(name string, dungeons []Dungeon, start *Level) *GameMap {
	return &GameMap{Name: name, Dungeons: dungeons, CurrentLevel: start}
}

// NewGeneratedGameMap generated a new set of dungeons and levels for the entire game.
func NewGeneratedGameMap(name string, dungeonCount int, levelCount int, floors, walls string) *GameMap {
	portalImg := CurrentGame().Images["portal"]

	dungeons := make([]Dungeon, 0)
	for i := 0; i < dungeonCount; i++ {
		d := NewDungeon(name+"-"+strconv.Itoa(i), floors, walls)
		var prevLevel *Level
		for j := 0; j < levelCount; j++ {
			l := NewLevel(d.FloorTypes, d.WallTypes)
			l.Generate()

			// connect portals between levels
			if prevLevel != nil {
				// connect the previous level to the current level
				prevLevel.SetExit(NewPortal("portal", &portalImg, l), prevLevel.OpenLocation())
				l.SetEntrance(NewPortal("portal", &portalImg, prevLevel), l.OpenLocation())
			}
			prevLevel = l

			d.Levels = append(d.Levels, l)
		}

		// put an exit on the last level?
		dungeons = append(dungeons, d)
	}

	logDebug("Total dungeons created: " + strconv.Itoa(len(dungeons)))

	return &GameMap{Name: name, Dungeons: dungeons, CurrentLevel: dungeons[0].Levels[0]}
}

// NewSingleGameMap creates a single level generated game map.
func NewSingleLevelGameMap() *GameMap {
	return NewGeneratedGameMap("Dungeon", 1, 1, "floor", "wall")
}

// NewSingleGameMapWithTerrain creates a single level generated game map.
func NewSingleGameMapWithTerrain(floors, walls string) *GameMap {
	return NewGeneratedGameMap("Dungeon", 1, 1, floors, walls)
}
