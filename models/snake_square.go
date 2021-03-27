package models

type SnakeSquare struct {
	Snake Snake
}

func (rs SnakeSquare) Shift() int {
	return rs.Snake.Tail
}
