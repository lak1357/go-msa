package user

import (
	"context"
	"errors"

	"example.com/go-msa/database"
	"github.com/go-kit/kit/log"
)

var ErrRepo = errors.New("unable to handle repo request")

type repo struct {
	logger log.Logger
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r *repo) CreateUser(ctx context.Context, user User) (string, error) {
	db := database.DBcon
	db.Create(&user)
	return user.UUID, nil
}

func (r *repo) GetUser(ctx context.Context, uid string) (string, error) {
	db := database.DBcon
	var user User
	db.Where("UUID = ?", uid).Find(&user)
	return user.Email, nil
}
