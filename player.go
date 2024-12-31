package tinyrogue

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

// Player represents the player character in the game.
type Player struct {
	*character
	fov              *FieldOfVision
	ViewRadius       int
	levelSwitchDelay int
}

// NewPlayer creates a new Player and initializes the data.
func NewPlayer(name string, kind string, img *firefly.Image, speed int) *Player {
	fov := &FieldOfVision{}
	fov.InitializeFOV()

	return &Player{
		character: &character{
			name:  name,
			kind:  kind,
			Image: img,
			speed: speed,
		},
		fov:        fov,
		ViewRadius: 3,
	}
}

// IsVisible always returns true, because the player is always visible.
func (p *Player) IsVisible() bool {
	return true
}

// SetVisible is just here to fulfill [Character] interface.
func (p *Player) SetVisible(visible bool) {
}

// Update updates the player.
func (p *Player) Update() {
	g := CurrentGame()
	x, y := 0, 0

	pad, _ := firefly.ReadPad(firefly.Combined)
	switch {
	case pad.DPad().Down:
		y = -1
	case pad.DPad().Up:
		y = 1
	case pad.DPad().Right:
		x = 1
	case pad.DPad().Left:
		x = -1
	}

	pos := g.Player.GetPosition()
	level := g.Map.CurrentLevel

	if g.UseFOV {
		p.fov.SetTorchRadius(p.ViewRadius)
		p.fov.RayCast(pos.X, pos.Y, level)
	}

	if !level.InBounds(pos.X+x, pos.Y+y) {
		return
	}

	index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
	tile := level.Tiles[index]

	switch tile.TileType {
	case ENTRANCE:
		// The player has reached the entrance to the previous level.
		p.levelSwitchDelay++
		if p.levelSwitchDelay > 10 {
			buttons := firefly.ReadButtons(firefly.Combined)
			if buttons.N || buttons.S || buttons.E || buttons.W {
				logDebug("Entrance reached")
				level.Block(pos, false)
				// TODO: handle if is dungeon entrance
				prevLevel := level.Entrance.Destination
				p.MoveTo(prevLevel.GetExitPosition())
				g.Map.CurrentLevel = prevLevel
				p.levelSwitchDelay = 0

				return
			}
		}
	case EXIT:
		// The player has reached the exit to the next level.
		p.levelSwitchDelay++
		if p.levelSwitchDelay > 10 {
			buttons := firefly.ReadButtons(firefly.Combined)
			if buttons.N || buttons.S || buttons.E || buttons.W {
				logDebug("Exit reached")
				level.Block(pos, false)
				// TODO: handle if is dungeon exit

				// generate a new level?
				nextLevel := level.Exit.Destination
				if !nextLevel.Generated {
					nextLevel.GenerateAndConnect(level)
				}

				p.MoveTo(nextLevel.GetEntrancePosition())
				g.Map.CurrentLevel = nextLevel
				p.levelSwitchDelay = 0

				return
			}
		}
	}

	if !tile.Blocked {
		// player is moving
		g.Player.Move(x, y)

		// unblock the previous tile
		level.Block(pos, false)
		// player has moved to this tile, so it is now blocked
		tile.Blocked = true
	} else if x != 0 || y != 0 {
		if tile.TileType != WALL {
			// Its a tile with a creature -- now what?
			creature := g.GetCreatureForTile(index)
			if creature != nil && g.ActionSystem != nil {
				g.ActionSystem.Action(g.Player, creature)
			}
		}
	}
}
