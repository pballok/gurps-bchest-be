package character

import (
	"testing"

	"github.com/pballok/gurps-bchest-be/internal/graph/model"
	"github.com/stretchr/testify/assert"
)

const defaultImportString string = `
{
    "CharacterName": "Test",
    "Playername": "Player",
    "TotalPoints": 100.0,
	"StrengthPoints": 10.0,
	"DexterityPoints": -20.0,
	"IntelligencePoints": 40.0,
	"HealthPoints": -20.0,
	"HitPointsPoints": 6.0,
	"HitPoints": 11.0,
	"WillpowerPoints": 20.0,
	"PerceptionPoints": -20.0,
	"FatiguePointsPoints": 15.0,
	"FatiguePoints": 8.0,
	"BasicSpeedPoints": 50.0,
	"BasicMovePoints": 5.0
}`

func TestCharacter_Creation(t *testing.T) {
	c := NewCharacter("Test", "Player", 10)

	assert.Equal(t, "Test", c.Name())
	assert.Equal(t, "Player", c.Player())
	assert.Equal(t, 10, c.Points())

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
	c := NewCharacter("Test", "Player", 10)

	c.Attribute(model.AttributeTypeSt).SetModifier(1.0)
	c.Attribute(model.AttributeTypeDx).SetModifier(-1.0)
	c.Attribute(model.AttributeTypeIq).SetModifier(2.0)
	c.Attribute(model.AttributeTypeHt).SetModifier(-2.0)
	c.Attribute(model.AttributeTypeHp).SetModifier(3.0)
	c.Attribute(model.AttributeTypeCurrHp).SetModifier(-3.0)
	c.Attribute(model.AttributeTypeWill).SetModifier(4.0)
	c.Attribute(model.AttributeTypePer).SetModifier(-4.0)
	c.Attribute(model.AttributeTypeFp).SetModifier(5.0)
	c.Attribute(model.AttributeTypeCurrFp).SetModifier(-5.0)
	c.Attribute(model.AttributeTypeBs).SetModifier(2.5)
	c.Attribute(model.AttributeTypeBm).SetModifier(1.0)

	assert.Equal(t, 11.0, c.Attribute(model.AttributeTypeSt).Value())
	assert.Equal(t, 1.0, c.Attribute(model.AttributeTypeSt).Modifier())
	assert.Equal(t, 10, c.Attribute(model.AttributeTypeSt).Cost())

	assert.Equal(t, 9.0, c.Attribute(model.AttributeTypeDx).Value())
	assert.Equal(t, -1.0, c.Attribute(model.AttributeTypeDx).Modifier())
	assert.Equal(t, -20, c.Attribute(model.AttributeTypeDx).Cost())

	assert.Equal(t, 12.0, c.Attribute(model.AttributeTypeIq).Value())
	assert.Equal(t, 2.0, c.Attribute(model.AttributeTypeIq).Modifier())
	assert.Equal(t, 40, c.Attribute(model.AttributeTypeIq).Cost())

	assert.Equal(t, 8.0, c.Attribute(model.AttributeTypeHt).Value())
	assert.Equal(t, -2.0, c.Attribute(model.AttributeTypeHt).Modifier())
	assert.Equal(t, -20, c.Attribute(model.AttributeTypeHt).Cost())

	assert.Equal(t, 14.0, c.Attribute(model.AttributeTypeHp).Value())
	assert.Equal(t, 3.0, c.Attribute(model.AttributeTypeHp).Modifier())
	assert.Equal(t, 6, c.Attribute(model.AttributeTypeHp).Cost())

	assert.Equal(t, 11.0, c.Attribute(model.AttributeTypeCurrHp).Value())
	assert.Equal(t, -3.0, c.Attribute(model.AttributeTypeCurrHp).Modifier())
	assert.Equal(t, 0, c.Attribute(model.AttributeTypeCurrHp).Cost())

	assert.Equal(t, 16.0, c.Attribute(model.AttributeTypeWill).Value())
	assert.Equal(t, 4.0, c.Attribute(model.AttributeTypeWill).Modifier())
	assert.Equal(t, 20, c.Attribute(model.AttributeTypeWill).Cost())

	assert.Equal(t, 8.0, c.Attribute(model.AttributeTypePer).Value())
	assert.Equal(t, -4.0, c.Attribute(model.AttributeTypePer).Modifier())
	assert.Equal(t, -20, c.Attribute(model.AttributeTypePer).Cost())

	assert.Equal(t, 13.0, c.Attribute(model.AttributeTypeFp).Value())
	assert.Equal(t, 5.0, c.Attribute(model.AttributeTypeFp).Modifier())
	assert.Equal(t, 15, c.Attribute(model.AttributeTypeFp).Cost())

	assert.Equal(t, 8.0, c.Attribute(model.AttributeTypeCurrFp).Value())
	assert.Equal(t, -5.0, c.Attribute(model.AttributeTypeCurrFp).Modifier())
	assert.Equal(t, 0, c.Attribute(model.AttributeTypeCurrFp).Cost())

	assert.Equal(t, 6.75, c.Attribute(model.AttributeTypeBs).Value())
	assert.Equal(t, 2.5, c.Attribute(model.AttributeTypeBs).Modifier())
	assert.Equal(t, 50, c.Attribute(model.AttributeTypeBs).Cost())

	assert.Equal(t, 7.0, c.Attribute(model.AttributeTypeBm).Value())
	assert.Equal(t, 1.0, c.Attribute(model.AttributeTypeBm).Modifier())
	assert.Equal(t, 5, c.Attribute(model.AttributeTypeBm).Cost())
}

