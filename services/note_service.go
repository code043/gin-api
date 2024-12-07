package services

import (
	"github.com/code043/gin-api/models"
	"github.com/code043/gin-api/repositories"
)

type NoteService struct {
	Repo *repositories.NoteRepository
}

func NewNoteService(repo *repositories.NoteRepository) *NoteService {
	return &NoteService{Repo: repo}
}

func (s *NoteService) Create(note *models.Note) error {
	return s.Repo.Create(note)
}
func (s *NoteService) GetAll() ([]models.Note, error) {
	return s.Repo.GetAll()
}
func (s *NoteService) GetByID(id uint) (*models.Note, error) {
	return s.Repo.GetByID(id)
}

func (s *NoteService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
