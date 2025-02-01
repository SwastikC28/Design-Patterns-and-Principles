package model

import (
	"errors"
	"fmt"
)

type Board interface {
	MarkPosition(int, int, *Player) error
	CheckWinner(row, col int) bool
	ShowBoard()
}

type boardImpl struct {
	grid [][]*Piece
}

func NewBoard(size int) Board {
	grid := make([][]*Piece, size)
	for i := range grid {
		grid[i] = make([]*Piece, size)
	}
	return &boardImpl{grid: grid}
}

func (board *boardImpl) MarkPosition(row, col int, player *Player) error {
	if row < 0 || row >= len(board.grid) || col < 0 || col >= len(board.grid[0]) {
		return errors.New("invalid move")
	}

	if board.grid[row][col] != nil {
		return errors.New("invalid move")
	}

	board.grid[row][col] = player.piece
	return nil
}

func (board *boardImpl) CheckWinner(row, col int) bool {
	win := false
	if win = board.verticalCheck(row, col); win {
		fmt.Println("VERTICAL WIN")
		return true
	}

	if win = board.horizontalCheck(row, col); win {
		fmt.Println("HORIZONONTAL WIN")
		return true
	}

	if win = board.leftDiagonalCheck(row, col); win {
		fmt.Println("leftDiagonalCheck WIN")
		return true
	}

	if win = board.rightDiagonalCheck(row, col); win {
		fmt.Println("rightDiagonalCheck WIN")
		return true
	}

	return false
}

func (board *boardImpl) ShowBoard() {
	fmt.Println("Displaying Board")
	fmt.Println()

	m := len(board.grid)
	n := len(board.grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			piece := board.grid[i][j]
			if piece == nil {
				fmt.Printf("%s \t", "-")
				continue
			}
			fmt.Printf("%s \t", piece.pieceType)
		}

		fmt.Println()
	}

	fmt.Println()
}

func (board *boardImpl) horizontalCheck(row, col int) bool {
	piece := board.grid[row][col]

	var horizontalWinner = true
	for i := 0; i < len(board.grid[row]); i++ {
		if board.grid[row][i] != piece {
			horizontalWinner = false
			break
		}
	}

	return horizontalWinner
}

func (board *boardImpl) verticalCheck(row, col int) bool {
	piece := board.grid[row][col]

	var verticalWinner = true
	for i := 0; i < len(board.grid); i++ {
		if board.grid[i][col] != piece {
			verticalWinner = false
			break
		}
	}

	return verticalWinner
}

func (board *boardImpl) leftDiagonalCheck(row, col int) bool {
	if row != col {
		return false
	}
	piece := board.grid[row][col]
	for i := 0; i < len(board.grid); i++ {
		if board.grid[i][i] != piece {
			return false
		}
	}
	return true
}

func (board *boardImpl) rightDiagonalCheck(row, col int) bool {
	if row+col != len(board.grid)-1 {
		return false
	}
	piece := board.grid[row][col]
	for i, j := 0, len(board.grid)-1; i < len(board.grid); i, j = i+1, j-1 {
		if board.grid[i][j] != piece {
			return false
		}
	}
	return true
}
