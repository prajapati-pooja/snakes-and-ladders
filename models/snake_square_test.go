package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnakeSquare_GetNextPositionShouldMoveThePlayerToLowerPosition(t *testing.T) {
	snakeSquare := SnakeSquare{
		snake: Snake{Head: 14, Tail: 8},
	}

	actualNextPosition, err := snakeSquare.getNextPosition()

	expectedNextPosition := 8
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedNextPosition, actualNextPosition)
}

func TestSnakeSquare_EnterShouldAssignTheGivenPlayer(t *testing.T) {
	snakeSquare := SnakeSquare{
		snake: Snake{Head: 14, Tail: 8},
	}

	snakeSquare.enter(Player{Name: "player 1"})

	expectedPlayer := Player{Name: "player 1"}
	assert.Equal(t, expectedPlayer, snakeSquare.player)
}

func TestSnakeSquare_LeaveShouldAssignTheEmptyPlayer(t *testing.T) {
	snakeSquare := SnakeSquare{
		snake: Snake{Head: 14, Tail: 8},
	}

	snakeSquare.leave()

	expectedPlayer := Player{}
	assert.Equal(t, expectedPlayer, snakeSquare.player)
}

func TestSnakeSquare_StringShouldReturnFormattedValueWithPlayer(t *testing.T) {
	snakeSquare := SnakeSquare{
		position: 14,
		snake: Snake{Head: 14, Tail: 8},
		player: Player{Name: "player 1"},
	}

	expectedValue := "{position: 14, player: {Name:player 1}, snake: {Head:14 Tail:8}}"
	assert.Equal(t, expectedValue, snakeSquare.String())
}

func TestSnakeSquare_StringShouldReturnFormattedValueWithoutPlayer(t *testing.T) {
	snakeSquare := SnakeSquare{
		position: 14,
		snake: Snake{Head: 14, Tail: 8},
	}

	expectedValue := "{position: 14, snake: {Head:14 Tail:8}}"
	assert.Equal(t, expectedValue, snakeSquare.String())
}