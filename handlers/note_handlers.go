package handlers

import (
	"net/http"
	"strconv"

	"github.com/code043/gin-api/models"
	"github.com/code043/gin-api/services"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	Service *services.NoteService
}

type NoteResponse struct {
	Quantity int         `json:"quantity"`
	Data     interface{} `json:"data"`
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
	if err := h.Service.Create(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}
func (h *NoteHandler) GetAll(c *gin.Context) {
	notes, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}
	response := NoteResponse{
		Quantity: len(notes),
		Data:     notes,
	}

	c.JSON(http.StatusOK, response)
}
func (h *NoteHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	note, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}
func (h *NoteHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedNote models.Note
	if err := c.ShouldBindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedNote.ID = uint(id)

	if err := h.Service.Update(&updatedNote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}
	c.JSON(http.StatusOK, updatedNote)
}

func (h *NoteHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
