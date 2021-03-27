package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegularDice_ThrowShouldReturnRandomValueBetween1To6(t *testing.T) {
	minValue := 1
	maxValue := 6
	dice := RegularDice{MinValue: minValue, MaxValue: maxValue}
	diceValue := dice.Throw()
	assert.Equal(t, true, diceValue >= minValue && diceValue <= maxValue)
}
