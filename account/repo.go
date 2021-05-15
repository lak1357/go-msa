package account

import (
	"context"
	"errors"

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

func (r *repo) CreateUser(ctx context.Context, user User) error {
	return ErrRepo
}

func (r *repo) GetUser(ctx context.Context, id string) (string, error) {
	return "lak1357@gmail.com", nil
}
