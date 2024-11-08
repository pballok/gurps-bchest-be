package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
	"github.com/pballok/gurps-bchest-be/internal/storage"
)

// ImportGCA5Character is the resolver for the importGCA5Character field.
func (r *mutationResolver) ImportGCA5Character(ctx context.Context, input model.CharacterGCA5Import) (*model.Character, error) {
	newChar, err := character.ImportGCA5Character(input.Campaign, []byte(input.Data))
	if err != nil {
		return nil, err
	}

	_, err = r.Storage.Characters.Add(newChar)
	if err != nil {
		return nil, err
	}

	modelChar := newChar.ToModel()

	slog.Info(fmt.Sprintf(`Imported new GCA5 character "%s" in campaign "%s"`, modelChar.Name, modelChar.Campaign))

	return &modelChar, nil
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context, campaign string) ([]*model.Character, error) {
	chars := r.Storage.Characters.List(storage.CharacterFilterType{Campaign: &campaign})

	modelChars := make([]*model.Character, 0)
	for _, char := range chars {
		modelChar := char.ToModel()
		modelChars = append(modelChars, &modelChar)
	}

	slog.Info(fmt.Sprintf(`Retrieved %d characters in campaign "%s"`, len(modelChars), campaign))
	return modelChars, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, campaign string, name string) (*model.Character, error) {
	c, err := r.Storage.Characters.Get(storage.CharacterKeyType{
		Campaign: campaign,
		Name:     name,
	})

	if err != nil {
		return nil, err
	}

	modelChar := c.ToModel()
	slog.Info(fmt.Sprintf(`Retrieved character "%s" in campaign "%s"`, modelChar.Name, modelChar.Campaign))
	return &modelChar, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
