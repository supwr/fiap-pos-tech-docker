package application

import (
	"github.com/supwr/fiap-pos-tech-docker/src/domain/contract"
	"github.com/supwr/fiap-pos-tech-docker/src/domain/entity"
)

type App struct {
	service contract.LanguageService
}

func NewApp(s contract.LanguageService) *App {
	return &App{
		service: s,
	}
}

func (a *App) ListLanguages() ([]*entity.Language, error) {
	return a.service.List()
}
