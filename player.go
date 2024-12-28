package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

type Player struct {
	*character
	ViewRadius int
}

func NewPlayer(name string, kind string, img *firefly.Image, speed int) *Player {
	return &Player{
		character: &character{
			name:  name,
			kind:  kind,
			Image: img,
			speed: speed,
		},
	}
}

func (p *Player) IsVisible() bool {
	return true
}

func (p *Player) SetVisible(visible bool) {
}

func (p *Player) Update() {
	g := CurrentGame()
	x, y := 0, 0

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
		level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
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
