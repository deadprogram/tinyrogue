package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Character is the interface for all characters in the game.
type Character interface {
	Name() string
	Kind() string
	SetImage(img *firefly.Image)
	GetSpeed() int
	SetSpeed(speed int)
	GetPosition() *Position
	Move(dx, dy int)
	MoveTo(pos Position)
	Draw()
	Update()
	IsVisible() bool
	SetVisible(visible bool)
}

// character is the base type for all characters in the game.
type character struct {
	name  string
	kind  string
	Image *firefly.Image
	pos   Position
	speed int
}

// Name returns the name of the character.
func (c *character) Name() string {
	return c.name
}

// Kind returns the kind of the character.
func (c *character) Kind() string {
	return c.kind
}

// SetImage sets the image for the character.
func (c *character) SetImage(img *firefly.Image) {
	c.Image = img
}

// GetSpeed returns the speed of the character. Lower is faster.
func (c *character) GetSpeed() int {
	return c.speed
}

// SetSpeed sets the speed of the character. Lower is faster.
func (c *character) SetSpeed(speed int) {
	c.speed = speed
}

// GetPosition returns the position of the character.
func (c *character) GetPosition() *Position {
	return &c.pos
}

// Move moves the character by the given amount.
func (c *character) Move(dx, dy int) {
	c.pos.X += dx
	c.pos.Y += dy
}

// MoveTo moves the character to the given position.
func (c *character) MoveTo(pos Position) {
	c.pos.X = pos.X
	c.pos.Y = pos.Y
}

// Draw draws the character on the screen.
func (c *character) Draw() {
	gd := CurrentGame().Data
	firefly.DrawImage(*c.Image, firefly.Point{X: c.pos.X * gd.TileWidth, Y: c.pos.Y * gd.TileHeight})
}
