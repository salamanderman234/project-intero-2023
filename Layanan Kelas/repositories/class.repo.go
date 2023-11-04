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
	_, err := basicCreateFunc(ctx, c.db, &data)
	return data, err
}
func (c *classRepository) Read(ctx context.Context, query string, id uint, page uint) ([]domain.ClassModel, uint,error) {
	var results []domain.ClassModel
	var maxPage int64
	var err error
	queryDB := c.db.Model(&domain.ClassModel{}).WithContext(ctx)
	if id == 0 {
		searchQuery := queryDB.
			Where("gurus.nama LIKE ?", fmt.Sprintf("%%%s%%", query)).
			Or("tahun_ajarans.tahun_mulai = ?", query).
			Or("grup_kelas.konsentrasi LIKE ?", fmt.Sprintf("%%%s%%", query)).
			Joins("join gurus on gurus.id = kelas.guru_id").
			Joins("join tahun_ajarans on tahun_ajarans.id = kelas.tahun_ajaran_id").
			Joins("join grup_kelas on grup_kelas.id = kelas.grup_kelas_id")
		_, maxPage, err = basicSearchFunc(ctx, c.db, *searchQuery, page, &results)
	} else {
		findQuery := queryDB.Where("id = ?", id)
		_, _, err = basicSearchFunc(ctx, c.db, *findQuery, page, &results)
	}
	return results, uint(maxPage) ,err
}
func (c *classRepository) Update(ctx context.Context, id uint, data domain.ClassModel) (int, domain.ClassModel, error) {
	result, err := basicUpdateFunc(ctx, c.db, id, &data)
	return int(result.RowsAffected), data, err
}
func (c *classRepository) Delete(ctx context.Context, id uint) (int, error) {
	result, err := basicDeleteFunc(ctx, c.db, id, &domain.ClassModel{})
	return int(result.RowsAffected), err
}