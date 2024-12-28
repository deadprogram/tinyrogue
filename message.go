package tinyrogue

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

const defaultMessageDelay = 60

type Message struct {
	Font              *firefly.Font
	FontColor         firefly.Color
	FillColor         firefly.Color
	Text1             string
	Text2             string
	NeedsConfirmation bool
	Confirmed         bool
	delay             int
}

func NewMessage(text string, font *firefly.Font, fontcolor, fillColor firefly.Color, needsConfirmation bool) *Message {
	return &Message{
		Text1:             text,
		Font:              font,
		NeedsConfirmation: needsConfirmation,
		Confirmed:         false,
		FontColor:         fontcolor,
		FillColor:         fillColor,
	}
}

func (m *Message) Update() {
	m.delay++
	if m.delay > defaultMessageDelay {
		if m.NeedsConfirmation {
			buttons := firefly.ReadButtons(firefly.Combined)
			if buttons.N || buttons.S || buttons.E || buttons.W {
				m.Confirmed = true
			}
		}
	}
}

func (m *Message) Draw() {
	pt := firefly.Point{X: 0, Y: 20}
	sz := firefly.Size{W: 260, H: 60}
	firefly.DrawRect(pt, sz, firefly.Style{FillColor: m.FillColor, StrokeColor: firefly.ColorBlack})

	firefly.DrawText(m.Text1, *m.Font, firefly.Point{X: 10, Y: 40}, firefly.ColorRed)

	y := 50
	if m.Text2 != "" {
		firefly.DrawText(m.Text2, *m.Font, firefly.Point{X: 10, Y: y}, firefly.ColorRed)
		y = 60
	}

	if m.NeedsConfirmation {
		firefly.DrawText("Press any button to continue...", *m.Font, firefly.Point{X: 10, Y: y}, firefly.ColorRed)
	}
}
