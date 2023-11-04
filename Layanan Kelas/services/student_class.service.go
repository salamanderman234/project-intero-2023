package service

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
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

func (s *studentClassService) AssignStudent(ctx context.Context, classId uint, studentId uint) (uint, uint, error) {
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	_, err := s.serviceRegistry.ClassServ.GetClassInfo(ctx, classId)
	if err != nil {
		return 0, 0, domain.ErrForeignKeyViolated
	}
	data := domain.StudentClassModel{
		SiswaID: &studentId,
		KelasID: &classId,
	}
	created, err := s.studentClassRepo.Create(ctx, data)
	if err != nil {
		return 0, 0, err
	}	
	return *created.KelasID, *created.SiswaID, nil
}
func (s *studentClassService) UnasssignStudent(ctx context.Context, classId uint, studentId uint) (bool, error) {
	// TODO: panggil api untuk mengecek apakah id benar-benar ada
	_, err := s.serviceRegistry.ClassServ.GetClassInfo(ctx, classId)
	if err != nil {
		return false, domain.ErrForeignKeyViolated
	}
	_, err = s.studentClassRepo.Delete(ctx, classId, studentId)
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