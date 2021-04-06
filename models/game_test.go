package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


type dummyDice struct {
}

func (dd dummyDice) Throw() int {
	return 5
}

func TestGame_StartShouldSetTheFirstSquareAsCurrentSquare(t *testing.T) {
	snake := Snake{
		Head: 7,
		Tail: 14,
	}
	board := NewBoard(100, snake, Ladder{Top: 35, Bottom: 11})
	player := Player{
		Name: "Player 1",
	}

	game := NewGame(board, player)

	game.Start()
	expectedSquare := &RegularSquare{
		position: 1,
		player:   player,
	}
	assert.Equal(t, expectedSquare, game.board.currentSquare)
}

func TestGame_PlayShouldUpdateTheCurrentSquareWithNextSquare(t *testing.T) {
	snake := Snake{
		Head: 7,
		Tail: 14,
	}
	board := NewBoard(100, snake, Ladder{Top: 35, Bottom: 11})
	player := Player{
		Name: "Player 1",
	}
	dice := dummyDice{}
	game := NewGame(board, player)

	game.Start()
	game.Play(dice)
	expectedSquare := &RegularSquare{
		position: 6,
		player:   player,
	}
	assert.Equal(t, expectedSquare, game.board.currentSquare)
}