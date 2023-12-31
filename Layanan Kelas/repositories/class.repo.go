package repository

import (
	"context"

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
func (c *classRepository) Read(ctx context.Context, query string, id uint, page uint, orderBy string, orderWith string) ([]domain.ClassModel, uint, error) {
	var results []domain.ClassModel
	var maxPage int64
	var err error
	queryDB := c.db.Model(&domain.ClassModel{}).WithContext(ctx)
	if id != 0 {
		findQuery := queryDB.Where("id = ?", id)
		_, _, err = basicSearchFunc(ctx, c.db, *findQuery, page, orderBy, orderWith, domain.ClassModel{}, &results)
	} else {
		searchQuery := queryDB.
			Where("focuses.focus LIKE ?", "%"+query+"%").
			Or("grades.grade = ?", "%"+query+"%").
			Or("classes.group LIKE ?", "%"+query+"%").
			Joins("join focuses on focuses.id = classes.focus_id").
			Joins("join grades on grades.id = classes.grade_id")
		_, maxPage, err = basicSearchFunc(ctx, c.db, *searchQuery, page, orderBy, orderWith, domain.ClassModel{}, &results)
	}
	return results, uint(maxPage), err
}
func (c *classRepository) Update(ctx context.Context, id uint, data domain.ClassModel) (int, domain.ClassModel, error) {
	result, err := basicUpdateFunc(ctx, c.db, id, &data)
	return int(result.RowsAffected), data, err
}
func (c *classRepository) Delete(ctx context.Context, id uint) (int, error) {
	result, err := basicDeleteFunc(ctx, c.db, id, &domain.ClassModel{})
	return int(result.RowsAffected), err
}

func (c *classRepository) GetAll(ctx context.Context) ([]domain.ClassModel, error) {
	var classes []domain.ClassModel
	result := c.db.WithContext(ctx).Model(&domain.ClassModel{}).Find(&classes)
	return classes, handleRepositoryError(result)
}
