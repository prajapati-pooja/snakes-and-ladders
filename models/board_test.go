package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_FindNextSquareShouldMoveThePlayerToHighPosition(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})
	board.currentSquare =  &RegularSquare{position: 2}
	player := Player{Name: "player 1"}
	board.setNextSquare(player, 4)

	expectedSquare := &RegularSquare{position: 6, player: player}
	assert.Equal(t, expectedSquare, board.currentSquare)
}

func TestBoard_FindNextSquareShouldNotMoveThePlayerIfNextPositionIsHigherThanGridSize(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})
	board.currentSquare =  &RegularSquare{position: 98}
	player := Player{Name: "player 1"}
	board.setNextSquare(player, 4)

	expectedSquare := &RegularSquare{position: 98}
	assert.Equal(t, expectedSquare, board.currentSquare)
}

func TestBoard_FindNextSquareShouldMoveThePlayerToLowerPosition(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})

	board.currentSquare = &RegularSquare{position: 10}
	player := Player{Name: "player 1"}

	board.setNextSquare(player, 4)

	expectedSquare := &SnakeSquare{
		position: 14,
		snake: Snake{Head: 14, Tail: 7},
		player: player,
	}
	assert.Equal(t, expectedSquare, board.currentSquare)
}

func TestBoard_StringShouldReturnFormattedValue(t *testing.T) {
	board  := NewBoard(100, Snake{Head: 14, Tail: 7}, Ladder{Top: 35, Bottom: 11})

	expectedString := `[{position: 1} {position: 2} {position: 3} {position: 4} {position: 5} {position: 6} {position: 7} {position: 8} {position: 9} {position: 10} {position: 11, ladder: {Top:35 Bottom:11}} {position: 12} {position: 13} {position: 14, snake: {Head:14 Tail:7}} {position: 15} {position: 16} {position: 17} {position: 18} {position: 19} {position: 20} {position: 21} {position: 22} {position: 23} {position: 24} {position: 25} {position: 26} {position: 27} {position: 28} {position: 29} {position: 30} {position: 31} {position: 32} {position: 33} {position: 34} {position: 35} {position: 36} {position: 37} {position: 38} {position: 39} {position: 40} {position: 41} {position: 42} {position: 43} {position: 44} {position: 45} {position: 46} {position: 47} {position: 48} {position: 49} {position: 50} {position: 51} {position: 52} {position: 53} {position: 54} {position: 55} {position: 56} {position: 57} {position: 58} {position: 59} {position: 60} {position: 61} {position: 62} {position: 63} {position: 64} {position: 65} {position: 66} {position: 67} {position: 68} {position: 69} {position: 70} {position: 71} {position: 72} {position: 73} {position: 74} {position: 75} {position: 76} {position: 77} {position: 78} {position: 79} {position: 80} {position: 81} {position: 82} {position: 83} {position: 84} {position: 85} {position: 86} {position: 87} {position: 88} {position: 89} {position: 90} {position: 91} {position: 92} {position: 93} {position: 94} {position: 95} {position: 96} {position: 97} {position: 98} {position: 99} {position: 100}]`
	assert.Equal(t, expectedString, board.String())
}
