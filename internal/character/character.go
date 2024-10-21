package character

import (
	"math"

	"github.com/pballok/gurps-bchest-be/internal/attribute"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
)

type Character interface {
	Name() string
	Attribute(attributeType model.AttributeType) attribute.Attribute
}

type character struct {
	name       string
	attributes map[model.AttributeType]attribute.Attribute
}

func NewCharacter(name string) Character {
	c := &character{
		name:       name,
		attributes: make(map[model.AttributeType]attribute.Attribute),
	}

	c.attributes[model.AttributeTypeSt] = attribute.NewAttribute(
		func() float64 { return 10.0 },
		func() int { return 10 },
	)
	c.attributes[model.AttributeTypeDx] = attribute.NewAttribute(
		func() float64 { return 10.0 },
		func() int { return 20 },
	)
	c.attributes[model.AttributeTypeIq] = attribute.NewAttribute(
		func() float64 { return 10.0 },
		func() int { return 20 },
	)
	c.attributes[model.AttributeTypeHt] = attribute.NewAttribute(
		func() float64 { return 10.0 },
		func() int { return 10 },
	)
	c.attributes[model.AttributeTypeHp] = attribute.NewAttribute(
		func() float64 { return c.attributes[model.AttributeTypeSt].Value() },
		func() int { return 2 },
	)
	c.attributes[model.AttributeTypeWill] = attribute.NewAttribute(
		func() float64 { return c.attributes[model.AttributeTypeIq].Value() },
		func() int { return 5 },
	)
	c.attributes[model.AttributeTypePer] = attribute.NewAttribute(
		func() float64 { return c.attributes[model.AttributeTypeIq].Value() },
		func() int { return 5 },
	)
	c.attributes[model.AttributeTypeFp] = attribute.NewAttribute(
		func() float64 { return c.attributes[model.AttributeTypeHt].Value() },
		func() int { return 3 },
	)
	c.attributes[model.AttributeTypeBs] = attribute.NewAttribute(
		func() float64 {
			return (c.attributes[model.AttributeTypeHt].Value() + c.attributes[model.AttributeTypeDx].Value()) / 4.0
		},
		func() int { return 20 },
	)
	c.attributes[model.AttributeTypeBm] = attribute.NewAttribute(
		func() float64 { return math.Floor(c.attributes[model.AttributeTypeBs].Value()) },
		func() int { return 5 },
	)

	return c
}

func (c *character) Name() string {
	return c.name
}

func (c *character) Attribute(attributeType model.AttributeType) attribute.Attribute {
	if attributeType.IsValid() {
		return c.attributes[attributeType]
	}

	return attribute.NewAttribute(func() float64 { return 0.0 }, func() int { return 0 })
}
