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

// Create - Create genre-.
func (uc *MovieUseCase) Create(ctx context.Context, title string, discription string) error {

	data := entity.Movie{
		Title:       title,
		Description: discription,
		Genres:      []string{},
	}

	if data.Title == "" || data.Description == "" {
		return fmt.Errorf("MovieUseCase - Create - s.movieRepo.Store: No values provided")
	}

	err := uc.movieRepo.Create(context.Background(), data)
	if err != nil {
		return fmt.Errorf("MovieUseCase - Create - s.movieRepo.Store: %w", err)
	}

	return nil
}

// Create - gets alls-.
func (uc *MovieUseCase) AddGenre(ctx context.Context, name string) error {

	if name == "" {
		return fmt.Errorf("MovieUseCase - Create - s.movieRepo.Store: No values provided")
	}

	err := uc.genreRepo.Create(context.Background(), name)
	if err != nil {
		return fmt.Errorf("MovieUseCase - Create - s.movieRepo.Store: %w", err)
	}

	return nil
}

// GetById - gets all movie by ID -.
func (uc *MovieUseCase) GetGenreById(ctx context.Context, genre_id string) (entity.Genre, error) {
	genre, err := uc.genreRepo.GetById(ctx, genre_id)
	if err != nil {
		return entity.Genre{}, fmt.Errorf("MovieUseCase - GetById - s.movieRepo.GetHistory: %w", err)
	}

	return genre, nil
}

// GetAll - gets alls-.
func (uc *MovieUseCase) GetAllGenre(ctx context.Context) ([]entity.Genre, error) {

	genres, err := uc.genreRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("MovieUseCase - GetAll - s.movieRepo.Store: %w", err)
	}

	return genres, nil
}
