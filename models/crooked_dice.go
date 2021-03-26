package models

import (
	"math/rand"
	"time"
)

type CrookedDice struct {
	MinValue int
	MaxValue int
}

func (d CrookedDice) Throw() int {
	if isOdd(d.MinValue) {
		d.MinValue = d.MinValue + 1
	}

	if isOdd(d.MaxValue) {
		d.MaxValue = d.MaxValue - 1
	}
	rand.Seed(time.Now().UnixNano())
	randomNumber := (rand.Intn(d.MaxValue-d.MinValue) + d.MinValue+1)/2
	return randomNumber*2
}

func isOdd(minValue int) bool {
	return minValue%2 == 1
}
