package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare_Regular(t *testing.T) {
	regularSquare := RegularSquare{
		position: 1,
		board: &Board{currentMoves: 5},
	}

	actualNextPosition := regularSquare.Shift()

	expectedNextPosition := 6
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}