// Package usecase implements application business logic. Each logic group in own file.
package internal

import (
	"context"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Movie -.
	Movie interface {
		GetById(context.Context, string) (entity.Movie, error)
		GetAll(context.Context) ([]entity.Movie, error)
	}

	// MovieRepo -.
	MovieRepo interface {
		GetAll(context.Context) ([]entity.Movie, error)
		GetById(context.Context, string) (entity.Movie, error)
		Create(context.Context, entity.Movie) error
		Update(context.Context, string, entity.Movie) error
		Delete(context.Context, string) error
	}

	// GenreRepo -.
	GenreRepo interface {
		GetById(context.Context, string) (entity.Genre, error)
		Create(context.Context, string) error
		Delete(context.Context, string) error
		Update(context.Context, entity.Genre) error
	}
	// MovieWebAPI -.
	WebAPI interface {
	}
)
