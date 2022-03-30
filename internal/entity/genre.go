// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Translation -.
type Genre struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"         example:"spongebob the movie"`
}
