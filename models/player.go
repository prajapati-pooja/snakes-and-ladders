package models

type Player struct {
	Board    *Board
	Position Square
}

func (p *Player) play(moves int) {
	p.Position = p.Board.findNextSquare(p.Position, moves)
}

func (p *Player) StartGame() {
	p.Position = p.Board.firstSquare()
}
