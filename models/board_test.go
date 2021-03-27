package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_FirstSquareShouldReturnFirstSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	actualSquare :=  board.firstSquare()

	expectedSquare := &RegularSquare{position: 1, board: board}
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
	board.currentSquare =  &RegularSquare{position: 2, board: board}
	player := Player{Name: "player 1"}
	board.findNextSquare(player, 4)

	expectedSquare := &RegularSquare{position: 6, board: board, player: player}
	assert.Equal(t, expectedSquare, board.currentSquare)
}

func TestBoard_FindNextSquareShouldReturnTailSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	board.currentSquare = &RegularSquare{position: 10, board: board}
	player := Player{Name: "player 1"}

	board.findNextSquare(player, 4)

	expectedSquare := &SnakeSquare{
		snake: Snake{Head: 14, Tail: 7},
		player: player,
	}
	assert.Equal(t, expectedSquare, board.currentSquare)
}