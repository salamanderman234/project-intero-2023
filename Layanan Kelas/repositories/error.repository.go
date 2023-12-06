package repository

import (
	"errors"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

func handleRepositoryError(result *gorm.DB) error {
	err := result.Error
	// fmt.Println(err)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return domain.ErrDuplicateEnties
	}
	if errors.Is(err, gorm.ErrRecordNotFound){
		return domain.ErrResourceNotFound
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return domain.ErrForeignKeyViolated
	}
	if err != nil {
		return domain.ErrGormInGeneral
	}
	return nil
}