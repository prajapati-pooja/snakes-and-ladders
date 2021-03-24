package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare_Regular(t *testing.T) {
	regularSquare := RegularSquare{
		Position: 1,
	}

	actualNextPosition := regularSquare.Shift(5)

	expectedNextPosition := 6
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}

func TestSquare_Snake(t *testing.T) {
	snakeSquare := SnakeSquare{
		Snake: Snake{head: 14, tail: 8},
	}

	actualNextPosition := snakeSquare.Shift(5)

	expectedNextPosition := 8
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}