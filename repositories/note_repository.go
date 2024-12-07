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
func (r *NoteRepository) GetAll() ([]models.Note, error) {
	var notes []models.Note
	err := r.DB.Find(&notes).Error
	return notes, err
}
func (r *NoteRepository) GetByID(id uint) (*models.Note, error) {
	var note models.Note
	err := r.DB.First(&note, id).Error
	return &note, err
}
func (r *NoteRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Note{}, id).Error
}
