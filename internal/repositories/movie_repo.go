package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/mongodb"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/security"
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

	var filter bson.M = bson.M{"_id": movie_id}
	curr, err := ur.Database.Collection(movieCollectionName).Find(context.Background(), filter)
	if err != nil {
		return entity.Movie{}, fmt.Errorf("MovieRepo - GetById - rows.Scan: %w", err)
	}
	defer curr.Close(context.Background())

	curr.All(context.Background(), &movie)

	return movie, nil
}

// Create -.
func (ur *MovieRepo) Create(ctx context.Context, movie entity.Movie, salt []byte) error {
	var hashedPassword = security.HashPassword(movie.Password, salt)

	movie.Password = hashedPassword
	_, err := ur.Database.Collection(movieCollectionName).InsertOne(context.Background(), movie)
	if err != nil {
		return fmt.Errorf("MovieRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Update -.
func (ur *MovieRepo) Update(ctx context.Context, movie_id string, movie entity.Movie, salt []byte) error {

	var hashedPassword = security.HashPassword(movie.Password, salt)
	movie.Password = hashedPassword
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
		bson.M{"_id": movie_id})

	if err != nil {
		return fmt.Errorf("MovieRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}

// Login -.
func (ur *MovieRepo) Login(ctx context.Context, movie entity.Movie) (entity.Movie, error) {
	moviedb := entity.Movie{}

	var filter bson.M = bson.M{"email": movie.Email}
	err := ur.Database.Collection(movieCollectionName).FindOne(context.Background(), filter).Decode(&moviedb)
	if err != nil {
		return entity.Movie{}, fmt.Errorf("MovieRepo - Login - rows.Scan: %w", err)
	}

	return moviedb, nil
}
