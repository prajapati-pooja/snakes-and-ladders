package models

import "fmt"

type Board struct {
	squares       []Square
	currentMoves  int
	currentSquare Square
}

func NewBoard(maxGridSize int, snake Snake) *Board {
	board := &Board{}
	var position int
	var squares []Square
	for position = 1; position <= maxGridSize; position++ {
		square := getSquare(position, snake, board)
		squares = append(squares, square)
	}
	board.squares = squares
	return board
}

func (b *Board) firstSquare() Square {
	return b.squares[0]
}

func (b *Board) enterInFirstSquare(player Player) {
	b.currentSquare = b.squares[0]
	b.currentSquare.enter(player)
}

func (b *Board) findNextSquare(player Player, moves int) {
	b.currentMoves = moves
	b.currentSquare.leave()
	nextPosition := b.currentSquare.getNextPosition()
	nextSquare := b.squares[nextPosition-1]
	nextSquare.enter(player)
	b.currentSquare = nextSquare
}

func (b *Board) currentMove() int {
	return b.currentMoves
}

func (b *Board) String() string {
	return fmt.Sprintf("%+v", b.squares)
}
