package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadderSquare_GetNextPositionShouldMoveThePlayerToHigherPosition(t *testing.T) {
	ladderSquare := LadderSquare{
		ladder: Ladder{Top: 35, Bottom: 11},
		player: Player{Name: "player 1"},
	}

	actualNextPosition, err := ladderSquare.getNextPosition(0)

	expectedNextPosition := 35
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}

func TestLadderSquare_EnterShouldAssignTheGivenPlayer(t *testing.T) {
	ladderSquare := LadderSquare{
		ladder: Ladder{Top: 35, Bottom: 11},
		player: Player{Name: "player 1"},
	}
	ladderSquare.enter(Player{Name: "player 1"})

	expectedPlayer := Player{Name: "player 1"}
	assert.Equal(t, expectedPlayer, ladderSquare.player)
}

func TestLadderSquare_LeaveShouldAssignTheEmptyPlayer(t *testing.T) {
	ladderSquare := LadderSquare{
		ladder: Ladder{Top: 35, Bottom: 11},
		player: Player{Name: "player 1"},
	}

	ladderSquare.leave()

	expectedPlayer := Player{}
	assert.Equal(t, expectedPlayer, ladderSquare.player)
}

func TestLadderSquare_StringShouldReturnFormattedValueWithPlayer(t *testing.T) {
	ladderSquare := LadderSquare{
		position: 11,
		ladder: Ladder{Top: 35, Bottom: 11},
		player: Player{Name: "player 1"},
	}

	expectedValue := "{position: 11, player: {Name:player 1}, ladder: {Top:35 Bottom:11}}"
	assert.Equal(t, expectedValue, ladderSquare.String())
}

func TestLadderSquare_StringShouldReturnFormattedValueWithoutPlayer(t *testing.T) {
	ladderSquare := LadderSquare{
		position: 11,
		ladder: Ladder{Top: 35, Bottom: 11},
	}

	expectedValue := "{position: 11, ladder: {Top:35 Bottom:11}}"
	assert.Equal(t, expectedValue, ladderSquare.String())
}