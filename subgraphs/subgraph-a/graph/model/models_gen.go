// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Location struct {
	ID string `json:"id"`
	// The calculated overall rating based on all reviews
	OverallRating *float64 `json:"overallRating,omitempty"`
	// All submitted reviews about this location
	ReviewsForLocation []*Review `json:"reviewsForLocation"`
}

func (Location) IsEntity() {}

type Review struct {
	ID string `json:"id"`
	// Written text
	Comment *string `json:"comment,omitempty"`
	// A number from 1 - 5 with 1 being lowest and 5 being highest
	Rating *int `json:"rating,omitempty"`
	// The location the review is about
	Location *Location `json:"location,omitempty"`
}
