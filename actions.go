package tinyrogue

// Actionable is an interface for actions that can be taken by characters.
type Actionable interface {
	Action(sender Character, receiver Character)
}

type DebugAction struct {
}

func (da *DebugAction) Action(sender Character, receiver Character) {
	logDebug("debug action")
}
