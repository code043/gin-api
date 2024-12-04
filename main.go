package main

import (
	"github.com/code043/gin-api/handlers"
	"github.com/code043/gin-api/models"
	"github.com/code043/gin-api/repositories"
	"github.com/code043/gin-api/routes"
	"github.com/code043/gin-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Note{})

	repo := repositories.NewNoteRepository(db)
	service := services.NewNoteService(repo)
	handler := handlers.NewNoteHandler(service)

	r := gin.Default()
	routes.Routes(r, handler)
	r.Run(":3000")
}
