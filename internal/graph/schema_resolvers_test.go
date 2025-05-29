package graph

import (
	"errors"
	"strconv"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
	"github.com/pballok/gurps-bchest-be/internal/mocks"
	"github.com/pballok/gurps-bchest-be/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testCharacterName1 = "Chuck Norris"
var testCharacterName2 = "Bruce Lee"
var testPlayerName1 = "Player 1"
var testPlayerName2 = "Player 2"
var testCampaign = "Test Campaign"
var testPoints = 100
var testCharacter1 = character.NewCharacter(
	testCharacterName1,
	testPlayerName1,
	testCampaign,
	testPoints,
)
var testCharacter2 = character.NewCharacter(
	testCharacterName2,
	testPlayerName2,
	testCampaign,
	testPoints,
)

const importData string = `
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

type graphqlResponse struct {
	ImportGCA5Character *model.Character   `json:"importGCA5Character"`
	Characters          []*model.Character `json:"characters"`
	Character           *model.Character   `json:"character"`
}

func createTestClient(resolver *Resolver) *client.Client {
	configGraph := Config{Resolvers: resolver}
	srv := handler.New(NewExecutableSchema(configGraph))

	srv.AddTransport(transport.POST{})

	return client.New(srv)
}

func TestSchemaResolvers_ImportGCA5Character_Success(t *testing.T) {
	testImport := `{
		"CharacterName": "` + testCharacterName1 + `",
		"Playername": "` + testPlayerName1 + `",
		"TotalPoints": ` + strconv.Itoa(testPoints) + `,
		"StrengthPoints": 20.0,
		"DexterityPoints": -20.0,
		"HitPoints": 11.0,
		"FatiguePoints": 8.0
	}`
	expectedCharacter := character.NewCharacter(testCharacterName1, testPlayerName1, testCampaign, testPoints)
	expectedCharacter.Attribute(model.AttributeTypeSt).SetModifier(2)
	expectedCharacter.Attribute(model.AttributeTypeDx).SetModifier(-1)
	expectedCharacter.Attribute(model.AttributeTypeCurrHp).SetModifier(-1)
	expectedCharacter.Attribute(model.AttributeTypeCurrFp).SetModifier(-2)

	mockedCharacterStorable := mocks.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Add(mock.Anything, mock.MatchedBy(func(c character.Character) bool {
		return c.Name() == testCharacterName1 &&
			c.Campaign() == testCampaign &&
			c.Player() == testPlayerName1 &&
			c.Points() == testPoints
	})).Return(storage.CharacterKeyType{Name: expectedCharacter.Name(), Campaign: expectedCharacter.Campaign()}, nil)

	mockedStorage := mocks.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage: mockedStorage,
	})
	importInput := model.ImportGCA5CharacterInput{
		Campaign: expectedCharacter.Campaign(),
		Data:     testImport,
	}
	query := `
      mutation importGCA5Character($input: ImportGCA5CharacterInput!) {
        importGCA5Character(input: $input) {
          campaign,
          name,
          player,
          availablePoints
        }
      }`
	response := graphqlResponse{}

	err := graphqlClient.Post(query, &response, client.Var("input", importInput))

	assert.NoError(t, err)
	assert.Equal(t, expectedCharacter.Name(), response.ImportGCA5Character.Name)
	assert.Equal(t, expectedCharacter.Player(), response.ImportGCA5Character.Player)
	assert.Equal(t, expectedCharacter.Campaign(), response.ImportGCA5Character.Campaign)
	assert.Equal(t, expectedCharacter.Points(), response.ImportGCA5Character.AvailablePoints)
}

func TestSchemaResolvers_ImportGCA5Character_ImportError(t *testing.T) {
	mockedStorage := mocks.NewMockStorage(t)

	graphqlClient := createTestClient(&Resolver{
		Storage: mockedStorage,
	})
	importInput := model.ImportGCA5CharacterInput{
		Campaign: "Test Campaign",
		Data:     "invalid data",
	}
	query := `
      mutation importGCA5Character($input: ImportGCA5CharacterInput!) {
        importGCA5Character(input: $input) {
          campaign,
          name,
          player,
          availablePoints
        }
      }`
	response := graphqlResponse{}

	err := graphqlClient.Post(query, &response, client.Var("input", importInput))

	assert.Error(t, err)
}

func TestSchemaResolvers_ImportGCA5Character_StorageAddError(t *testing.T) {
	testImport := `{
		"CharacterName": "Test",
		"Playername": "Player",
		"TotalPoints": 100.0,
		"StrengthPoints": 20.0,
		"DexterityPoints": -10.0
	}`

	mockedCharacterStorable := mocks.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Add(mock.Anything, mock.Anything).Return(storage.CharacterKeyType{}, errors.New("uh-oh"))

	mockedStorage := mocks.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage: mockedStorage,
	})
	importInput := model.ImportGCA5CharacterInput{
		Campaign: "Test Campaign",
		Data:     testImport,
	}
	query := `
      mutation importGCA5Character($input: ImportGCA5CharacterInput!) {
        importGCA5Character(input: $input) {
          campaign,
          name,
          player,
          availablePoints
        }
      }`
	response := graphqlResponse{}

	err := graphqlClient.Post(query, &response, client.Var("input", importInput))

	assert.Error(t, err)
}

func TestSchemaResolvers_ListCharacters_Success(t *testing.T) {
	mockedCharacterStorable := mocks.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().List(mock.Anything, storage.CharacterFilterType{Campaign: &testCampaign}).Return([]character.Character{testCharacter1, testCharacter2}, nil)

	mockedStorage := mocks.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage: mockedStorage,
	})
	query := `
      query listCharacters {
        characters(campaign: "` + testCampaign + `") {
          campaign,          
          name,
          player,
          availablePoints
        }
      }`
	response := graphqlResponse{}

	err := graphqlClient.Post(query, &response)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(response.Characters))
	assert.Equal(t, testCharacterName1, response.Characters[0].Name)
	assert.Equal(t, testPlayerName1, response.Characters[0].Player)
	assert.Equal(t, testCampaign, response.Characters[0].Campaign)
	assert.Equal(t, testPoints, response.Characters[0].AvailablePoints)
	assert.Equal(t, testCharacterName2, response.Characters[1].Name)
	assert.Equal(t, testPlayerName2, response.Characters[1].Player)
	assert.Equal(t, testCampaign, response.Characters[1].Campaign)
	assert.Equal(t, testPoints, response.Characters[1].AvailablePoints)
}

func TestSchemaResolvers_GetCharacter_Success(t *testing.T) {
	mockedCharacterStorable := mocks.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Get(mock.Anything, storage.CharacterKeyType{Campaign: testCampaign, Name: testCharacterName1}).Return(testCharacter1, nil)

	mockedStorage := mocks.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage: mockedStorage,
	})
	query := `
      query getCharacter {
        character(campaign: "` + testCampaign + `", name: "` + testCharacterName1 + `") {
          campaign,          
          name,
          player,
          availablePoints
        }
      }`
	response := graphqlResponse{}

	err := graphqlClient.Post(query, &response)

	assert.NoError(t, err)
	assert.Equal(t, testCharacterName1, response.Character.Name)
	assert.Equal(t, testPlayerName1, response.Character.Player)
	assert.Equal(t, testCampaign, response.Character.Campaign)
	assert.Equal(t, testPoints, response.Character.AvailablePoints)
}

func TestSchemaResolvers_GetCharacter_StorageGetError(t *testing.T) {
	mockedCharacterStorable := mocks.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Get(mock.Anything, storage.CharacterKeyType{Campaign: testCampaign, Name: testCharacterName1}).Return(nil, errors.New("uh-oh"))

	mockedStorage := mocks.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage: mockedStorage,
	})
	query := `
      query getCharacter {
        character(campaign: "` + testCampaign + `", name: "` + testCharacterName1 + `") {
          campaign,          
          name,
          player,
          availablePoints
        }
      }`
	response := graphqlResponse{}

	err := graphqlClient.Post(query, &response)

	assert.Error(t, err)
}
