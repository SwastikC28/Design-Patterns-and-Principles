package main

import (
	"fmt"
	"tictactoe/internal/model"
)

func main() {
	fmt.Println("Tic tac toe")

	pieceX, err := model.NewPiece("X")
	if err != nil {
		return
	}

	pieceY, err := model.NewPiece("O")
	if err != nil {
		return
	}

	player1 := model.NewPlayer("PlayerX", pieceX)
	player2 := model.NewPlayer("PlayerY", pieceY)

	boardSize := 3

	var winChecker model.WinChecker
	switch boardSize {
	case 3:
		winChecker = &model.TicTacToeWinChecker{}
	default:
		fmt.Println("Invalid Board Size")
		return
	}

	// Open Closed Principle
	choice := "CLI"
	var inputHandler model.InputHandler
	switch choice {
	case "CLI":
		inputHandler = &model.CLIInputHandler{}
	default:
		fmt.Println("Invalid Input Handler")
		return
	}

	game := model.NewGame(boardSize, winChecker, inputHandler, player1, player2)
	winner := game.Start()
	if winner != nil {
		fmt.Println("Winner of the game is", winner.GetName())
	}
}