func TestCharacter_ImportProperties(t *testing.T) {
	c, err := ImportGCA5Character([]byte(defaultImportString))

	assert.NoError(t, err)
	assert.Equal(t, "Test", c.Name())
	assert.Equal(t, "Player", c.Player())
	assert.Equal(t, 100, c.Points())
}

func TestCharacter_ImportAttributes(t *testing.T) {
	c, err := ImportGCA5Character([]byte(defaultImportString))

	assert.NoError(t, err)

	assert.Equal(t, 11.0, c.Attribute(model.AttributeTypeSt).Value())
	assert.Equal(t, 1.0, c.Attribute(model.AttributeTypeSt).Modifier())
	assert.Equal(t, 10, c.Attribute(model.AttributeTypeSt).Cost())

	assert.Equal(t, 9.0, c.Attribute(model.AttributeTypeDx).Value())
	assert.Equal(t, -1.0, c.Attribute(model.AttributeTypeDx).Modifier())
	assert.Equal(t, -20, c.Attribute(model.AttributeTypeDx).Cost())

	assert.Equal(t, 12.0, c.Attribute(model.AttributeTypeIq).Value())
	assert.Equal(t, 2.0, c.Attribute(model.AttributeTypeIq).Modifier())
	assert.Equal(t, 40, c.Attribute(model.AttributeTypeIq).Cost())

	assert.Equal(t, 8.0, c.Attribute(model.AttributeTypeHt).Value())
	assert.Equal(t, -2.0, c.Attribute(model.AttributeTypeHt).Modifier())
	assert.Equal(t, -20, c.Attribute(model.AttributeTypeHt).Cost())

	assert.Equal(t, 14.0, c.Attribute(model.AttributeTypeHp).Value())
	assert.Equal(t, 3.0, c.Attribute(model.AttributeTypeHp).Modifier())
	assert.Equal(t, 6, c.Attribute(model.AttributeTypeHp).Cost())

	assert.Equal(t, 11.0, c.Attribute(model.AttributeTypeCurrHp).Value())
	assert.Equal(t, -3.0, c.Attribute(model.AttributeTypeCurrHp).Modifier())
	assert.Equal(t, 0, c.Attribute(model.AttributeTypeCurrHp).Cost())

	assert.Equal(t, 16.0, c.Attribute(model.AttributeTypeWill).Value())
	assert.Equal(t, 4.0, c.Attribute(model.AttributeTypeWill).Modifier())
	assert.Equal(t, 20, c.Attribute(model.AttributeTypeWill).Cost())

	assert.Equal(t, 8.0, c.Attribute(model.AttributeTypePer).Value())
	assert.Equal(t, -4.0, c.Attribute(model.AttributeTypePer).Modifier())
	assert.Equal(t, -20, c.Attribute(model.AttributeTypePer).Cost())

	assert.Equal(t, 13.0, c.Attribute(model.AttributeTypeFp).Value())
	assert.Equal(t, 5.0, c.Attribute(model.AttributeTypeFp).Modifier())
	assert.Equal(t, 15, c.Attribute(model.AttributeTypeFp).Cost())

	assert.Equal(t, 8.0, c.Attribute(model.AttributeTypeCurrFp).Value())
	assert.Equal(t, -5.0, c.Attribute(model.AttributeTypeCurrFp).Modifier())
	assert.Equal(t, 0, c.Attribute(model.AttributeTypeCurrFp).Cost())

	assert.Equal(t, 6.75, c.Attribute(model.AttributeTypeBs).Value())
	assert.Equal(t, 2.5, c.Attribute(model.AttributeTypeBs).Modifier())
	assert.Equal(t, 50, c.Attribute(model.AttributeTypeBs).Cost())

	assert.Equal(t, 7.0, c.Attribute(model.AttributeTypeBm).Value())
	assert.Equal(t, 1.0, c.Attribute(model.AttributeTypeBm).Modifier())
	assert.Equal(t, 5, c.Attribute(model.AttributeTypeBm).Cost())
}
