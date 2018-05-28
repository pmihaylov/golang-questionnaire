package models

import (
	"github.com/google/uuid"
)

type Result struct {
	BaseModel
	ResultID uuid.UUID `json:"resultId"`
	Title    string    `json:"title"`
}
