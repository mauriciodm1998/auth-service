package mocks

import (
	"context"
	"tech-challenge-auth/internal/canonical"

	"github.com/stretchr/testify/mock"
)

type LoginRepositoryMock struct {
	mock.Mock
}

func (l *LoginRepositoryMock) GetUserByLogin(_ context.Context, login string) (*canonical.User, error) {
	args := l.Called(login)

	return args.Get(0).(*canonical.User), args.Error(1)
}

func (l *LoginRepositoryMock) MockGetUserByLogin(login string, user canonical.User, errorReturned error, times int) {
	l.On("GetUserByLogin", mock.MatchedBy(func(u string) bool {
		return u == login
	})).Return(&user, errorReturned).Times(times)
}

func (l *LoginRepositoryMock) SaveAccess(_ context.Context, access canonical.Access) error {
	args := l.Called(access)

	return args.Error(0)
}

func (l *LoginRepositoryMock) MockSaveAccess(access canonical.Access, errorReturned error, times int) {
	l.On("SaveAccess", mock.MatchedBy(func(u canonical.Access) bool {
		return u.USERID == access.USERID
	})).Return(errorReturned).Times(times)
}
