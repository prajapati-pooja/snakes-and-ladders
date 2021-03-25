package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_InitializeSquareShouldReturnRegularSquare(t *testing.T) {
	actualSquare :=  initializeSquare(2, Snake{Head: 14, Tail: 7})
	expectedSquare := RegularSquare{Position: 2}
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestBoard_InitializeSquareShouldReturnSnakeSquare(t *testing.T) {
	actualSquare :=  initializeSquare(14, Snake{Head: 14, Tail: 7})
	expectedSquare := SnakeSquare{
		Snake: Snake{Head: 14, Tail: 7},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestBoard_FirstSquareShouldReturnFirstSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	actualSquare :=  board.firstSquare()

	expectedSquare := RegularSquare{Position: 1}
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestBoard_FirstSquareShouldNotReturnOtherSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	actualSquare :=  board.firstSquare()

	expectedSquare := RegularSquare{Position: 2}
	assert.NotEqual(t, expectedSquare, actualSquare)
}

func TestBoard_FindNextSquareShouldReturnRelativeNextSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	firstSquare := RegularSquare{Position: 2}
	actualSquare :=  board.findNextSquare(firstSquare, 4)

	expectedSquare := RegularSquare{Position: 6}
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestBoard_FindNextSquareShouldReturnTailSquare(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7})

	firstSquare := RegularSquare{Position: 10}
	actualSquare :=  board.findNextSquare(firstSquare, 4)

	expectedSquare := SnakeSquare{
		Snake: Snake{Head: 14, Tail: 7},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}