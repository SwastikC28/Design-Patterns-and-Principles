package model

import (
	"context"

	"github.com/rs/zerolog"
)

type Game interface {
	Start()
	GetWinner() *Player
}

type gameImpl struct {
	dices      []Dice
	board      Board
	players    []*Player
	winner     *Player
	winchecker WinChecker
}

func NewGame(board Board, winchecker WinChecker, dice []Dice, players []*Player) Game {
	game := &gameImpl{
		dices:      make([]Dice, 0),
		board:      board,
		players:    make([]*Player, 0),
		winner:     nil,
		winchecker: winchecker,
	}

	game.dices = append(game.dices, dice...)
	game.players = append(game.players, players...)
	return game
}

func (game *gameImpl) Start() {
	logger := zerolog.Ctx(context.Background()).With().Str("module", "game").Logger()

	i := 0
	for game.winner == nil {
		i = i % len(game.players)
		currentPlayer := game.players[i]
		initialPosition := currentPlayer.Position

		steps := 0

		for _, dice := range game.dices {
			steps += dice.RollDice()
		}

		game.board.MakeMove(currentPlayer, steps)

		finalPosition := currentPlayer.Position

		finished := game.winchecker.IsWinningMove(currentPlayer)

		logger.Info().Msgf("%s rolled a %d and moved from %d to %d.", currentPlayer.Name, steps, initialPosition, finalPosition)

		if finished {
			game.winner = currentPlayer
			return
		}

		i++
	}
}

func (game *gameImpl) GetWinner() *Player {
	return game.winner
}
