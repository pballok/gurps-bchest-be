package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/pballok/gurps-bchest-be/internal/character"
	"github.com/pballok/gurps-bchest-be/internal/graph/model"
)

// ImportGCA5Character is the resolver for the importGCA5Character field.
func (r *mutationResolver) ImportGCA5Character(ctx context.Context, input model.CharacterGCA5Import) (*model.Character, error) {
	newChar, err := character.ImportGCA5Character(input.Campaign, []byte(input.Data))
	if err != nil {
		return nil, err
	}

	id, err := r.Storage.Characters.Add(newChar)
	if err != nil {
		return nil, err
	}

	modelChar := newChar.ToModel()
	modelChar.ID = id

	return &modelChar, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, name string) (*model.Character, error) {
	return &model.Character{
		Name: "Test Character",
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
