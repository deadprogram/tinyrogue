package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Game holds all data the entire game will need.
type Game struct {
	Debug       bool
	Map         GameMap
	Data        GameData
	Turn        TurnState
	TurnCounter int
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

// Update is called on each frame loop
// The default value is 1/60 [s]
func (g *Game) Update() error {
	g.TurnCounter++
	// Update is called 60 times a second, so this
	// just adds a small delay which is
	// good enough for a turn-based game.
	// if g.Turn == PlayerTurn && g.TurnCounter > 20 {
	// 	TakePlayerAction(g)
	// }

	// if g.Turn == MonsterTurn {
	// 	UpdateMonster(g)
	// }
	return nil
}

// Draw is called each on each frame loop
func (g *Game) Render() {
	//Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel()

	// // Draw other renderables
	// ProcessRenderables(g, level, screen)

	// if g.Debug {
	// 	gd := NewGameData()
	// 	debug := fmt.Sprintf(
	// 		"FPS: %.0f\nSize: %d rows x %d cols\nDimensions: %dx%dpx\nTurnCounter: %d",
	// 		ebiten.ActualFPS(),
	// 		gd.Cols,
	// 		gd.Rows,
	// 		gd.GameWidth(),
	// 		gd.GameHeight(),
	// 		g.TurnCounter)
	// 	ebitenutil.DebugPrint(screen, debug)
	// }

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
