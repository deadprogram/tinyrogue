package tinyrogue

type Player struct {
	*Character
}

func NewPlayer() *Player {
	return &Player{
		Character: &Character{},
	}
}
