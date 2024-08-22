package models

import (
	"github.com/google/uuid"
)

type TokenPayload struct {
	ID		uuid.UUID		`json:"id"`
	Email	string			`json:"email"`
}