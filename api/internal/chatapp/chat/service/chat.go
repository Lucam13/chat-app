package service

import (
	"context"

	"github.com/lean1097/chat-back/internal/chatapp/chat"
	"github.com/lean1097/chat-back/internal/chatapp/repository"
)

type (
	ChatService interface {
		Get(ctx context.Context) (map[string][]chat.Chat, error)
		GetByAreaID(ctx context.Context, areaID int64) ([]chat.Chat, error)
		Save(ctx context.Context, fromAreaID, toAreaID int64) error
		Delete(ctx context.Context, chatID int64) error
	}

	chatService struct {
		repo repository.ChatRepository
	}
)

func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatService{
		repo: repo,
	}
}

func (s *chatService) Get(ctx context.Context) (map[string][]chat.Chat, error) {
	return s.repo.Get(ctx)
}

func (s *chatService) GetByAreaID(ctx context.Context, areaID int64) ([]chat.Chat, error) {
	return s.repo.GetByAreaID(ctx, areaID)
}

func (s *chatService) Save(ctx context.Context, fromAreaID, toAreaID int64) error {
	return s.repo.Save(ctx, fromAreaID, toAreaID)
}

func (s *chatService) Delete(ctx context.Context, chatID int64) error {
	return s.repo.Delete(ctx, chatID)
}
