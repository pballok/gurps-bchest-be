package graph

import (
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

var testCharacterName = "Test Character"
var testPlayerName = "Test Player"
var testCampaign = "Test Campaign"
var testPoints = 100
var testCharacter = character.NewCharacter(
	testCharacterName,
	testPlayerName,
	testCampaign,
	testPoints,
)

type graphqlResponse struct {
	ImportGCA5Character *model.Character `json:"importGCA5Character"`
}

func createTestClient(resolver *Resolver) *client.Client {
	configGraph := Config{Resolvers: resolver}
	srv := handler.New(NewExecutableSchema(configGraph))

	srv.AddTransport(transport.POST{})

	return client.New(srv)
}

func TestSchemaResolvers_ImportGCA5Character_Success(t *testing.T) {
	mockedImporter := character.NewMockImporterFunc(t)
	mockedImporter.EXPECT().Execute("Test Campaign", mock.Anything).Return(testCharacter, nil)

	mockedCharacterStorable := storage.NewMockStorable[storage.CharacterKeyType, character.Character, storage.CharacterFilterType](t)
	mockedCharacterStorable.EXPECT().Add(mock.Anything).Return(storage.CharacterKeyType{Name: testCharacterName, Campaign: testCampaign}, nil)

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
	graphqlClient.MustPost(query, &response, client.Var("input", importInput))

	assert.Equal(t, testCharacterName, response.ImportGCA5Character.Name)
	assert.Equal(t, testPlayerName, response.ImportGCA5Character.Player)
	assert.Equal(t, testCampaign, response.ImportGCA5Character.Campaign)
	assert.Equal(t, testPoints, response.ImportGCA5Character.AvailablePoints)
}
