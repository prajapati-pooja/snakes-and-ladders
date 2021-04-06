package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetSquareShouldReturnRegularSquare(t *testing.T) {
	actualSquare :=  getSquare(2, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})
	expectedSquare := &RegularSquare{position: 2}
	assert.Equal(t, expectedSquare, actualSquare)
}

func Test_GetSquareShouldReturnSnakeSquare(t *testing.T) {
	actualSquare :=  getSquare(14, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})
	expectedSquare := &SnakeSquare{
		position: 14,
		snake: Snake{Head: 14, Tail: 7},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}

func Test_GetSquareShouldReturnLadderSquare(t *testing.T) {
	actualSquare :=  getSquare(11, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})
	expectedSquare := &LadderSquare{
		position: 11,
		ladder: Ladder{Top: 35, Bottom:11},
	}
	assert.Equal(t, expectedSquare, actualSquare)
}
