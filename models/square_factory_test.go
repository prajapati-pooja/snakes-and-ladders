package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetSquareShouldReturnRegularSquare(t *testing.T) {
	board := &Board{}
	actualSquare :=  getSquare(2, Snake{Head: 14, Tail: 7}, board)
	expectedSquare := &RegularSquare{position: 2, board: board}
	assert.Equal(t, expectedSquare, actualSquare)
}

func Test_GetSquareShouldReturnSnakeSquare(t *testing.T) {
	board := &Board{}
	actualSquare :=  getSquare(14, Snake{Head: 14, Tail: 7}, board)
	expectedSquare := &SnakeSquare{
		position: 14,
		snake: Snake{Head: 14, Tail: 7},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}

