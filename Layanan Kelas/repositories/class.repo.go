package repository

import (
	"context"
	"fmt"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(conn *gorm.DB) domain.ClassRepository {
	return &classRepository{
		db: conn,
	}
}

func (c *classRepository) Create(ctx context.Context, data domain.ClassModel) (domain.ClassModel, error) {
	result := c.db.WithContext(ctx).Create(&data)
	err := result.Error
	return data, HandleRepositoryError(err)
}
func (c *classRepository) Read(ctx context.Context, query string, id uint, page uint) ([]domain.ClassModel, uint,error) {
	var results []domain.ClassModel
	var err error
	queryDB := c.db.WithContext(ctx)
	var result *gorm.DB
	var maxPage int64
	if id != 0 {
		getQuery := queryDB.
			Model(&domain.ClassModel{}).
			Where("gurus.nama LIKE ?", fmt.Sprintf("%%%s%%", query)).
			Where("tahun_ajarans.tahun_mulai = ?", query).
			Where("grup_kelas.konsen LIKE ?", fmt.Sprintf("%%%s%%", query))
		result = getQuery.Scopes(PaginateScope(page)).Find(&results)
		getQuery.Count(&maxPage)
		maxPage /= DATA_PERPAGE
	} else {
		result = queryDB.
		Where("id = ?", id).
		Find(&results)
	}

	err = result.Error
	return results, uint(maxPage) ,HandleRepositoryError(err)
}
func (c *classRepository) Update(ctx context.Context, id uint, data domain.ClassModel) (int, domain.ClassModel, error) {
	result := c.db.WithContext(ctx).
		Where("id = ?", id).
		Updates(&data)
	err := result.Error
	aff := result.RowsAffected
	if aff != 1 && err == nil {
		return 0, data, domain.ErrResourceNotFound
	}
	return int(aff), data, HandleRepositoryError(err)
}
func (c *classRepository) Delete(ctx context.Context, id uint) (int, error) {
	result := c.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.ClassModel{})
	err := result.Error
	aff := result.RowsAffected
	if aff != 1 && err == nil {
		return 0, domain.ErrResourceNotFound
	}
	return int(aff), HandleRepositoryError(err)
}