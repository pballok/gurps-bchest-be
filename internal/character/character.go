package character

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/pballok/gurps-bchest-be/internal/attribute"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
)

type Character interface {
	Name() string
	Player() string
	Campaign() string
	Points() int
	Attribute(attributeType model.AttributeType) attribute.Attribute
}

type character struct {
	name       string
	player     string
	campaign   string
	points     int
	attributes map[model.AttributeType]attribute.Attribute
}

type characterGCA5Import struct {
	Name     string  `json:"CharacterName"`
	Player   string  `json:"Playername"`
	Points   float64 `json:"TotalPoints"`
	STCost   float64 `json:"StrengthPoints"`
	DXCost   float64 `json:"DexterityPoints"`
	IQCost   float64 `json:"IntelligencePoints"`
	HTCost   float64 `json:"HealthPoints"`
	PerCost  float64 `json:"PerceptionPoints"`
	WillCost float64 `json:"WillpowerPoints"`
	BSCost   float64 `json:"BasicSpeedPoints"`
	BMCost   float64 `json:"BasicMovePoints"`
	HPCost   float64 `json:"HitPointsPoints"`
	CurrHP   float64 `json:"HitPoints"`
	FPCost   float64 `json:"FatiguePointsPoints"`
	CurrFP   float64 `json:"FatiguePoints"`
}

func NewCharacter(name string, player string, campaign string, points int) Character {
	c := &character{
		name:       name,
		player:     player,
		campaign:   campaign,
		points:     points,
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
	c.attributes[model.AttributeTypeCurrHp] = attribute.NewAttribute(
		func() float64 { return c.attributes[model.AttributeTypeHp].Value() },
		func() int { return 0 },
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
	c.attributes[model.AttributeTypeCurrFp] = attribute.NewAttribute(
		func() float64 { return c.attributes[model.AttributeTypeFp].Value() },
		func() int { return 0 },
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

func ImportGCA5Character(campaign string, jsonString []byte) (Character, error) {
	characterData := characterGCA5Import{}
	err := json.Unmarshal(jsonString, &characterData)
	if err != nil {
		return nil, fmt.Errorf("character import error: %w", err)
	}

	c := NewCharacter(characterData.Name, characterData.Player, campaign, int(characterData.Points))
	c.Attribute(model.AttributeTypeSt).SetCost(int(characterData.STCost))
	c.Attribute(model.AttributeTypeDx).SetCost(int(characterData.DXCost))
	c.Attribute(model.AttributeTypeIq).SetCost(int(characterData.IQCost))
	c.Attribute(model.AttributeTypeHt).SetCost(int(characterData.HTCost))
	c.Attribute(model.AttributeTypeHp).SetCost(int(characterData.HPCost))
	c.Attribute(model.AttributeTypeWill).SetCost(int(characterData.WillCost))
	c.Attribute(model.AttributeTypePer).SetCost(int(characterData.PerCost))
	c.Attribute(model.AttributeTypeFp).SetCost(int(characterData.FPCost))
	c.Attribute(model.AttributeTypeBs).SetCost(int(characterData.BSCost))
	c.Attribute(model.AttributeTypeBm).SetCost(int(characterData.BMCost))
	c.Attribute(model.AttributeTypeCurrHp).SetModifier(characterData.CurrHP - c.Attribute(model.AttributeTypeHp).Value())
	c.Attribute(model.AttributeTypeCurrFp).SetModifier(characterData.CurrFP - c.Attribute(model.AttributeTypeFp).Value())

	return c, nil
}

func (c *character) Name() string {
	return c.name
}

func (c *character) Player() string {
	return c.player
}

func (c *character) Campaign() string {
	return c.campaign
}

func (c *character) Points() int {
	return c.points
}

func (c *character) Attribute(attributeType model.AttributeType) attribute.Attribute {
	if attributeType.IsValid() {
		return c.attributes[attributeType]
	}

	return attribute.NewAttribute(func() float64 { return 0.0 }, func() int { return 0 })
}
