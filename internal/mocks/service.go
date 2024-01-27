package mocks

import (
	"context"
	"tech-challenge-auth/internal/canonical"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (c *MockService) Login(_ context.Context, login canonical.Login) (string, error) {
	args := c.Called(login)
	return args.Get(0).(string), args.Error(1)
}

func (c *MockService) MockLogin(login canonical.Login, token string, errorToReturn error, times int) {
	c.On("Login", mock.MatchedBy(func(l canonical.Login) bool {
		return l.Login == login.Login
	})).Return(token, errorToReturn).Times(times)
}

func (c *MockService) Bypass(_ context.Context) (string, error) {
	return "fakebypasstoken", nil
}

func (c *MockService) MockBypass(token string, error error, times int) {
	c.On("Bypass").Return(token, error).Times(times)
}
