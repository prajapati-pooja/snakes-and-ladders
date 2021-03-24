package models

type SnakeSquare struct {
	Snake Snake
}

func (rs SnakeSquare) Shift(moves int) int {
	return rs.Snake.tail
}
