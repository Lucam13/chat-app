package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                             // Permite solicitudes desde este origen
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Encabezados permitidos
		AllowCredentials: true,
	}))

	handlers := buildHandlers()

	configureMappings(router, handlers)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
