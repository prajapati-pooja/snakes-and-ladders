package models

import "fmt"

type SnakeSquare struct {
	position int
	snake    Snake
	player   Player
}

func (ss *SnakeSquare) getNextPosition(moves int) (int, error) {
	return ss.snake.Tail, nil
}

func (ss *SnakeSquare) enter(player Player) {
	ss.player = player
}
func (ss *SnakeSquare) leave() {
	ss.player = Player{}
}

func (ss *SnakeSquare) String() string {
	if isEmpty(ss.player.Name) {
		return fmt.Sprintf("{position: %d, snake: %+v}", ss.position, ss.snake)
	}
	return fmt.Sprintf("{position: %d, player: %+v, snake: %+v}", ss.position, ss.player, ss.snake)
}

func isEmpty(playerName string) bool {
	return playerName == ""
}
