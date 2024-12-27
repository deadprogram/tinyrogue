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
func NewCreature(name string, img *firefly.Image, speed int) *Creature {
	return &Creature{
		character: &character{
			name:  name,
			Image: img,
			speed: speed,
		},
	}
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
	l := CurrentGame().Map.CurrentLevel
	playerPosition := CurrentGame().Player.GetPosition()
	creaturePos := c.GetPosition()

	//creatureFoV := fov.New()
	//creatureFoV.Compute(&l, creaturePos.X, creaturePos.Y, 8)
	//if creatureFoV.IsVisible(playerPosition.X, playerPosition.Y) {
	isVisible := true
	if isVisible {
		if creaturePos.GetManhattanDistance(playerPosition) == 1 {
			// The creature is right next to the player. Now what?
			if CurrentGame().ActionSystem != nil {
				CurrentGame().ActionSystem.Action(c, CurrentGame().Player)
			}
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

func (c *Creature) Avoid() {
	logDebug("Avoid is not implemented yet")
}
