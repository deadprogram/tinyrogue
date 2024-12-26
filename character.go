package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

type Character struct {
	Image *firefly.Image
	pos   Position
}

func NewCharacter() *Character {
	return &Character{}
}

func (c *Character) SetImage(img *firefly.Image) {
	c.Image = img
}

func (c *Character) GetPosition() *Position {
	return &c.pos
}

func (c *Character) SetPosition(pos *Position) {
	c.pos = *pos
}

func (c *Character) Move(dx, dy int) {
	c.pos.X += dx
	c.pos.Y += dy
}

func (c *Character) MoveTo(pos *Position) {
	c.pos = *pos
}

func (c *Character) Draw() {
	gd := CurrentGame().Data
	firefly.DrawImage(*c.Image, firefly.Point{X: c.pos.X * gd.TileWidth, Y: c.pos.Y * gd.TileHeight})
}
