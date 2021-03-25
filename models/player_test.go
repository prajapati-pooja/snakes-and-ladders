package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_ShouldSetCurrentPositionAsFirstSquareOnGameStart(t *testing.T) {
	player := Player{
		Board: NewBoard(100, Snake{Head: 14, Tail: 7}),
	}
	player.StartGame()

	expectedPosition := RegularSquare{Position: 1}
	assert.Equal(t, expectedPosition, player.Position)
}

func TestPlayer_ShouldMoveToRelativeSquare(t *testing.T) {
	player := Player{
		Board: NewBoard(100, Snake{Head: 14, Tail: 7}),
	}
	player.StartGame()
	player.Move(5)

	expectedPosition := RegularSquare{Position: 6}

	assert.Equal(t, expectedPosition, player.Position)
}