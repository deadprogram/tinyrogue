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

// Draw draws the dialog box to the screen.
func (d *Dialog) Draw() {
	pt := firefly.Point{X: 0, Y: 20}
	sz := firefly.Size{W: 260, H: 60}
	firefly.DrawRect(pt, sz, firefly.Style{FillColor: d.FillColor})

	firefly.DrawText(d.Text1, *d.Font, firefly.Point{X: 10, Y: 40}, d.FontColor)

	y := 50
	if d.Text2 != "" {
		firefly.DrawText(d.Text2, *d.Font, firefly.Point{X: 10, Y: y}, d.FontColor)
		y = 60
	}

	if d.NeedsConfirmation {
		firefly.DrawText("Press any button to continue...", *d.Font, firefly.Point{X: 10, Y: y}, d.FontColor)
	}
}
