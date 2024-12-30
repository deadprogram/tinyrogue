package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Player represents the player character in the game.
type Player struct {
	*character
	ViewRadius int
}

// NewPlayer creates a new Player and initializes the data.
func NewPlayer(name string, kind string, img firefly.Image, speed int) *Player {
	return &Player{
		character: &character{
			name:  name,
			kind:  kind,
			Image: img,
			speed: speed,
		},
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
		level.SetViewRadius(p.ViewRadius)
		level.RayCast(pos.X, pos.Y)
	}

	if !level.InBounds(pos.X+x, pos.Y+y) {
		return
	}

	index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
	tile := level.Tiles[index]

	if !tile.Blocked {
		// player is moving away from this tile
		level.Block(pos.X, pos.Y, false)
		g.Player.Move(x, y)

		// player has moved to this tile, so it is now blocked
		level.Tiles[index].Blocked = true
	} else if x != 0 || y != 0 {
		if level.Tiles[index].TileType != WALL {
			// Its a tile with a creature -- now what?
			creature := g.GetCreatureForTile(index)
			if creature != nil && g.ActionSystem != nil {
				g.ActionSystem.Action(g.Player, creature)
			}
		}
	}
}
