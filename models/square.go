package models

type Square interface {
	Shift(moves int) int
}

type RegularSquare struct {
	Position int
}

func (rs RegularSquare) Shift(moves int) int {
	return rs.Position + moves
}