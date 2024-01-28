package main

import (
	"tech-challenge-auth/internal/channels/rest"
	"tech-challenge-auth/internal/config"
	"tech-challenge-auth/internal/repositories"
	"tech-challenge-auth/internal/service"

	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	service := service.NewLoginService(repositories.NewUserRepo(repositories.New()))
	if err := rest.NewRestChannel(service).Start(); err != nil {
		logrus.Panic()
	}
}
