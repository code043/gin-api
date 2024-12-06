package routes

import (
	"github.com/code043/gin-api/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, noteHandler *handlers.NoteHandler) {
	r.POST("/notes", noteHandler.Create)
	r.GET("/notes", noteHandler.GetAll)
	r.GET("/notes/:id", noteHandler.GetByID)
}
