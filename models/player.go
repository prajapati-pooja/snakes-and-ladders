package models

type Player struct {
	Board    *Board
	Position Square
}

func (p *Player) Move(distance int) {
	p.Position = p.Board.findNextSquare(p.Position, distance)
}

func (p *Player) StartGame() {
	p.Position = p.Board.firstSquare()
}
