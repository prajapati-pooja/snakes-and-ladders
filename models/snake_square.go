package models

type SnakeSquare struct {
	snake  Snake
	player Player
}

func (ss *SnakeSquare) getNextPosition() int {
	return ss.snake.Tail
}

func (ss *SnakeSquare) enter(player Player) {
	ss.player = player
}
func (ss *SnakeSquare) leave() {
	ss.player = Player{}
}

