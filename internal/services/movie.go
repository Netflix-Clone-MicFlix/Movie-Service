package services

import (
	"context"
	"fmt"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
)

// MovieUseCase -.
type MovieUseCase struct {
	movieRepo internal.MovieRepo
	genreRepo internal.GenreRepo
	webAPI    internal.WebAPI
}

// New -.
func NewMovieUseCase(r internal.MovieRepo, s internal.GenreRepo, w internal.WebAPI) *MovieUseCase {
	return &MovieUseCase{
		movieRepo: r,
		genreRepo: s,
		webAPI:    w,
	}
}

// GetById - gets all movie by ID -.
func (uc *MovieUseCase) GetById(ctx context.Context, movie_id string) (entity.Movie, error) {
	movie, err := uc.movieRepo.GetById(ctx, movie_id)
	if err != nil {
		return entity.Movie{}, fmt.Errorf("MovieUseCase - GetById - s.movieRepo.GetHistory: %w", err)
	}

	return movie, nil
}

// GetAll - gets alls-.
func (uc *MovieUseCase) GetAll(ctx context.Context) ([]entity.Movie, error) {

	movies, err := uc.movieRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("MovieUseCase - GetAll - s.movieRepo.Store: %w", err)
	}

	return movies, nil
}
