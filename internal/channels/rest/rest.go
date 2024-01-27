package rest

import (
	"context"
	"fmt"
	"net/http"
	"tech-challenge-auth/internal/config"
	"tech-challenge-auth/internal/middlewares"
	"tech-challenge-auth/internal/service"

	"github.com/labstack/echo"
)

var (
	cfg = &config.Cfg
)

type Login interface {
	RegisterGroup(*echo.Group)
	Login(echo.Context) error
	Bypass(echo.Context) error
	Start() error
}

type login struct {
	service service.LoginService
}

func NewRestChannel(svc service.LoginService) Login {
	return &login{
		service: svc,
	}
}

func (u *login) RegisterGroup(g *echo.Group) {
	g.POST("/login", u.Login)
	g.POST("/bypass", u.Bypass)
}

func (u *login) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainGroup := router.Group("/api")

	authGroup := mainGroup.Group("/user")
	u.RegisterGroup(authGroup)

	return router.Start(":" + cfg.Server.Port)
}

func (u *login) Login(c echo.Context) error {
	var userLogin LoginRequest

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	token, err := u.service.Login(c.Request().Context(), userLogin.toCanonical())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response{err.Error()})
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}

func (u *login) Bypass(c echo.Context) error {
	token, _ := u.service.Bypass(context.Background())
	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}
