package app

import (
	"github.com/lean1097/chat-back/cmd/api/app/handler"
	areaService "github.com/lean1097/chat-back/internal/chatapp/area/service"
	chatService "github.com/lean1097/chat-back/internal/chatapp/chat/service"
	messageService "github.com/lean1097/chat-back/internal/chatapp/message/service"
	"github.com/lean1097/chat-back/internal/chatapp/repository"
	userService "github.com/lean1097/chat-back/internal/chatapp/user/service"
	"github.com/lean1097/chat-back/internal/platform/storage/sql"
)

type handlerContainer struct {
	chatWebSocketHandler handler.ChatWebSocketHandler
	chatHandler          handler.ChatHandler
	areaHandler          handler.AreaHandler
	messageHandler       handler.MessageHandler
	userHandler          handler.UserHandler
}

func buildHandlers() handlerContainer {
	dbCredentials := sql.BuildDatabaseCredentials()
	dbInitializer := sql.NewDatabaseInitializer()

	sqlDatabaseService := sql.NewSQLDatabaseService(dbCredentials, dbInitializer)
	sqlDatabase, err := sqlDatabaseService.StartSQLDatabase()
	if err != nil {
		panic(err)
	}

	chatRepository := repository.NewChatRepository(sqlDatabase.GetDatabase())
	messageRepository := repository.NewMessageRepository(sqlDatabase.GetDatabase())
	userRepository := repository.NewUserRepository(sqlDatabase.GetDatabase())
	areaRepository := repository.NewAreaRepository(sqlDatabase.GetDatabase())

	chatService := chatService.NewChatService(chatRepository)
	messageService := messageService.NewMessageService(messageRepository)
	userService := userService.NewUserService(userRepository)
	areaService := areaService.NewAreaService(areaRepository)

	chatWebSocketHandler := handler.NewChatWebSocketHandler(messageService)
	chatHandler := handler.NewChatHandler(chatService)
	areaHandler := handler.NewAreaHandler(areaService)
	messageHandler := handler.NewMessageHandler(messageService)
	userHandler := handler.NewUserHandler(userService)

	return handlerContainer{
		chatWebSocketHandler: chatWebSocketHandler,
		chatHandler:          chatHandler,
		areaHandler:          areaHandler,
		messageHandler:       messageHandler,
		userHandler:          userHandler,
	}
}
