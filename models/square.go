package models

import (
	"errors"
	"fmt"
)

type Square interface {
	getNextPosition(moves int) (int, error)
	enter(player Player)
	leave()
}

type RegularSquare struct {
	position int
	player   Player
}

func (rs *RegularSquare) getNextPosition(moves int) (int, error) {
	nextPosition := rs.position + moves
	if nextPosition > MaxGridSize {
		return 0, errors.New("cant move ahead")
	}
	return nextPosition, nil
}

func (rs *RegularSquare) enter(player Player) {
	rs.player = player
}

func (rs *RegularSquare) leave() {
	rs.player = Player{}
}

func (rs *RegularSquare) String() string {
	if isEmpty(rs.player.Name) {
		return fmt.Sprintf("{position: %d}", rs.position)
	}
	return fmt.Sprintf("{position: %d, player: %+v}", rs.position, rs.player)
}
