package tinyrogue

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

type CreatureBehavior int

const (
	CreatureIgnore CreatureBehavior = iota
	CreatureApproach
	CreatureAvoid
)

type Creature struct {
	*Character
	CurrentBehavior CreatureBehavior
}

func NewCreature(img *firefly.Image, speed int) *Creature {
	return &Creature{
		Character: &Character{
			Image: img,
			speed: speed,
		},
	}
}

func (c *Creature) SetBehavior(b CreatureBehavior) {
	c.CurrentBehavior = b
}

func (c *Creature) Update() {
	switch c.CurrentBehavior {
	case CreatureApproach:
		c.Approach()
	}
}

func (c *Creature) Approach() {
	l := CurrentGame().Map.CurrentLevel
	playerPosition := CurrentGame().Player.GetPosition()
	creaturePos := c.GetPosition()

	//creatureFoV := fov.New()
	//creatureFoV.Compute(&l, creaturePos.X, creaturePos.Y, 8)
	//if creatureFoV.IsVisible(playerPosition.X, playerPosition.Y) {
	if true {
		if creaturePos.GetManhattanDistance(playerPosition) == 1 {
			// The creature is right next to the player. Now what?
			// AttackSystem(game, creaturePos, &playerPosition)
		} else {
			path := AStar{}.GetPath(l, creaturePos, playerPosition)
			if len(path) > 1 {
				nextTile := l.Tiles[l.GetIndexFromXY(path[1].X, path[1].Y)]
				if !nextTile.Blocked {
					l.Tiles[l.GetIndexFromXY(creaturePos.X, creaturePos.Y)].Blocked = false

					c.MoveTo(Position{X: path[1].X, Y: path[1].Y})
					nextTile.Blocked = true
				}
			}
		}
	}
}
