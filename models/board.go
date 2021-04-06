package models

import "fmt"

const GridStartPosition = 1

type Board struct {
	squares       []Square
	currentMove   int
	currentSquare Square
}

func NewBoard(maxGridSize int, snake Snake, ladder Ladder) *Board {
	board := &Board{}
	var position int
	var squares []Square
	for position = GridStartPosition; position <= maxGridSize; position++ {
		square := getSquare(position, snake, ladder)
		squares = append(squares, square)
	}
	board.squares = squares
	return board
}

func (b *Board) registerInFirstSquare(player Player) {
	b.currentSquare = b.squares[0]
	b.currentSquare.enter(player)
}

func (b *Board) setNextSquare(player Player, moves int) {
	b.currentMove = moves
	nextPosition, err  := b.currentSquare.getNextPosition(moves)
	if err != nil {
		return
	}
	b.currentSquare.leave()
	nextSquare := b.squares[nextPosition-1]
	nextSquare.enter(player)
	b.currentSquare = nextSquare
}

func (b *Board)String() string {
	return fmt.Sprintf("%+v", b.squares)
}
