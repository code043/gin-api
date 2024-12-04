package repositories

import (
	"github.com/code043/gin-api/models"
	"gorm.io/gorm"
)

type NoteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{
		DB: db,
	}
}
func (r *NoteRepository) Create(note *models.Note) error {
	return r.DB.Create(note).Error
}
