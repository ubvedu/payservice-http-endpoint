package repository

import "github.com/google/uuid"

type Users interface {
    AddUser(email, password string) uuid.UUID
    HasId(id uuid.UUID) bool
}
