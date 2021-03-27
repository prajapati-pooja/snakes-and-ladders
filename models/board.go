package models

type Board struct {
	squares      []Square
	currentMoves int
}

func NewBoard(maxGridSize int, snake Snake) *Board {
	board := &Board{}
	var position int
	var squares []Square
	for position = 1; position <= maxGridSize; position++ {
		square := getSquare(position, snake, board)
		squares = append(squares, square)
	}
	board.squares = squares
	return board
}


func (b *Board) firstSquare() Square {
	return b.squares[0]
}

func (b *Board) findNextSquare(square Square, moves int) Square {
	b.currentMoves = moves
	nextPosition := square.Shift()
	return b.squares[nextPosition-1]
}

func (b *Board) currentMove() int {
	return b.currentMoves
}