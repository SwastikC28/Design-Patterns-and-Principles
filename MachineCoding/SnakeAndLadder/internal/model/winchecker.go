package model

type WinChecker interface {
	IsWinningMove(*Player) bool
}

func NewWinChecker() WinChecker {
	return &SimpleWinChecker{}
}

type SimpleWinChecker struct{}

func (winchecker *SimpleWinChecker) IsWinningMove(player *Player) bool {
	return player.Position == 100
}
