package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare_Snake(t *testing.T) {
	snakeSquare := SnakeSquare{
		Snake: Snake{Head: 14, Tail: 8},
	}

	actualNextPosition := snakeSquare.Shift()

	expectedNextPosition := 8
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}