package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_ShouldSetCurrentPositionAsFirstSquareOnGameStart(t *testing.T) {
	board := NewBoard(100, Snake{Head: 14, Tail: 7})
	player := Player{
		Board: board,
	}
	player.StartGame()

	expectedPosition := RegularSquare{position: 1, board: board}
	assert.Equal(t, expectedPosition, player.Position)
}

func TestPlayer_ShouldMoveToRelativeSquare(t *testing.T) {
	board := NewBoard(100, Snake{Head: 14, Tail: 7})
	player := Player{
		Board: board,
	}
	player.StartGame()
	player.Move(5)

	expectedPosition := RegularSquare{position: 6, board: board}

	assert.Equal(t, expectedPosition, player.Position)
}