package attribute

import "math"

type Attribute interface {
	Modifier() float64
	SetModifier(float64)
	Value() float64
	Cost() int
}

type attribute struct {
	baseValue       func() float64
	costPerModifier func() int
	modifier        float64
}

func NewAttribute(baseValueFunc func() float64, costPerModifierFunc func() int) Attribute {
	return &attribute{
		baseValue:       baseValueFunc,
		costPerModifier: costPerModifierFunc,
		modifier:        0.0,
	}
}

func (a *attribute) Modifier() float64 {
	return a.modifier
}

func (a *attribute) SetModifier(modifier float64) {
	a.modifier = modifier
}

func (a *attribute) Value() float64 {
	return a.baseValue() + a.modifier
}

func (a *attribute) Cost() int {
	return int(math.Ceil(a.modifier * float64(a.costPerModifier())))
}
