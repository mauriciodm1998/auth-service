package service

import (
	"context"
	"fmt"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/repositories"
	"tech-challenge-auth/internal/security"
	"tech-challenge-auth/internal/token"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type LoginService interface {
	Login(context.Context, canonical.Login) (string, error)
	Bypass(context.Context) (string, error)
}

type loginService struct {
	repositories.LoginRepository
}

func NewLoginService(repo repositories.LoginRepository) LoginService {
	return &loginService{
		repo,
	}
}

func (u *loginService) Login(ctx context.Context, user canonical.Login) (string, error) {
	baseUser, err := u.LoginRepository.GetUserByLogin(ctx, user.Login)
	if err != nil {
		err = fmt.Errorf("error getting user by email: %w", err)
		logrus.WithError(err).Info()
		return "", err
	}

	if err = security.CheckPassword(baseUser.Password, user.Password); err != nil {
		err = fmt.Errorf("error checking password: %w", err)
		logrus.WithError(err).Info()
		return "", err
	}

	if err = u.SaveAccess(ctx, canonical.Access{
		ID:            uuid.New().String(),
		AccessLevelID: baseUser.AccessLevelID,
		USERID:        baseUser.Id,
		AccessedAt:    time.Now(),
	}); err != nil {
		logrus.WithError(err).Error("Error when save access in the database")
		return "", err
	}

	token, _ := token.GenerateToken(baseUser.Id)

	return token, nil
}

func (u *loginService) Bypass(_ context.Context) (string, error) {
	return token.GenerateToken("")
}
