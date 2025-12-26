package tinyrogue

// TurnState represents the current state of the game turn.
type TurnState int

const (
	BeforePlayerAction = iota
	PlayerTurn
	CreatureTurn
	GameOver
)

// GetNextState returns the next turn state based on the current state.
func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return CreatureTurn
	case CreatureTurn:
		return BeforePlayerAction
	case GameOver:
		return GameOver
	default:
		return PlayerTurn
	}
}
