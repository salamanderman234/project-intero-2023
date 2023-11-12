package service

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type studentClassService struct {
	studentClassRepo domain.StudentClassRepository
	classService domain.ClassService
}

func NewStudentClassService(r domain.StudentClassRepository, sr domain.ClassService) domain.StudentClassService {
	return &studentClassService{
		studentClassRepo: r,
		classService: sr,
	}
}

func (s *studentClassService) AssignStudent(ctx context.Context, assignForm domain.AssignStudentForm) (uint, uint, error) {
	if ok, errs := helper.ValidateForm(assignForm); !ok {
		return 0, 0, errs
	}
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	results, _ := s.GetStudentClassList(ctx ,assignForm.StudentID,assignForm.ClassID, assignForm.Year)
	if len(results) > 0 {
		return 0,0,domain.ErrDuplicateEnties
	}
	data := domain.StudentClassModel{
		StudentID: &assignForm.StudentID,
		ClassID: &assignForm.ClassID,
		Year: &assignForm.Year,
	}
	created, err := s.studentClassRepo.Create(ctx, data)
	if err != nil {
		return 0, 0, err
	}	
	return *created.ClassID, *created.StudentID, nil
}
func (s *studentClassService) UnasssignStudent(ctx context.Context, assignForm domain.AssignStudentForm) (bool, error) {
	if ok, errs := helper.ValidateForm(assignForm); !ok {
		return false, errs
	}
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	_, err := s.studentClassRepo.Delete(ctx, assignForm.ClassID, assignForm.StudentID, assignForm.Year)
	if err != nil {
		return false, err
	}	
	return true, nil
}

func (s *studentClassService) GetStudentClassList(ctx context.Context, studentId uint, classId uint, year uint) ([]domain.StudentClassEntity, error) {
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	var resultsClass []domain.StudentClassEntity
	results, err := s.studentClassRepo.Read(ctx, classId, studentId, year)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		temp := domain.StudentClassEntity{}
		err := helper.Convert(result, &temp)
		if err != nil {
			return nil, domain.ErrConversionType
		}
		class, err := s.classService.GetClassInfo(ctx, *result.ClassID)
		if err == nil {
			temp.Class = class
			resultsClass = append(resultsClass, temp)
		}
	}
	return resultsClass, nil
}