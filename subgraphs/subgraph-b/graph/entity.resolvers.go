package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"router-coprocessor-proj/subgraph-b/graph/model"
)

// FindLocationByID is the resolver for the findLocationByID field.
func (r *entityResolver) FindLocationByID(ctx context.Context, id string) (*model.Location, error) {
	var loc, notFound = GetLocationFromList(id)

	if notFound {
		return nil, fmt.Errorf("Location not found")
	}

	return &loc, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
