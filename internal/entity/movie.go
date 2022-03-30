// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"time"
)

// Translation -.
type Movie struct {
	Id          string    `json:"id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	Title       string    `json:"name"         example:"spongebob the movie"`
	Description string    `json:"description"         example:"spongebob the movie"`
	Genres      []string  `json:"genre_ids"              example:"horror = 1"`
	CreatedAt   time.Time `json:"created_at"   example:"2022-02-17 13:39:03.809450"`
}
