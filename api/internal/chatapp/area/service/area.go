package service

import (
	"context"

	"github.com/lean1097/chat-back/internal/chatapp/area"
	"github.com/lean1097/chat-back/internal/chatapp/repository"
)

type (
	AreaService interface {
		Get(ctx context.Context) ([]area.Area, error)
		GetByID(ctx context.Context, areaID int64) (area.Area, error)
		Save(ctx context.Context, name string) error
		Delete(ctx context.Context, areaID int64) error
	}

	areaService struct {
		repo repository.AreaRepository
	}
)

func NewAreaService(repo repository.AreaRepository) AreaService {
	return &areaService{
		repo: repo,
	}
}

func (s *areaService) Get(ctx context.Context) ([]area.Area, error) {
	return s.repo.Get(ctx)
}

func (s *areaService) GetByID(ctx context.Context, areaID int64) (area.Area, error) {
	return s.repo.GetByID(ctx, areaID)
}

func (s *areaService) Save(ctx context.Context, name string) error {
	return s.repo.Save(ctx, name)
}

func (s *areaService) Delete(ctx context.Context, areaID int64) error {
	return s.repo.Delete(ctx, areaID)
}
