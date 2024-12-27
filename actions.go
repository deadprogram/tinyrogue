package tinyrogue

type Actionable interface {
	Action(sender Character, receiver Character)
}

type DebugAction struct {
}

func (da *DebugAction) Action(sender Character, receiver Character) {
	logDebug("debug action")
}
