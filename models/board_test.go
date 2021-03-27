package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_FirstSquareShouldReturnFirstSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	actualSquare :=  board.firstSquare()

	expectedSquare := RegularSquare{position: 1, board: board}
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestBoard_FirstSquareShouldNotReturnOtherSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	actualSquare :=  board.firstSquare()

	expectedSquare := RegularSquare{position: 2, board: board}
	assert.NotEqual(t, expectedSquare, actualSquare)
}

func TestBoard_FindNextSquareShouldReturnRelativeNextSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	firstSquare := RegularSquare{position: 2, board: board}
	actualSquare :=  board.findNextSquare(firstSquare, 4)

	expectedSquare := RegularSquare{position: 6, board: board}
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestBoard_FindNextSquareShouldReturnTailSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	firstSquare := RegularSquare{position: 10, board: board}
	actualSquare :=  board.findNextSquare(firstSquare, 4)

	expectedSquare := SnakeSquare{
		Snake: Snake{Head: 14, Tail: 7},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}