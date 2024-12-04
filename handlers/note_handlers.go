package handlers

import (
	"net/http"

	"github.com/code043/gin-api/models"
	"github.com/code043/gin-api/services"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	Service *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{
		Service: service,
	}
}
func (h *NoteHandler) Create(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
