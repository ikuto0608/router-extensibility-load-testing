package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"router-coprocessor-proj/subgraph-b/graph/model"
)

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	return GetAllLocations(), nil
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, id string) (*model.Location, error) {
	var loc, notFound = GetLocationFromList(id)

	if notFound {
		return nil, fmt.Errorf("Location not found")
	}

	return &loc, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
