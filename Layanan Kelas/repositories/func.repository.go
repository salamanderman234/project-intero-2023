package repository

import (
	"context"
	"math"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

func basicCreateFunc(ctx context.Context, c *gorm.DB, data any, customQuery ...*gorm.DB) (*gorm.DB, error) {
	var result *gorm.DB
	if len(customQuery) == 1 {
		result = customQuery[0]
	} else {
		result = c.WithContext(ctx).Create(data)
	}
	return result, handleRepositoryError(result)
}

func basicSearchFunc(ctx context.Context, c *gorm.DB, query gorm.DB, page uint, orderBy string, orderWith string, model any, results any) (*gorm.DB, int64 ,error) {
	var maxPage int64
	query.Count(&maxPage)
	result := *query.Scopes(orderScope(model, orderBy, orderWith)).Scopes(paginateScope(page)).Find(results)
	if result.RowsAffected <= 0 {
		return &result, 0, domain.ErrResourceNotFound
	}
	maxPage = int64(math.Ceil(float64(maxPage)/DATA_PERPAGE))
	return &result, maxPage, handleRepositoryError(&result)
}

func basicUpdateFunc(ctx context.Context, c *gorm.DB, id uint, data any, customQuery ...*gorm.DB) (*gorm.DB, error){
	var result *gorm.DB
	if len(customQuery) == 1 {
		result = customQuery[0]
	} else {
		result = c.WithContext(ctx).
			Where("id = ?", id).
			Updates(data)
	}
	aff := result.RowsAffected
	err := handleRepositoryError(result)
	if aff != 1 && err == nil{
		return result, domain.ErrResourceNotFound
	}
	return result, handleRepositoryError(result)
}

func basicDeleteFunc(ctx context.Context, c *gorm.DB, id uint, model any, customQuery ...*gorm.DB) (*gorm.DB, error) {
	var result *gorm.DB
	if len(customQuery) == 1 {
		result = customQuery[0]
	} else {
		result = c.WithContext(ctx).
			Where("id = ?", id).
			Delete(model)
	}
	aff := result.RowsAffected
	err := handleRepositoryError(result)
	if aff != 1 && err == nil{
		return result, domain.ErrResourceNotFound
	}
	return result, nil
}