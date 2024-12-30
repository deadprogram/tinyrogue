package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Portal is the type for all portals between levels in the game.
type Portal struct {
	PortalType  string
	Image       *firefly.Image
	Destination *Level
}

// NewPortal creates a new Portal and initializes the data.
func NewPortal(pt string, img *firefly.Image, destination *Level) *Portal {
	return &Portal{
		PortalType:  pt,
		Image:       img,
		Destination: destination,
	}
}
