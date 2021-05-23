package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID     string `json:"uuid:omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
