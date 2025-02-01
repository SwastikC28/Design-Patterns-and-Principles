package model

import "fmt"

type Game interface {
	Start() *Player
	GetWinner() *Player
}

type gameImpl struct {
	boardSize    int
	board        Board
	players      []*Player
	finished     bool
	winner       *Player
	winchecker   WinChecker
	inputHandler InputHandler
}

func NewGame(boardSize int, winChecker WinChecker, inputHandler InputHandler, players ...*Player) Game {
	game := &gameImpl{
		board:        NewBoard(boardSize),
		boardSize:    boardSize,
		players:      make([]*Player, 0),
		finished:     false,
		winner:       nil,
		winchecker:   winChecker,
		inputHandler: inputHandler,
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
		var err error
		for {
			row, col, err = game.inputHandler.GetMove(currentPlayer, game.board)
			if err != nil {
				fmt.Println("Invalid input, try again.")
				continue
			}

			if err := game.board.MarkPosition(row, col, currentPlayer); err != nil {
				fmt.Println(err)
				continue
			}
			break
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
