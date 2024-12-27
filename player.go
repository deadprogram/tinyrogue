package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

type Player struct {
	*character
}

func NewPlayer(name string, img *firefly.Image, speed int) *Player {
	return &Player{
		character: &character{
			name:  name,
			Image: img,
			speed: speed,
		},
	}
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
	index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
	tile := level.Tiles[index]

	if g.UseFOV {
		level.RayCast(pos.X, pos.Y)
	}

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
