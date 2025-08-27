package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vantonietti/gin-testing/internal/handler"
	"github.com/vantonietti/gin-testing/internal/infrastructure"
	"github.com/vantonietti/gin-testing/internal/usecase"
)

func main() {
	dsn := "postgres://username:password@localhost:5432/clean_arch_db?sslmode=diable"
	db := infrastructure.NewPostgresDB(dsn)

	userRepo := infrastructure.NewUserRepositoryPostgres(db)

	userUC := usecase.NewUserUsecase(userRepo)

	userHandler := handler.NewUserHandler(userUC)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)

	log.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
