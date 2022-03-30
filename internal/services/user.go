package services

import (
	"context"
	"fmt"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/security"
	"github.com/google/uuid"
)

// MovieUseCase -.
type MovieUseCase struct {
	userRepo internal.MovieRepo
	saltRepo internal.SaltRepo
	webAPI   internal.WebAPI
}

// New -.
func NewMovieUseCase(r internal.MovieRepo, s internal.SaltRepo, w internal.WebAPI) *MovieUseCase {
	return &MovieUseCase{
		userRepo: r,
		saltRepo: s,
		webAPI:   w,
	}
}

// GetById - gets all user by ID -.
func (uc *MovieUseCase) GetById(ctx context.Context, user_id string) (entity.Movie, error) {
	user, err := uc.userRepo.GetById(ctx, user_id)
	if err != nil {
		return entity.Movie{}, fmt.Errorf("MovieUseCase - GetById - s.userRepo.GetHistory: %w", err)
	}

	return user, nil
}

// GetAll - gets alls-.
func (uc *MovieUseCase) GetAll(ctx context.Context) ([]entity.Movie, error) {

	users, err := uc.userRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("MovieUseCase - GetAll - s.userRepo.Store: %w", err)
	}

	return users, nil
}

// Register - gets alls-.
func (uc *MovieUseCase) Register(ctx context.Context, user entity.Movie) error {

	user.Id = uuid.New().String() //generate guid

	salt, err := uc.saltRepo.Create(context.Background(), user.Id) //generate salt

	if err != nil {
		return fmt.Errorf("MovieUseCase - Register - s.userRepo.Store: %w", err)
	}

	err = uc.userRepo.Create(context.Background(), user, salt)
	return err
}

// Login - gets alls-.
func (uc *MovieUseCase) Login(ctx context.Context, user entity.Movie) error {

	userdb, err := uc.userRepo.Login(context.Background(), user)
	if err != nil {
		return fmt.Errorf("MovieUseCase - Login - s.userRepo.Store: %w", err)
	}

	salt, err := uc.saltRepo.GetById(context.Background(), userdb.Id)
	if err != nil {
		return fmt.Errorf("MovieUseCase - Login - s.userRepo.Store: %w", err)
	}

	if userdb.Email != user.Email && !security.CheckPasswordsMatch(userdb.Password, user.Password, salt.SaltData) {
		return fmt.Errorf("MovieRepo - Login Email- rows.Scan: %w", err)
	}

	updatedSalt, err := uc.saltRepo.Update(context.Background(), userdb.Id)
	if err != nil {
		return fmt.Errorf("MovieUseCase - Login - s.userRepo.Store: %w", err)
	}

	err = uc.userRepo.Update(context.Background(), userdb.Id, user, updatedSalt)
	if err != nil {
		return fmt.Errorf("MovieUseCase - Login - s.userRepo.Store: %w", err)
	}

	return err
}
