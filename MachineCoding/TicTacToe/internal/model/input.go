package model

import (
	"errors"
	"fmt"
)

type InputHandler interface {
	GetMove(player *Player, board Board) (int, int, error)
}

type CLIInputHandler struct {
}

func (cli *CLIInputHandler) GetMove(player *Player, board Board) (int, int, error) {
	var position int
	fmt.Printf("%s, enter the position: ", player.GetName())
	_, err := fmt.Scanln(&position)
	if err != nil {
		return -1, -1, errors.New("invalid input")
	}

	boardSize := len(board.GetGrid())
	position--
	row, col := position/boardSize, position%boardSize

	return row, col, nil
}
