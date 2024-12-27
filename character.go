package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

type Character interface {
	SetImage(img *firefly.Image)
	GetSpeed() int
	SetSpeed(speed int)
	GetPosition() *Position
	Move(dx, dy int)
	MoveTo(pos Position)
	Draw()
}

type character struct {
	Image *firefly.Image
	pos   Position
	speed int
}

func NewCharacterizer() *character {
	return &character{}
}

func (c *character) SetImage(img *firefly.Image) {
	c.Image = img
}

func (c *character) GetSpeed() int {
	return c.speed
}

func (c *character) SetSpeed(speed int) {
	c.speed = speed
}

func (c *character) GetPosition() *Position {
	return &c.pos
}

func (c *character) Move(dx, dy int) {
	c.pos.X += dx
	c.pos.Y += dy
}

func (c *character) MoveTo(pos Position) {
	c.pos.X = pos.X
	c.pos.Y = pos.Y
}

func (c *character) Draw() {
	gd := CurrentGame().Data
	firefly.DrawImage(*c.Image, firefly.Point{X: c.pos.X * gd.TileWidth, Y: c.pos.Y * gd.TileHeight})
}
