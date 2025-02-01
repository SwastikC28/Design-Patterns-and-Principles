package model

type Board interface {
	MakeMove(*Player, int)
	GetBoard() map[int]Entity
}

type SimpleBoard struct {
	board map[int]Entity
}

func NewBoard(entities ...Entity) Board {
	board := &SimpleBoard{
		board: make(map[int]Entity),
	}

	for _, entity := range entities {
		if entity.GetStartPosition() == 100 {
			continue
		}
		board.board[entity.GetStartPosition()] = entity
	}

	return board
}

func (board *SimpleBoard) MakeMove(player *Player, steps int) {
	currentPosition := player.Position + steps

	if currentPosition > 100 {
		return
	}

	entity, exists := board.board[currentPosition]

	if exists {
		currentPosition = entity.GetEndPosition()
	}

	player.Position = currentPosition
}

func (board *SimpleBoard) GetBoard() map[int]Entity {
	return board.board
}
