package models

import "fmt"

type LadderSquare struct {
	position int
	ladder   Ladder
	player   Player
}

func (ls LadderSquare) getNextPosition(moves int) (int, error) {
	return ls.ladder.Top, nil
}

func (ls *LadderSquare) enter(player Player) {
	ls.player = player
}

func (ls *LadderSquare) leave() {
	ls.player = Player{}
}

func (ls *LadderSquare) String() string {
	if isEmpty(ls.player.Name) {
		return fmt.Sprintf("{position: %d, ladder: %+v}", ls.position, ls.ladder)
	}
	return fmt.Sprintf("{position: %d, player: %+v, ladder: %+v}", ls.position, ls.player, ls.ladder)
}
