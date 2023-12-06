package service

import (
	"context"
	"math"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type classService struct {
	classRepo domain.ClassRepository
	serviceList domain.ServiceRegistry
}

func NewClassService(r domain.ClassRepository, s domain.ServiceRegistry) domain.ClassService {
	return &classService{
		classRepo: r,
		serviceList: s,
	}
}

func (c *classService) CreateClass(ctx context.Context, data domain.ClassCreateForm) (domain.ClassEntity, error) {
	var result domain.ClassEntity
	// validate the form
	if ok, errs := helper.ValidateForm(data); !ok {
		return result, errs
	}
	// TODO: panggil api untuk mengecek setiap id dan jika tidak ada maka return validation error dengan unpross entity
	// convert to model
	var dataModel domain.ClassModel
	if err := helper.Convert(data, &dataModel); err != nil {
		return result, domain.ErrConversionType
	}
	// call the repo
	resultRepo, err := c.classRepo.Create(ctx, dataModel)
	if err != nil {
		return result, err
	}
	// convert back to entity
	if err := helper.Convert(resultRepo, &result); err != nil {
		return result, domain.ErrConversionType
	}
	return result, nil
}
func (c *classService) GetClassList(ctx context.Context, query string, page uint,orderBy string, orderWith string) ([]domain.ClassEntity, domain.Pagination, error) {
	var results []domain.ClassEntity
	pagination := domain.Pagination{}
	resultsRepos, maxPage, err := c.classRepo.Read(ctx, query, 0, uint(math.Max(float64(page), 1)), orderBy, orderWith)
	if err != nil {
		return results, pagination, err
	}
	// TODO: panggil api lain untuk mendapatkan detail (preload)
	for _, resultRepo := range resultsRepos {
		var temp domain.ClassEntity
		if err := helper.Convert(resultRepo, &temp); err != nil {
			return results, pagination, domain.ErrConversionType
		}
		results = append(results, temp)
	}
	paginationQueries := map[string]any{}
	if query != "" {
		paginationQueries["q"] =  query
	} 
	if orderBy != "" {
		paginationQueries["order_by"] = orderBy
	}
	if orderWith != "" {
		paginationQueries["order"] = orderWith
	}
	pagination = domain.CreatePagination(page, maxPage, paginationQueries)
	return results, pagination, nil
}
func (c *classService) GetClassInfo(ctx context.Context, id uint) (domain.ClassEntity, error) {
	var result domain.ClassEntity
	resultRepo, _ ,err := c.classRepo.Read(ctx, "", id, 1, "", "")
	if err != nil {
		return result, err
	}
	// TODO: panggil api lain untuk mendapatkan detail (preload)
	if err := helper.Convert(resultRepo[0], &result); err != nil {
		return result, domain.ErrConversionType
	}
	return result, nil
}
func (c *classService) UpdateClassInfo(ctx context.Context, id uint, updateData domain.ClassUpdateForm) (int, domain.ClassEntity, error) {
	var updatedFields domain.ClassEntity
	if ok, errs := helper.ValidateForm(updateData); !ok {
		return 0, updatedFields, errs
	}
	var updatedFieldsModel domain.ClassModel
	if err := helper.Convert(updateData, &updatedFieldsModel); err != nil {
		return 0, updatedFields, domain.ErrConversionType
	}
	aff, updatedFieldsModel, err := c.classRepo.Update(ctx, id, updatedFieldsModel)
	if err != nil {
		return 0, updatedFields, err
	}
	if err := helper.Convert(updatedFieldsModel, &updatedFields); err != nil {
		return 0, updatedFields, domain.ErrConversionType
	}
	return aff, updatedFields, nil
}
func (c *classService) DeleteClass(ctx context.Context, id uint) (bool, error) {
	_, err := c.classRepo.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
