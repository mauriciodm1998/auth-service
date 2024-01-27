package service

import (
	"context"
	"errors"
	"fmt"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	"golang.org/x/crypto/bcrypt"
)

var (
	service LoginService

	repositoryMock *mocks.LoginRepositoryMock
)

func init() {
	repositoryMock = new(mocks.LoginRepositoryMock)
	service = NewLoginService(repositoryMock)
}

func TestLogin(t *testing.T) {
	login := canonical.Login{
		Login:    "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:            "123123",
		Login:         "fulano",
		Password:      "$2a$10$GjN8aPVbp5u/jFZFwmgda.XpLnj7oqb6hPsN1v57JhbCIqN/M.04O",
		AccessLevelID: 1,
	}

	repositoryMock.MockGetUserByLogin(login.Login, user, nil, 1)
	repositoryMock.MockSaveAccess(canonical.Access{
		USERID: "123123",
	}, nil, 1)

	token, err := service.Login(context.Background(), login)

	assert.Nil(t, err)
	assert.NotNil(t, token)
	repositoryMock.AssertExpectations(t)
}

func TestLoginGetUserError(t *testing.T) {
	login := canonical.Login{
		Login:    "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:            "123123",
		Login:         "fulano",
		Password:      "$2a$10$GjN8aPVbp5u/jFZFwmgda.XpLnj7oqb6hPsN1v57JhbCIqN/M.04O",
		AccessLevelID: 1,
	}

	repositoryMock.MockGetUserByLogin(login.Login, user, errors.New("generic error"), 1)

	_, err := service.Login(context.Background(), login)

	assert.Equal(t, fmt.Errorf("error getting user by email: %w", errors.New("generic error")), err)
}

func TestLoginCheckPasswordError(t *testing.T) {
	login := canonical.Login{
		Login:    "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:            "123123",
		Login:         "fulano",
		Password:      "$2a$10$GjN8aPVbp5u/jFZFwmg/M.04O",
		AccessLevelID: 1,
	}

	patch, _ := mpatch.PatchMethod(bcrypt.CompareHashAndPassword, func(hashedPassword []byte, password []byte) error {
		return errors.New("generic error")
	})

	defer patch.Unpatch()

	repositoryMock.MockGetUserByLogin(login.Login, user, nil, 1)

	_, err := service.Login(context.Background(), login)

	assert.Equal(t, fmt.Errorf("error checking password: %w", errors.New("generic error")), err)
}

func TestLoginSaveAccessError(t *testing.T) {
	login := canonical.Login{
		Login:    "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:            "123123",
		Login:         "fulano",
		Password:      "$2a$10$GjN8aPVbp5u/jFZFwmgda.XpLnj7oqb6hPsN1v57JhbCIqN/M.04O",
		AccessLevelID: 1,
	}

	repositoryMock.MockGetUserByLogin(login.Login, user, nil, 1)
	repositoryMock.MockSaveAccess(canonical.Access{
		USERID: "123123",
	}, errors.New("generic error"), 1)

	_, err := service.Login(context.Background(), login)

	assert.Error(t, err)
	repositoryMock.AssertExpectations(t)
}

func TestBypass(t *testing.T) {
	result, _ := service.Bypass(context.Background())

	assert.NotNil(t, result)
}
