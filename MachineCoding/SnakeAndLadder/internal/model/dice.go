package model

import (
	"errors"
	"math/rand"
)

type Dice interface {
	RollDice() int
}

func NewDice(sides int) (Dice, error) {
	switch sides {
	case 4:
		return &TetraHedronDice{}, nil
	case 6:
		return &CubicDice{}, nil
	case 8:
		return &OctaHedronDice{}, nil
	default:
		return nil, errors.New("invalid dice sides")
	}
}

type CubicDice struct{}

func (dice *CubicDice) RollDice() int {
	num := rand.Intn(6) + 1
	return num
}

type TetraHedronDice struct{}

func (dice *TetraHedronDice) RollDice() int {
	num := rand.Intn(4) + 1
	return num
}

type OctaHedronDice struct{}

func (dice *OctaHedronDice) RollDice() int {
	num := rand.Intn(8) + 1
	return num
}
