package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}

}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {

	logger := log.With(s.logger, "method", "CreateUser")
	uuid, _ := uuid.NewV4()
	uid := uuid.String()
	user := User{
		UUID:     uid,
		Email:    email,
		Password: password,
	}

	id, err := s.repository.CreateUser(ctx, user)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", uid)

	return id, nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {

	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("get user", email)

	return email, nil
}
