package tinyrogue

import "github.com/firefly-zero/firefly-go/firefly"

// Portal is the type for all portals between levels in the game.
type Portal struct {
	PortalType      string
	Visible         bool
	Image           *firefly.Image
	DungeonName     string
	DestinationName string
}

// NewPortal creates a new Portal and initializes the data.
func NewPortal(pt string, img *firefly.Image, dungeon *Dungeon, destination *Level) *Portal {
	return &Portal{
		PortalType:      pt,
		Visible:         true,
		Image:           img,
		DungeonName:     dungeon.Name,
		DestinationName: destination.Name,
	}
}

func (p *Portal) Dungeon() *Dungeon {
	logDebug("Portal.Dungeon: " + p.DungeonName)
	return CurrentGame().Map.Dungeon(p.DungeonName)
}

func (p *Portal) Destination() *Level {
	logDebug("Portal.Destination: " + p.DestinationName)
	return CurrentGame().Map.Dungeon(p.DungeonName).Level(p.DestinationName)
}
