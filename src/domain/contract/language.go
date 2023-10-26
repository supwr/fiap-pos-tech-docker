package contract

import "github.com/supwr/fiap-pos-tech-docker/src/domain/entity"

type LanguageService interface {
	List() ([]*entity.Language, error)
}

type LanguageRepository interface {
	List() ([]*entity.Language, error)
}

type LanguageApplication interface {
	ListLanguages() ([]*entity.Language, error)
}
