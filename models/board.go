package models

type Board struct {
	squares []Square
}

func NewBoard(maxGridSize int, snake Snake) Board {
	var position int
	var squares []Square
	for position = 1; position <= maxGridSize; position++ {
		square := initializeSquare(position, snake)
		squares = append(squares, square)
	}

	return Board{squares: squares}
}

func initializeSquare(position int, snake Snake) Square {
	var square Square
	if position == snake.Head {
		square = SnakeSquare{Snake: snake}
	} else {
		square = RegularSquare{Position: position}
	}
	return square
}

func (b Board) firstSquare() Square {
	return b.squares[0]
}


func (b Board) findNextSquare(square Square, shifts int) Square {
	nextPosition := square.Shift(shifts)
	return b.squares[nextPosition-1]
}