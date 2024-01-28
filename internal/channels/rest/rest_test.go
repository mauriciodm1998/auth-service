package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/mocks"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	svc *mocks.MockService

	rest Login
)

func init() {
	svc = new(mocks.MockService)

	rest = NewRestChannel(svc)
}

func TestStart(t *testing.T) {
	go func() {
		err := rest.Start()
		assert.NoError(t, err)
	}()

	<-time.After(100 * time.Millisecond)
}

func TestLogin(t *testing.T) {
	login := canonical.Login{
		Login:    "loginfulano",
		Password: "12345",
	}

	svc.MockLogin(login, "fake-token", nil, 1)

	request := LoginRequest{
		Login:    "loginfulano",
		Password: "12345",
	}

	req := createJsonRequest(http.MethodPost, "/login", request)

	rec := httptest.NewRecorder()

	err := rest.Login(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	svc.AssertExpectations(t)
}

func TestLoginError(t *testing.T) {
	login := canonical.Login{
		Login:    "loginfulano",
		Password: "12345",
	}

	svc.MockLogin(login, "fake-token", errors.New("generic error"), 1)

	request := LoginRequest{
		Login:    "loginfulano",
		Password: "12345",
	}

	req := createJsonRequest(http.MethodPost, "/login", request)

	rec := httptest.NewRecorder()

	err := rest.Login(echo.New().NewContext(req, rec))

	assert.Error(t, err)
	svc.AssertExpectations(t)
}

func TestLoginErrorBind(t *testing.T) {
	req := createJsonRequest(http.MethodPost, "/login", "")

	rec := httptest.NewRecorder()

	err := rest.Login(echo.New().NewContext(req, rec))

	assert.Equal(t, "code=400, message={invalid data}", err.Error())
	svc.AssertExpectations(t)
}

func TestBypass(t *testing.T) {
	expectedResult := `{"token":"fakebypasstoken"}
`
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	svc.MockBypass("fakebypasstoken", nil, 1)

	err := rest.Bypass(echo.New().NewContext(req, rec))

	assert.Equal(t, expectedResult, rec.Body.String())
	assert.Nil(t, err)
}

func createJsonRequest(method, endpoint string, request interface{}) *http.Request {
	json, _ := json.Marshal(request)
	req := httptest.NewRequest(method, endpoint, bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	return req
}
