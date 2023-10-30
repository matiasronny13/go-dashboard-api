package notes

import (
	model "github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/google/uuid"
)

type service struct {
	repository NotesRepository
}

type NotesService interface {
	GetNotes() ([]model.Notes, error)
	GetNotesById(id uuid.UUID) (model.Notes, error)
	CreateNotes(model *model.Notes) error
	UpdateNotes(model *model.Notes) error
	DeleteNotes(id uuid.UUID) error
}

func (svc *service) GetNotes() ([]model.Notes, error) {
	return svc.repository.GetNotes()
}

func (svc *service) GetNotesById(id uuid.UUID) (model.Notes, error) {
	return svc.repository.GetNotesById(id)
}

func (svc *service) CreateNotes(model *model.Notes) error {
	return svc.repository.CreateNotes(model)
}

func (svc *service) UpdateNotes(model *model.Notes) error {
	return svc.repository.UpdateNotes(model)
}

func (svc *service) DeleteNotes(id uuid.UUID) error {
	return svc.repository.DeleteNotes(id)
}
