// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Translation -.
type Movie struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"name"         example:"spongebob the movie"`
	description string             `json:"name"         example:"spongebob the movie"`
	CreatedAt   time.Time          `json:"created_at"   example:"2022-02-17 13:39:03.809450"`
}
