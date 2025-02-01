package model

type Player struct {
	name  string
	piece *Piece
}

func NewPlayer(name string, piece *Piece) *Player {
	return &Player{
		name:  name,
		piece: piece,
	}
}

func (player *Player) GetName() string {
	return player.name
}
