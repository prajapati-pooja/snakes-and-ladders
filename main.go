package main

import (
	"fmt"
	"snakes-and-ladders/models"
)

const MaxRunCount = 10

func main() {
	snake := models.Snake{
		Head: 14,
		Tail: 8,
	}
	board := models.NewBoard(100, snake)
	player := models.Player{
		Name: "Player 1",
	}

	game := models.NewGame(board, player)
	game.Start()

	regularDice := models.RegularDice{MinValue: 1, MaxValue: 6}
	playWith(regularDice, game)

	crookedDice := models.CrookedDice{MinValue: 1, MaxValue: 6}
	playWith(crookedDice, game)

}

func playWith(dice models.Dice, game models.Game) {
	maxRunCount := MaxRunCount
	runCount := 0

	for runCount < maxRunCount {
		game.Play(dice)
		runCount = runCount+1
	}
	fmt.Printf("%+v", game.State())
}

