package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getSquareShouldReturnRegularSquare(t *testing.T) {
	board := &Board{}
	actualSquare :=  getSquare(2, Snake{Head: 14, Tail: 7}, board)
	expectedSquare := &RegularSquare{position: 2, board: board}
	assert.Equal(t, expectedSquare, actualSquare)
}

func Test_getSquareShouldReturnSnakeSquare(t *testing.T) {
	board := &Board{}
	actualSquare :=  getSquare(14, Snake{Head: 14, Tail: 7}, board)
	expectedSquare := &SnakeSquare{
		snake: Snake{Head: 14, Tail: 7},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}

