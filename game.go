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
	Images      map[string]*firefly.Image
}

var currentGame *Game

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Debug = true

	g.Images = make(map[string]*firefly.Image)
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

// Update is called on each frame loop
// The default value is 1/60 [s]
func (g *Game) Update() error {
	g.TurnCounter++
	x, y := 0, 0
	if g.Turn == PlayerTurn && g.TurnCounter > 20 {
		buttons := firefly.ReadButtons(firefly.Combined)
		switch {
		case buttons.N:
			y = -1
		case buttons.S:
			y = 1
		case buttons.E:
			x = 1
		case buttons.W:
			x = -1
		}
	}

	pos := g.Player.GetPosition()
	level := g.Map.CurrentLevel
	index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
	tile := level.Tiles[index]

	if !tile.Blocked {
		level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
		g.Player.Move(x, y)
		level.Tiles[index].Blocked = true
		//level.PlayerFoV.Compute(level, pos.X, pos.Y, 8)
	} else if x != 0 || y != 0 {
		// if level.Tiles[index].TileType != WALL {
		// 	// Its a tile with a monster -- Fight it
		// 	monsterPosition := Position{X: pos.X + x, Y: pos.Y + y}
		// 	AttackSystem(g, pos, &monsterPosition)
		// }
	}

	// if g.Turn == MonsterTurn {
	// 	UpdateMonster(g)
	// }
	return nil
}

// Draw is called each on each frame loop
func (g *Game) Render() {
	// Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel()

	// Draw the player
	g.Player.Draw()

	// TODO: Draw the monsters
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
