package tinyrogue

// GameMap holds all the level and aggregate information for the entire world.
type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel Level
}

// NewGameMap creates a new set of maps for the entire game.
// For now we will only create 1 map (default)
func NewGameMap() GameMap {
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{Name: "default", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	gm := GameMap{Dungeons: dungeons, CurrentLevel: l}
	return gm
}
