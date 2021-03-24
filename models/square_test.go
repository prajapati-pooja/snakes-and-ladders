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