package tinyrogue

type TurnState int

const (
	BeforePlayerAction = iota
	PlayerTurn
	CreatureTurn
	GameOver
)

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
