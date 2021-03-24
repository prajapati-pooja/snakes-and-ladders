package models

import (
	"math/rand"
	"time"
)

type Dice struct {
	MinValue int
	MaxValue int
}

func (d Dice) Throw() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.MaxValue-d.MinValue) + d.MinValue
}
