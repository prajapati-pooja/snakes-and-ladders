package main

import (
	"fmt"
	. "snakes-and-ladders/models"
)

func main() {
	snake := Snake{
		Head: 14,
		Tail: 8,
	}
	ladder := Ladder{Top: 35, Bottom: 11}
	board := NewBoard(MaxGridSize, snake, ladder)
	player := Player{
		Name: "Player 1",
	}

	game := NewGame(board, player)
	game.Start()

	regularDice := RegularDice{MinValue: MinDiceValue, MaxValue: MaxDiceValue}
	playWith(regularDice, game)

	crookedDice := CrookedDice{MinValue: MinDiceValue, MaxValue: MaxDiceValue}
	playWith(crookedDice, game)
}

func playWith(dice Dice, game Game) {
	maxRunCount := MaxRunCount
	runCount := 0

	for runCount < maxRunCount {
		game.Play(dice)
		runCount = runCount+1
	}
	fmt.Printf("%+v", game.State())
}

