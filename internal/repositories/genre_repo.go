package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const genreCollectionName = "genre"

// GenreRepo -.
type GenreRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewGenreRepo(mdb *mongodb.MongoDB) *GenreRepo {
	return &GenreRepo{mdb}
}

// GetById -.
func (sr *GenreRepo) GetById(ctx context.Context, genre_id string) (entity.Genre, error) {
	genre := entity.Genre{}

	var filter bson.M = bson.M{"id": genre_id}
	err := sr.Database.Collection(genreCollectionName).FindOne(context.Background(), filter).Decode(&genre)
	if err != nil {
		return entity.Genre{}, fmt.Errorf("GenreRepo - GetById - rows.Scan: %w", err)
	}

	return genre, nil
}

// Create -.
func (sr *GenreRepo) Create(ctx context.Context, genre_name string) error {

	guid := uuid.New().String()
	genre := entity.Genre{
		Id:   guid,
		Name: genre_name,
	}
	_, err := sr.Database.Collection(genreCollectionName).InsertOne(context.Background(), genre)
	if err != nil {
		return fmt.Errorf("GenreRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (sr *GenreRepo) Delete(ctx context.Context, genre_id string) error {
	_, err := sr.Database.Collection(genreCollectionName).DeleteOne(
		context.Background(),
		bson.M{"id": genre_id})

	if err != nil {
		return fmt.Errorf("GenreRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}

// Update -.
func (sr *GenreRepo) Update(ctx context.Context, newGenre entity.Genre) error {

	genre := entity.Genre{
		Id:   newGenre.Id,
		Name: newGenre.Name,
	}

	update := bson.M{"$set": genre}

	_, err := sr.Database.Collection(genreCollectionName).UpdateOne(
		context.Background(),
		bson.M{"id": newGenre.Id},
		update)

	if err != nil {
		fmt.Errorf("GenreRepo - Create - rows.Scan: %w", err)
	}
	return nil
}
