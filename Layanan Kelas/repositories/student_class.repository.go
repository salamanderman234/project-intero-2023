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
func (s *studentClassRepository) Read(ctx context.Context, classId uint, studentId uint) ([]domain.StudentClassModel, error) {
	var results []domain.StudentClassModel
	var err error
	queryDB := s.conn.Model(&domain.ClassModel{}).WithContext(ctx)
	searchQuery := queryDB
	if studentId != 0 {
		searchQuery = searchQuery.Where("kelas_id = ?", classId)
	}
	if classId != 0 {
		searchQuery = searchQuery.Where("siswa_id = ?", studentId)
	}
	_, _, err = basicSearchFunc(ctx, s.conn, *searchQuery, 1, "", "", domain.StudentClassModel{}, &results)
	return results, err
}
func (s *studentClassRepository) Delete(ctx context.Context, classId uint, studentId uint) (int, error) {
	query := s.conn.WithContext(ctx).
		Where("kelas_id = ?", classId).
		Where("siswa_id = ?", studentId).
		Delete(&domain.StudentClassModel{})
	result, err := basicDeleteFunc(ctx, s.conn, 0, nil, query)
	return int(result.RowsAffected), err
}