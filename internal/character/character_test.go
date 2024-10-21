package character

import (
	"testing"

	"github.com/pballok/gurps-bchest-be/internal/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestCharacter_Creation(t *testing.T) {
	c := NewCharacter("Test")

	assert.Equal(t, "Test", c.Name())

	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeSt).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeDx).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeIq).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeHt).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeHp).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeWill).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypePer).Value())
	assert.Equal(t, 10.0, c.Attribute(model.AttributeTypeFp).Value())
	assert.Equal(t, 5.0, c.Attribute(model.AttributeTypeBs).Value())
	assert.Equal(t, 5.0, c.Attribute(model.AttributeTypeBm).Value())
}

func TestCharacter_Attribute(t *testing.T) {
	c := NewCharacter("Test")

	//TODO: Complete this
	c.Attribute(model.AttributeTypeSt).SetModifier(1.0)
	c.Attribute(model.AttributeTypeDx).SetModifier(2.0)
	c.Attribute(model.AttributeTypeIq).SetModifier(3.0)
	c.Attribute(model.AttributeTypeHt).SetModifier(4.0)
	c.Attribute(model.AttributeTypeHp).SetModifier(5.0)
	c.Attribute(model.AttributeTypeWill).SetModifier(6.0)
	c.Attribute(model.AttributeTypePer).SetModifier(7.0)
	c.Attribute(model.AttributeTypeFp).SetModifier(8.0)
	c.Attribute(model.AttributeTypeFp).SetModifier(0.5)

	assert.Equal(t, 11.0, c.Attribute(model.AttributeTypeSt).Value())
	assert.Equal(t, 12.0, c.Attribute(model.AttributeTypeDx).Value())
	assert.Equal(t, 13.0, c.Attribute(model.AttributeTypeIq).Value())
	assert.Equal(t, 14.0, c.Attribute(model.AttributeTypeHt).Value())
	assert.Equal(t, 15.0, c.Attribute(model.AttributeTypeHp).Value())
	assert.Equal(t, 19.0, c.Attribute(model.AttributeTypeWill).Value())
	assert.Equal(t, 20.0, c.Attribute(model.AttributeTypePer).Value())
	assert.Equal(t, 22.0, c.Attribute(model.AttributeTypeFp).Value())
}
