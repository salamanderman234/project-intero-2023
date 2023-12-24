package repository

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type studentClassRepository struct {
	conn *gorm.DB
}

func NewStudentClassRepository(c *gorm.DB) domain.StudentClassRepository {
	return &studentClassRepository{
		conn: c,
	}
}

func (s *studentClassRepository) Create(ctx context.Context, data domain.StudentClassModel) (domain.StudentClassModel, error) {
	_, err := basicCreateFunc(ctx, s.conn, &data)
	return data, err
}
func (s *studentClassRepository) Read(ctx context.Context, classId uint, studentId uint, year uint) ([]domain.StudentClassModel, error) {
	var results []domain.StudentClassModel
	var err error
	searchQuery := s.conn.Model(&domain.StudentClassModel{}).WithContext(ctx)
	if classId != 0 {
		searchQuery = searchQuery.Where("class_id = ?", classId)
	}
	if studentId != 0{
		searchQuery =searchQuery.Where("student_id = ?", studentId)
	}
	if year != 0 {
		searchQuery =searchQuery.Where("year = ?", year)
	}
	_, _, err = basicSearchFunc(ctx, s.conn, *searchQuery, 1, "", "", domain.StudentClassModel{}, &results)
	return results, err
}
func (s *studentClassRepository) Delete(ctx context.Context, classId uint, studentId uint, year uint) (int, error) {
	query := s.conn.WithContext(ctx).
		Where("class_id = ?", classId).
		Where("student_id = ?", studentId).
		Where("year = ?", year).
		Delete(&domain.StudentClassModel{})
	result, err := basicDeleteFunc(ctx, s.conn, 0, domain.StudentClassModel{}, query)
	return int(result.RowsAffected), err
}