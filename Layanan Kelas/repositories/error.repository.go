package repository

import (
	"errors"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

func HandleRepositoryError(err error) error {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return domain.ErrDuplicateEnties
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrResourceNotFound
	}
	if err != nil {
		return domain.ErrGormInGeneral
	}
	return nil
}