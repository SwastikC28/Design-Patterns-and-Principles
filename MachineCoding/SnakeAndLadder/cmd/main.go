package main

import (
	"context"
	"fmt"
	"os"
	"snakeandladder/internal/model"

	"github.com/rs/zerolog"
)

func init() {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stdout,
	}).With().Timestamp().Logger()

	zerolog.DefaultContextLogger = &logger
}

func main() {
	ctx := context.Background()

	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Logger Initialized")

	players := []*model.Player{
		model.NewPlayer("PlayerX"),
		model.NewPlayer("PlayerO"),
	}

	entities := []model.Entity{
		model.NewSnake(50, 10),
		model.NewSnake(30, 20),
		model.NewSnake(65, 15),
	}

	dice1, err := model.NewDice(6)
	if err != nil {
		logger.Err(err).Msg("Error creating dice")
		return
	}

	dice2, err := model.NewDice(6)
	if err != nil {
		logger.Err(err).Msg("Error creating dice")
		return
	}

	dices := []model.Dice{dice1, dice2}

	board := model.NewBoard(entities...)
	winchecker := model.NewWinChecker()

	game := model.NewGame(board, winchecker, dices, players)
	game.Start()

	winner := game.GetWinner()
	fmt.Println("Winner of the game is ", winner.Name)
}
