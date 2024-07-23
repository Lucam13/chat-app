package service

import (
	"context"

	"github.com/lean1097/chat-back/internal/chatapp/repository"
	"github.com/lean1097/chat-back/internal/chatapp/user"
	"github.com/lean1097/chat-back/internal/platform/errors"
)

type (
	UserService interface {
		Get(ctx context.Context) ([]user.User, error)
		GetByID(ctx context.Context, userID int64) (user.User, error)
		Save(ctx context.Context, username string, rol user.Rol, areaID int64) error
		Delete(ctx context.Context, userID int64) error
	}

	userService struct {
		repo repository.UserRepository
	}
)

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Get(ctx context.Context) ([]user.User, error) {
	return s.repo.Get(ctx)
}

func (s *userService) GetByID(ctx context.Context, userID int64) (user.User, error) {
	return s.repo.GetByID(ctx, userID)
}

func (s *userService) Save(ctx context.Context, username string, rol user.Rol, areaID int64) error {
	if !user.IsValidRol(rol) {
		return errors.ErrInvalidRol
	}

	return s.repo.Save(ctx, username, rol, areaID)
}

func (s *userService) Delete(ctx context.Context, userID int64) error {
	return s.repo.Delete(ctx, userID)
}
