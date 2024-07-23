package service

import (
	"context"

	"github.com/lean1097/chat-back/internal/chatapp/message"
	"github.com/lean1097/chat-back/internal/chatapp/repository"
)

type (
	MessageService interface {
		GetByChatID(ctx context.Context, chatID int64) ([]message.Message, error)
		Save(ctx context.Context, text string, userID, chatID int64) error
	}

	messageService struct {
		repo repository.MessageRepository
	}
)

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{
		repo: repo,
	}
}

func (s *messageService) GetByChatID(ctx context.Context, chatID int64) ([]message.Message, error) {
	return s.repo.GetByChatID(ctx, chatID)
}

func (s *messageService) Save(ctx context.Context, text string, userID, chatID int64) error {
	return s.repo.Save(ctx, text, userID, chatID)
}
