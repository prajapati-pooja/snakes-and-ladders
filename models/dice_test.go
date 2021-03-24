package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDice_Roll(t *testing.T) {
	minValue := 1
	maxValue := 6
	dice := Dice{MinValue: minValue, MaxValue: maxValue}
	diceValue := dice.Roll()
	assert.Equal(t, true, diceValue > minValue && diceValue < maxValue)
}
