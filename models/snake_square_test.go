package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnakeSquare_ShouldGetNextPosition(t *testing.T) {
	snakeSquare := SnakeSquare{
		snake: Snake{Head: 14, Tail: 8},
	}

	actualNextPosition := snakeSquare.getNextPosition()

	expectedNextPosition := 8
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