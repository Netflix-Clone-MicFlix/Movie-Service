package repositories

import (
	"context"
	"time"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const movieCollectionName = "movies"

// MovieRepo -.
type MovieRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewMovieRepo(mdb *mongodb.MongoDB) *MovieRepo {
	return &MovieRepo{mdb}
}

// GetAll -.
func (ur *MovieRepo) GetAll(ctx context.Context) ([]entity.Movie, error) {

	movies := []entity.Movie{}

	collection := ur.Database.Collection(movieCollectionName)

	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("MovieRepo - GetAll - rows.Scan: %w", err)
	}
	if err = curr.All(context.Background(), &movies); err != nil {
		return nil, fmt.Errorf("MovieRepo - GetAll - rows.Scan: %w", err)
	}

	return movies, nil
}

// GetById -.
func (ur *MovieRepo) GetById(ctx context.Context, movie_id string) (entity.Movie, error) {
	movie := entity.Movie{}

	collection := ur.Database.Collection(movieCollectionName)

	var filter bson.M = bson.M{"id": movie_id}
	if err := collection.FindOne(ctx, filter).Decode(&movie); err != nil {
		return entity.Movie{}, fmt.Errorf("MovieRepo - GetAll - rows.Scan: %w", err)
	}
	return movie, nil
}

// Create -.
func (ur *MovieRepo) Create(ctx context.Context, movie entity.Movie) error {
	guid := uuid.New().String()
	movie.Id = guid
	movie.CreatedAt = time.Now()

	_, err := ur.Database.Collection(movieCollectionName).InsertOne(context.Background(), movie)
	if err != nil {
		return fmt.Errorf("MovieRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Update -.
func (ur *MovieRepo) Update(ctx context.Context, movie_id string, movie entity.Movie) error {

	movie.Id = movie_id
	update := bson.M{"$set": movie}

	_, err := ur.Database.Collection(movieCollectionName).UpdateOne(
		context.Background(),
		bson.M{"id": movie_id},
		update)

	if err != nil {
		return fmt.Errorf("MovieRepo - Update - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (ur *MovieRepo) Delete(ctx context.Context, movie_id string) error {
	_, err := ur.Database.Collection(movieCollectionName).DeleteOne(
		context.Background(),
		bson.M{"id": movie_id})

	if err != nil {
		return fmt.Errorf("MovieRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}
