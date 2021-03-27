package models

import "fmt"

type Game struct {
	board *Board
	player Player
}

func NewGame(board *Board, player Player) Game {
	return Game{
		board:  board,
		player: player,
	}
}

func (g Game) Start() {
	g.board.registerInFirstSquare(g.player)
}

func (g Game) Play(dice Dice) {
	moves := dice.Throw()
	g.board.findNextSquare(g.player, moves)
}

func (g Game) State() string {
	return fmt.Sprint(g.board)
}

