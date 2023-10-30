package model

import "github.com/google/uuid"

type Notes struct {
	Id      *uuid.UUID `json:"id" format:"uuid"`
	Title   string     `json:"title"`
	Content string     `json:"content"`
}
