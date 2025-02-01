package model

type WinChecker interface {
	IsWinningMove(board Board, row, col int, piece *Piece) bool
}

type TicTacToeWinChecker struct{}

func (t *TicTacToeWinChecker) IsWinningMove(board Board, row, col int, piece *Piece) bool {
	return checkDirection(board, row, col, piece, 1, 0, 3) || // Horizontal
		checkDirection(board, row, col, piece, 0, 1, 3) || // Vertical
		checkDirection(board, row, col, piece, 1, 1, 3) || // Diagonal ↘
		checkDirection(board, row, col, piece, 1, -1, 3) // Diagonal ↙
}

func checkDirection(board Board, row, col int, piece *Piece, dx, dy, winLength int) bool {
	count := 1
	size := len(board.GetGrid()) // Assuming Board has a method GetGrid() to return the grid

	// Check forward
	for i := 1; i < winLength; i++ {
		r, c := row+dx*i, col+dy*i
		if r < 0 || r >= size || c < 0 || c >= size || board.GetGrid()[r][c] != piece {
			break
		}
		count++
	}

	// Check backward
	for i := 1; i < winLength; i++ {
		r, c := row-dx*i, col-dy*i
		if r < 0 || r >= size || c < 0 || c >= size || board.GetGrid()[r][c] != piece {
			break
		}
		count++
	}

	return count >= winLength
}
