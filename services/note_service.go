package services

import (
	"github.com/code043/gin-api/models"
	"github.com/code043/gin-api/repositories"
)

type NoteService struct {
	Repo *repositories.NoteRepository
}

func (s *NoteService) Create(note *models.Note) error {
	return s.Repo.Create(note)
}
