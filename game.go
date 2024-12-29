package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Game holds all data the entire game will need.
type Game struct {
	Debug bool
	Map   GameMap
	Data  GameData

	// TurnBased is a flag to determine if the game is turn based or real-time.
	TurnBased   bool
	Turn        TurnState
	TurnCounter int

	// Player and Creatures
	Player    Character
	Creatures []Character

	// Images that are cached for space efficiency. Used for tiles and creatures.
	Images map[string]firefly.Image

	// UseFOV is a flag to determine if the game should use Field of View.
	UseFOV bool

	// ActionSystem is the interface for the game to handle actions between characters.
	ActionSystem Actionable

	// DialogShowing is a flag to determine if a dialog is currently showing.
	DialogShowing bool

	// currentDialog is the dialog that is currently showing.
	currentDialog *Dialog
}

var currentGame *Game

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{
		Images:    make(map[string]firefly.Image),
		Creatures: make([]Character, 0),
	}
	g.Debug = true

	g.Turn = PlayerTurn
	g.TurnCounter = 0

	// only one game at a time
	currentGame = g
	return g
}

// SetMap sets the map for the game.
func (g *Game) SetMap(m GameMap) {
	g.Map = m
}

// SetData sets the data for the game.
func (g *Game) SetData(d GameData) {
	g.Data = d
}

// SetPlayer sets the player for the game.
func (g *Game) SetPlayer(p Character) {
	g.Player = p
}

// SetActionSystem sets the action system for the game.
func (g *Game) SetActionSystem(a Actionable) {
	g.ActionSystem = a
}

// AddCreature adds a creature to the game.
func (g *Game) AddCreature(c Character) {
	g.Creatures = append(g.Creatures, c)
}

// RemoveCreature removes a creature from the game.
func (g *Game) RemoveCreature(c Character) {
	for i, creature := range g.Creatures {
		if creature == c {
			g.Creatures = append(g.Creatures[:i], g.Creatures[i+1:]...)
			break
		}
	}
}

// GetCreatureByName returns a creature by name.
func (g *Game) GetCreatureByName(name string) Character {
	for _, creature := range g.Creatures {
		if creature.Name() == name {
			return creature
		}
	}
	return nil
}

// Update is called on each frame loop
// The default value is 1/60 [s]
func (g *Game) Update() {
	if g.DialogShowing {
		g.currentDialog.Update()
		if !g.currentDialog.Confirmed {
			return
		}
		g.DialogShowing = false
	}

	if g.Turn == GameOver {
		return
	}

	g.TurnCounter++
	if g.TurnCounter%g.Player.GetSpeed() == 0 {
		if g.Turn == PlayerTurn || !g.TurnBased {
			g.Player.Update()
			g.Turn = CreatureTurn
		}
	}
	if g.Turn == CreatureTurn || !g.TurnBased {
		for _, c := range g.Creatures {
			if g.TurnCounter%c.GetSpeed() == 0 {
				c.Update()
			}
		}
		if g.Turn != GameOver {
			g.Turn = PlayerTurn
		}
	}
}

// Draw is called each on each frame loop
func (g *Game) Render() {
	firefly.ClearScreen(firefly.ColorBlack)

	// Draw the Map
	g.Map.CurrentLevel.Draw()

	// Draw the player
	g.Player.Draw()

	// Draw the creatures
	for _, c := range g.Creatures {
		if c.IsVisible() || !g.UseFOV {
			c.Draw()
		}
	}

	if g.DialogShowing && g.currentDialog != nil {
		g.currentDialog.Draw()
	}
}

// Layout accepts an outside size, which is a window size on desktop,
// and returns the game's logical screen size.
func (g *Game) Layout(w, h int) (int, int) {
	return g.Data.GameWidth(), g.Data.GameHeight()
}

// CurrentGame returns the current game.
func CurrentGame() *Game {
	return currentGame
}

// GetCreatureForTile returns the creature for the given tile index.
func (g *Game) GetCreatureForTile(index int) Character {
	for _, c := range g.Creatures {
		pos := c.GetPosition()
		if g.GetIndexFromXY(pos.X, pos.Y) == index {
			return c
		}
	}
	return nil
}

// GetIndexFromXY returns the index for the given x and y coordinates.
func (g *Game) GetIndexFromXY(x int, y int) int {
	return (y * g.Data.Cols) + x
}

// ShowDialog shows a dialog on the screen.
func (g *Game) ShowDialog(dlg *Dialog) {
	g.DialogShowing = true
	g.currentDialog = dlg
}

func (g *Game) LoadImage(name string) *firefly.Image {
	img := firefly.LoadFile(name, nil).Image()
	g.Images[name] = img

	return &img
}
