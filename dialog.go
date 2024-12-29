package tinyrogue

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

const minimumDialogDelay = 60

// Dialog is a simple dialog box that can be displayed to the player.
type Dialog struct {
	Font              *firefly.Font
	FontColor         firefly.Color
	FillColor         firefly.Color
	Text1             string
	Text2             string
	NeedsConfirmation bool
	Confirmed         bool
	delay             int
	Point             firefly.Point
	Size              firefly.Size
}

// NewDialog creates a new dialog box with the given text and font.
func NewDialog(text1 string, text2 string, font *firefly.Font, fontcolor, fillColor firefly.Color, needsConfirmation bool) *Dialog {
	return &Dialog{
		Text1:             text1,
		Text2:             text2,
		Font:              font,
		NeedsConfirmation: needsConfirmation,
		Confirmed:         false,
		FontColor:         fontcolor,
		FillColor:         fillColor,
		Point:             firefly.Point{X: 0, Y: 20},
		Size:              firefly.Size{W: 260, H: 60},
	}
}

// Update updates the dialog box, basically just used to dismiss it.
func (d *Dialog) Update() {
	d.delay++
	if d.delay > minimumDialogDelay {
		if d.NeedsConfirmation {
			buttons := firefly.ReadButtons(firefly.Combined)
			if buttons.N || buttons.S || buttons.E || buttons.W {
				d.Confirmed = true
			}
		}
	}
}

const (
	xMargin     = 10
	yMargin     = 20
	lineSpacing = 10
)

// Draw draws the dialog box to the screen.
func (d *Dialog) Draw() {
	firefly.DrawRect(d.Point, d.Size, firefly.Style{FillColor: d.FillColor})

	x := d.Point.X + xMargin
	y := d.Point.Y + yMargin

	firefly.DrawText(d.Text1, *d.Font, firefly.Point{X: x, Y: y}, d.FontColor)

	y += lineSpacing
	if d.Text2 != "" {
		firefly.DrawText(d.Text2, *d.Font, firefly.Point{X: x, Y: y}, d.FontColor)
		y += lineSpacing
	}

	if d.NeedsConfirmation {
		firefly.DrawText("Press any button to continue...", *d.Font, firefly.Point{X: x, Y: y}, d.FontColor)
	}
}
