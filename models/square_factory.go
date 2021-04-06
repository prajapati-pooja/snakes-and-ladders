package models

func getSquare(position int, snake Snake, ladder Ladder)  Square {
	switch position {
	case ladder.Bottom:
		return &LadderSquare{position: position, ladder: ladder}
	case snake.Head:
		return &SnakeSquare{position: position, snake: snake}
	default:
		return &RegularSquare{position: position}
	}
}
