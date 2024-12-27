package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Game holds all data the entire game will need.
type Game struct {
	Debug       bool
	Map         GameMap
	Data        GameData
	Turn        TurnState
	TurnCounter int
	Player      *Player
	Creatures   []*Creature
	Images      map[string]*firefly.Image
}

var currentGame *Game

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{
		Images:    make(map[string]*firefly.Image),
		Creatures: make([]*Creature, 0),
	}
	g.Debug = true

	g.Turn = PlayerTurn
	g.TurnCounter = 0

	// only one game at a time
	currentGame = g
	return g
}

func (g *Game) SetMap(m GameMap) {
	g.Map = m
}

func (g *Game) SetData(d GameData) {
	g.Data = d
}

func (g *Game) SetPlayer(p *Player) {
	g.Player = p
}

func (g *Game) AddCreature(c *Creature) {
	g.Creatures = append(g.Creatures, c)
}

// Update is called on each frame loop
// The default value is 1/60 [s]
func (g *Game) Update() error {
	g.TurnCounter++
	if g.TurnCounter%g.Player.GetSpeed() == 0 {
		if g.Turn == PlayerTurn {
			g.Player.Update()
			g.Turn = CreatureTurn
		}
	}
	if g.Turn == CreatureTurn {
		for _, c := range g.Creatures {
			if g.TurnCounter%c.GetSpeed() == 0 {
				c.Update()
			}
		}
		g.Turn = PlayerTurn
	}

	return nil
}

// Draw is called each on each frame loop
func (g *Game) Render() {
	// Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel()

	// Draw the player
	g.Player.Draw()

	// Draw the creatures
	for _, c := range g.Creatures {
		c.Draw()
	}
}

// Layout accepts an outside size, which is a window size on desktop,
// and returns the game's logical screen size.
func (g *Game) Layout(w, h int) (int, int) {
	gd := g.Data
	return gd.GameWidth(), gd.GameHeight()
}

func CurrentGame() *Game {
	return currentGame
}
