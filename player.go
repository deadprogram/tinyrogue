package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

type Player struct {
	Image *firefly.Image
	pos   Position
}

func NewPlayer() *Player {
	return &Player{
		Image: CurrentGame().Images["player"],
		pos:   Position{X: 1, Y: 1},
	}
}

func (p *Player) GetPosition() *Position {
	return &p.pos
}

func (p *Player) SetPosition(pos *Position) {
	p.pos = *pos
}

func (p *Player) Move(dx, dy int) {
	p.pos.X += dx
	p.pos.Y += dy
}

func (p *Player) MoveTo(pos *Position) {
	p.pos = *pos
}

func (p *Player) Draw() {
	gd := CurrentGame().Data
	firefly.DrawImage(*p.Image, firefly.Point{X: p.pos.X * gd.TileWidth, Y: p.pos.Y * gd.TileHeight})
}
