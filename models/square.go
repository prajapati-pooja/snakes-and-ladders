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

func (rs *RegularSquare) getNextPosition() int {
	moves := rs.board.getCurrentMove()
	return rs.position + moves
}

func (rs *RegularSquare) enter(player Player) {
	rs.player = player
}

func (rs *RegularSquare) leave() {
	rs.player = Player{}
}

func (rs *RegularSquare) String() string {
	if isEmpty(rs.player.Name) {
		return fmt.Sprintf("{position: %d}", rs.position)
	}
	return fmt.Sprintf("{position: %d, player: %+v}", rs.position, rs.player)
}
