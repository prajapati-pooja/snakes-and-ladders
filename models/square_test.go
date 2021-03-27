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

	actualNextPosition := regularSquare.getNextPosition()

	expectedNextPosition := 6
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}

func TestRegularSquare_EnterShouldAssignTheGivenPlayer(t *testing.T) {
	regularSquare := RegularSquare{
		position: 1,
		board: &Board{currentMoves: 5},
	}

	regularSquare.enter(Player{Name: "player 1"})
	expectedPlayer := Player{Name: "player 1"}
	assert.Equal(t, expectedPlayer, regularSquare.player)
}

func TestRegularSquare_LeaveShouldAssignEmptyPlayer(t *testing.T) {
	regularSquare := RegularSquare{
		position: 1,
		board: &Board{currentMoves: 5},
	}

	regularSquare.leave()
	expectedPlayer := Player{}
	assert.Equal(t, expectedPlayer, regularSquare.player)
}