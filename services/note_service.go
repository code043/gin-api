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
