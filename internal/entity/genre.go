// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Genre struct {
	Id   string `json:"id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	Name string `json:"name"  example:"spongebob the movie"`
}
