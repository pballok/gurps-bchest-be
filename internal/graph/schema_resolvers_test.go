package graph

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
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
	mockedImporter := character.NewMockImporterFunc(t)
	mockedImporter.EXPECT().Execute("Test Campaign", mock.Anything).Return(testCharacter1, nil)

	mockedCharacterStorable := storage.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Add(mock.Anything).Return(storage.CharacterKeyType{Name: testCharacterName1, Campaign: testCampaign}, nil)

	mockedStorage := storage.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage:           mockedStorage,
		CharacterImporter: mockedImporter.Execute,
	})
	importInput := model.ImportGCA5CharacterInput{
		Campaign: "Test Campaign",
		Data:     "import data",
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
	assert.Equal(t, testCharacterName1, response.ImportGCA5Character.Name)
	assert.Equal(t, testPlayerName1, response.ImportGCA5Character.Player)
	assert.Equal(t, testCampaign, response.ImportGCA5Character.Campaign)
	assert.Equal(t, testPoints, response.ImportGCA5Character.AvailablePoints)
}

func TestSchemaResolvers_ImportGCA5Character_ImportError(t *testing.T) {
	mockedImporter := character.NewMockImporterFunc(t)
	mockedImporter.EXPECT().Execute("Test Campaign", mock.Anything).Return(nil, errors.New("uh-oh"))

	mockedStorage := storage.NewMockStorage(t)

	graphqlClient := createTestClient(&Resolver{
		Storage:           mockedStorage,
		CharacterImporter: mockedImporter.Execute,
	})
	importInput := model.ImportGCA5CharacterInput{
		Campaign: "Test Campaign",
		Data:     "import data",
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
	mockedImporter := character.NewMockImporterFunc(t)
	mockedImporter.EXPECT().Execute("Test Campaign", mock.Anything).Return(testCharacter1, nil)

	mockedCharacterStorable := storage.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Add(mock.Anything).Return(storage.CharacterKeyType{}, errors.New("uh-oh"))

	mockedStorage := storage.NewMockStorage(t)
	mockedStorage.EXPECT().Characters().Return(mockedCharacterStorable)

	graphqlClient := createTestClient(&Resolver{
		Storage:           mockedStorage,
		CharacterImporter: mockedImporter.Execute,
	})
	importInput := model.ImportGCA5CharacterInput{
		Campaign: "Test Campaign",
		Data:     "import data",
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
	mockedCharacterStorable := storage.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().List(storage.CharacterFilterType{Campaign: &testCampaign}).Return([]character.Character{testCharacter1, testCharacter2})

	mockedStorage := storage.NewMockStorage(t)
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
	mockedCharacterStorable := storage.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Get(storage.CharacterKeyType{Campaign: testCampaign, Name: testCharacterName1}).Return(testCharacter1, nil)

	mockedStorage := storage.NewMockStorage(t)
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
	mockedCharacterStorable := storage.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Get(storage.CharacterKeyType{Campaign: testCampaign, Name: testCharacterName1}).Return(nil, errors.New("uh-oh"))

	mockedStorage := storage.NewMockStorage(t)
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
