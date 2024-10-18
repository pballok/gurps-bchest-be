package attribute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute_DefaultCreation(t *testing.T) {
	attr := NewAttribute(func() float64 { return 1.0 }, func() int { return 0 })

	assert.Equal(t, 0.0, attr.Modifier())
	assert.Equal(t, 1.0, attr.Value())
}

func TestAttribute_SetModifier(t *testing.T) {
	attr := NewAttribute(func() float64 { return 1.0 }, func() int { return 0 })
	attr.SetModifier(2.0)

	assert.Equal(t, 2.0, attr.Modifier())
}

func TestAttribute_BaseValueIsAddedToValue(t *testing.T) {
	var base = 1.0
	attr := NewAttribute(func() float64 { return base }, func() int { return 0 })

	assert.Equal(t, 1.0, attr.Value())

	base = 2.0
	assert.Equal(t, 2.0, attr.Value())
}

func TestAttribute_ModifierIsAddedToValue(t *testing.T) {
	attr := NewAttribute(func() float64 { return 1.0 }, func() int { return 0 })
	attr.SetModifier(2.0)

	assert.Equal(t, 3.0, attr.Value())
}

func TestAttribute_CostCalculation(t *testing.T) {
	attr := NewAttribute(func() float64 { return 1.0 }, func() int { return 5 })
	attr.SetModifier(2.0)

	assert.Equal(t, 10, attr.Cost())
}
