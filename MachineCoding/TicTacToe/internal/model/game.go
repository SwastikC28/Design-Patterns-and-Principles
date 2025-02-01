package model

import "fmt"

type Game interface {
	Start() *Player
	GetWinner() *Player
}

type gameImpl struct {
	boardSize int
	board     Board
	players   []*Player
	finished  bool
	winner    *Player
}

func NewGame(boardSize int, players ...*Player) Game {
	game := &gameImpl{
		board:     NewBoard(boardSize),
		boardSize: boardSize,
		players:   make([]*Player, 0),
		finished:  false,
		winner:    nil,
	}

	game.players = append(game.players, players...)
	return game
}

func (game *gameImpl) Start() *Player {
	if game.finished {
		return game.winner
	}

	for turn := 0; turn < game.boardSize*game.boardSize; turn++ {
		currentPlayer := game.players[turn%len(game.players)]

		var row, col int
		fmt.Printf("TURN %d: %s, enter row and column: ", turn+1, currentPlayer.GetName())
		fmt.Scanln(&row, &col)

		row, col = row-1, col-1

		err := game.board.MarkPosition(row, col, currentPlayer)
		if err != nil {
			fmt.Println("Invalid Move")
			turn-- // Retry the same player
			continue
		}

		game.board.ShowBoard()

		if game.board.CheckWinner(row, col) {
			fmt.Println(currentPlayer.GetName(), "won the game")
			game.winner = currentPlayer
			return currentPlayer
		}
	}

	fmt.Println("Game Over")
	return nil
}

func (game *gameImpl) GetWinner() *Player {
	return game.winner
}
