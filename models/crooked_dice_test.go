package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrookedDice_ShouldThrowEvenNumbersBetween1To6(t *testing.T) {
	minValue := 1
	maxValue := 6
	dice := CrookedDice{MinValue: minValue, MaxValue: maxValue}
	diceValue := dice.Throw()
	assert.Equal(t, true, diceValue%2 == 0)
}

func TestCrookedDice_ShouldNotThrowOddNumbers(t *testing.T) {
	minValue := 1
	maxValue := 6
	dice := CrookedDice{MinValue: minValue, MaxValue: maxValue}
	diceValue := dice.Throw()
	assert.Equal(t, false, diceValue%2 == 1)
}