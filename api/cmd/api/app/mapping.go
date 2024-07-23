package app

import "github.com/gin-gonic/gin"

func configureMappings(router *gin.Engine, handlers handlerContainer) {
	baseGroup := router.Group("/api")

	// Chat web socket operations.
	chatWebSocketGroup := baseGroup.Group("/chat")
	chatWebSocketGroup.GET("/ws", handlers.chatWebSocketHandler.Handle())

	// Chat operations.
	chatGroup := baseGroup.Group("/chats")
	chatGroup.GET("", handlers.chatHandler.Get())
	chatGroup.GET("/:id", handlers.chatHandler.GetByAreaID())
	chatGroup.POST("", handlers.chatHandler.Save())
	chatGroup.DELETE("/:id", handlers.chatHandler.Delete())

	// Area operations.
	areaGroup := baseGroup.Group("/areas")
	areaGroup.GET("", handlers.areaHandler.Get())
	areaGroup.GET("/:id", handlers.areaHandler.GetByID())
	areaGroup.POST("", handlers.areaHandler.Save())
	areaGroup.DELETE("/:id", handlers.areaHandler.Delete())

	// Message operations.
	messageGroup := baseGroup.Group("/messages")
	messageGroup.GET("/:chat_id", handlers.messageHandler.GetByChatID())
	messageGroup.POST("", handlers.messageHandler.Save())

	// User operations.
	userGroup := baseGroup.Group("/users")
	userGroup.GET("", handlers.userHandler.Get())
	userGroup.GET("/:id", handlers.userHandler.GetByID())
	userGroup.POST("", handlers.userHandler.Save())
	userGroup.DELETE("/:id", handlers.userHandler.Delete())
}
