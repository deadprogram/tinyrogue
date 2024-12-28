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

// Creature is the type for all creatures in the game.
type Creature struct {
	*character
	CurrentBehavior CreatureBehavior
	Visible         bool
}

// NewCreature creates a new Creature and initializes the data
func NewCreature(name string, kind string, img *firefly.Image, speed int) *Creature {
	return &Creature{
		character: &character{
			name:  name,
			kind:  kind,
			Image: img,
			speed: speed,
		},
	}
}

func (c *Creature) IsVisible() bool {
	return c.Visible
}

func (c *Creature) SetVisible(visible bool) {
	c.Visible = visible
}

// SetBehavior sets the behavior of the creature.
func (c *Creature) SetBehavior(b CreatureBehavior) {
	c.CurrentBehavior = b
}

// Update updates the creature.
func (c *Creature) Update() {
	switch c.CurrentBehavior {
	case CreatureApproach:
		c.Approach()
	case CreatureAvoid:
		c.Avoid()
	}
}

// Approach moves the creature towards the player.
func (c *Creature) Approach() {
	level := CurrentGame().Map.CurrentLevel
	playerPosition := CurrentGame().Player.GetPosition()
	creaturePos := c.GetPosition()

	if c.Visible || !CurrentGame().UseFOV {
		if creaturePos.GetManhattanDistance(playerPosition) == 1 {
			// The creature is right next to the player. Now what?
			if CurrentGame().ActionSystem != nil {
				CurrentGame().ActionSystem.Action(c, CurrentGame().Player)
			}
		} else {
			path := AStar{}.GetPath(level, creaturePos, playerPosition)
			if len(path) > 1 {
				nextTile := level.Tiles[level.GetIndexFromXY(path[1].X, path[1].Y)]
				if !nextTile.Blocked {
					level.Block(creaturePos.X, creaturePos.Y, false)

					c.MoveTo(Position{X: path[1].X, Y: path[1].Y})
					nextTile.Blocked = true
				}
			}
		}
	}
}

func (c *Creature) Avoid() {
	logDebug("Avoid is not implemented yet")
}
