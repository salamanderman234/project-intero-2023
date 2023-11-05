package service

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	helper "github.com/salamanderman234/project-intro-2023/layanan-kelas/helpers"
)

type studentClassService struct {
	studentClassRepo domain.StudentClassRepository
	serviceRegistry domain.ServiceRegistry
}

func NewStudentClassService(r domain.StudentClassRepository, sr domain.ServiceRegistry) domain.StudentClassService {
	return &studentClassService{
		studentClassRepo: r,
		serviceRegistry: sr,
	}
}

func (s *studentClassService) AssignStudent(ctx context.Context, assignForm domain.AssignStudentForm) (uint, uint, error) {
	if ok, errs := helper.ValidateForm(assignForm); !ok {
		return 0, 0, errs
	}
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	_, err := s.serviceRegistry.ClassServ.GetClassInfo(ctx, *assignForm.KelasID)
	if err != nil {
		return 0, 0, domain.ErrForeignKeyViolated
	}
	data := domain.StudentClassModel{
		SiswaID: assignForm.SiswaID,
		KelasID: assignForm.KelasID,
	}
	created, err := s.studentClassRepo.Create(ctx, data)
	if err != nil {
		return 0, 0, err
	}	
	return *created.KelasID, *created.SiswaID, nil
}
func (s *studentClassService) UnasssignStudent(ctx context.Context, assignForm domain.AssignStudentForm) (bool, error) {
	if ok, errs := helper.ValidateForm(assignForm); !ok {
		return false, errs
	}
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	_, err := s.serviceRegistry.ClassServ.GetClassInfo(ctx, *assignForm.KelasID)
	if err != nil {
		return false, domain.ErrForeignKeyViolated
	}
	_, err = s.studentClassRepo.Delete(ctx, *assignForm.KelasID, *assignForm.SiswaID)
	if err != nil {
		return false, err
	}	
	return true, nil
}

func (s *studentClassService) GetStudentClassList(ctx context.Context, studentId uint) ([]domain.ClassEntity, error) {
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	var resultsClass []domain.ClassEntity
	results, err := s.studentClassRepo.Read(ctx, 0, studentId)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		class, err := s.serviceRegistry.ClassServ.GetClassInfo(ctx, *result.KelasID)
		if err != nil {
			return nil, err
		}
		resultsClass = append(resultsClass, class)
	}
	return resultsClass, nil
}