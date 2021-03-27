package models

import "fmt"

type Square interface {
	getNextPosition() int
	enter(player Player)
	leave()
}

type RegularSquare struct {
	position int
	board    *Board
	player   Player
}

func (rs RegularSquare) getNextPosition() int {
	moves := rs.board.currentMove()
	return rs.position + moves
}

func (rs *RegularSquare) enter(player Player) {
	rs.player = player
}

func (rs *RegularSquare) leave() {
	rs.player = Player{}
}

func (rs *RegularSquare) String() string {
	return fmt.Sprintf("position: %d, %v", rs.position, rs.player)
}
