package repository

import (
	"errors"
	"github.com/supwr/fiap-pos-tech-docker/src/domain/entity"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{
		db: db,
	}
}

func (r *LanguageRepository) List() ([]*entity.Language, error) {
	var languages []*entity.Language

	if err := r.db.Order("name asc").Find(&languages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New("an error occurred from repository")
	}

	return languages, nil
}
