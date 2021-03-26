package models

import (
	"math/rand"
	"time"
)

type Dice interface {
	Throw() int
}

type RegularDice struct {
	MinValue int
	MaxValue int
}

func (d RegularDice) Throw() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.MaxValue-d.MinValue) + d.MinValue
}
