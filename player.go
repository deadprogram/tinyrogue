package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

type Player struct {
	*Character
}

func NewPlayer() *Player {
	return &Player{
		Character: &Character{},
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

	if !tile.Blocked {
		level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
		g.Player.Move(x, y)
		level.Tiles[index].Blocked = true
		//level.PlayerFoV.Compute(level, pos.X, pos.Y, 8)
	} else if x != 0 || y != 0 {
		//if level.Tiles[index].TileType != WALL {
		// Its a tile with a creature -- now what?
		// monsterPosition := Position{X: pos.X + x, Y: pos.Y + y}
		//AttackSystem(g, pos, &monsterPosition)
		//}
	}

	//g.Turn = CreatureTurn
}
