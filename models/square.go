package models

type Square interface {
	Shift() int
}

type RegularSquare struct {
	position int
	board    *Board
}

func (rs RegularSquare) Shift() int {
	moves := rs.board.currentMove()
	return rs.position + moves
}