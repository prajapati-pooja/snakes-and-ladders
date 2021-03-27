package models

func getSquare(position int, snake Snake, board *Board)  Square {
	var square Square
	if position == snake.Head {
		square = &SnakeSquare{position: position, snake: snake}
	} else {
		square = &RegularSquare{position: position, board: board}
	}
	return square
}
