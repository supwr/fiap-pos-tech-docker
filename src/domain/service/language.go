package service

import (
	"github.com/supwr/fiap-pos-tech-docker/src/domain/contract"
	"github.com/supwr/fiap-pos-tech-docker/src/domain/entity"
)

type LanguageService struct {
	repository contract.LanguageRepository
}

func NewLanguageService(r contract.LanguageRepository) *LanguageService {
	return &LanguageService{
		repository: r,
	}
}

func (s *LanguageService) List() ([]*entity.Language, error) {
	return s.repository.List()
}
