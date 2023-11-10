package service

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type classSubjectService struct {
	repo domain.ClassSubjectRepository
	services domain.ServiceRegistry
}

func NewClassSubjectService(r domain.ClassSubjectRepository, s domain.ServiceRegistry) domain.ClassSubjectService {
	return &classSubjectService{
		repo: r,
		services: s,
	}
}

func (cs *classSubjectService) CreateClassSubject(c context.Context, data domain.ClassSubjectCreateForm) (domain.ClassSubjectEntity, error) {
	var result domain.ClassSubjectEntity
	// validate the form
	if ok, errs := helper.ValidateForm(data); !ok {
		return result, errs
	}
	// TODO: panggil api untuk mengecek setiap id dan jika tidak ada maka return validation error dengan unpross entity
	// convert to model
	var dataModel domain.ClassSubjectModel
	if err := helper.Convert(data, &dataModel); err != nil {
		return result, domain.ErrConversionType
	}
	// call the repo
	resultRepo, err := cs.repo.Create(c, dataModel)
	if err != nil {
		return result, err
	}
	// convert back to entity
	if err := helper.Convert(resultRepo, &result); err != nil {
		return result, domain.ErrConversionType
	}
	return result, nil
}
func (cs *classSubjectService) GetStudentClassSubject(c context.Context, studentID uint, year uint, page uint, orderBy string, orderWith string) ([]domain.ClassSubjectEntity, domain.Pagination, error) {
	return cs.basicClassSubjectSearch(c, 0, 0, studentID, 0, year, page, orderBy, orderWith)
}
func (cs *classSubjectService) GetClassSubject(c context.Context, classID uint, year uint, page uint, orderBy string, orderWith string) ([]domain.ClassSubjectEntity, domain.Pagination, error) {
	return cs.basicClassSubjectSearch(c, 0, classID, 0, 0, year, page, orderBy, orderWith)
}
func (cs *classSubjectService) GetTeacherClassSubject(c context.Context, teacherID uint, year uint, page uint, orderBy string, orderWith string) ([]domain.ClassSubjectEntity, domain.Pagination, error) {
	return cs.basicClassSubjectSearch(c, 0, 0, 0, teacherID, year, page, orderBy, orderWith)
}
func (cs *classSubjectService) FindClassSubject(c context.Context, id uint) (domain.ClassSubjectEntity, error) {
	results, _, err := cs.basicClassSubjectSearch(c, id, 0, 0, 0, 0, 0, "", "")
	if len(results) == 1 {
		return results[0], nil
	}
	return domain.ClassSubjectEntity{}, err
}
func (cs *classSubjectService) UpdateClassSubject(c context.Context, id uint, data domain.ClassSubjectUpdateForm) (domain.ClassSubjectEntity, error) {
	var updatedFields domain.ClassSubjectEntity
	if ok, errs := helper.ValidateForm(data); !ok {
		return updatedFields, errs
	}
	var updatedFieldsModel domain.ClassSubjectModel
	if err := helper.Convert(data, &updatedFieldsModel); err != nil {
		return updatedFields, domain.ErrConversionType
	}
	_, updatedFieldsModel, err := cs.repo.Update(c, id, updatedFieldsModel)
	if err != nil {
		return updatedFields, err
	}
	if err := helper.Convert(updatedFieldsModel, &updatedFields); err != nil {
		return updatedFields, domain.ErrConversionType
	}
	return  updatedFields, nil
}
func (cs *classSubjectService) DeleteClassSubject(c context.Context, id uint) (error) {
	_, err := cs.repo.Delete(c, id)
	return err
}

func (cs *classSubjectService) basicClassSubjectSearch(c context.Context, id uint, classId uint, studentId uint, teacherId uint, year uint, page uint, orderBy string, orderWith string) ([]domain.ClassSubjectEntity, domain.Pagination, error){
	var results []domain.ClassSubjectEntity
	var pagination domain.Pagination
	// TODO: panggil api untuk mengecek setiap id dan jika tidak ada maka return validation error dengan unpross entity
	// call the repo
	resultRepos, maxPage, err := cs.repo.Read(c, id, studentId, teacherId, classId, year, orderBy, orderWith)
	if err != nil {
		return results, pagination, err
	}
	// TODO: panggil api lain untuk mendapatkan detail (preload)
	for _, resultRepo := range resultRepos {
		var temp domain.ClassSubjectEntity
		if err := helper.Convert(resultRepo, &temp); err != nil {
			return results, pagination, domain.ErrConversionType
		}
		results = append(results, temp)
	}
	pagination = domain.CreatePagination(page, maxPage)
	return results, pagination, nil
}