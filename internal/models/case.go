package models

import (
	"time"

	"github.com/google/uuid"
)

type CaseStatus string

const (
	StatusOpen   CaseStatus = "open"
	StatusClosed CaseStatus = "closed"
)

type Case struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Title       string     `json:"title" db:"title" validate:"required"`
	Description string     `json:"description" db:"description"`
	Status      CaseStatus `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewCase(title, description string) *Case {
	return &Case{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      StatusOpen,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
