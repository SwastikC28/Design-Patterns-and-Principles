package model

import "errors"

var list map[string]bool = make(map[string]bool)

type Piece struct {
	pieceType string
}

func NewPiece(pieceType string) (*Piece, error) {
	_, exists := list[pieceType]
	if exists {
		return nil, errors.New("piece type already exists")
	}

	piece := &Piece{pieceType: pieceType}
	list[piece.pieceType] = true
	return piece, nil
}
