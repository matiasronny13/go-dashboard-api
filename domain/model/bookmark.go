package model

import (
	"github.com/google/uuid"
)

type Bookmark struct {
	Id      *uuid.UUID `json:"id" format:"uuid"`
	Name    string     `json:"name"`
	URL     string     `json:"url"`
	GroupID int64      `json:"groupId"`
	Icon    string     `json:"icon"`
}
