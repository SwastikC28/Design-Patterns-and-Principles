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

	game := model.NewGame(boardSize, player1, player2)
	winner := game.Start()
	if winner != nil {
		fmt.Println("Winner of the game is", winner.GetName())
	}
}
