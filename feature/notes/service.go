package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/entity"
)

type service struct {
	repository NotesRepository
}

type NotesService interface {
	GetNotes() ([]entity.Notes, error)
}

func (svc *service) GetNotes() ([]entity.Notes, error) {
	return svc.repository.GetNotes()
}
