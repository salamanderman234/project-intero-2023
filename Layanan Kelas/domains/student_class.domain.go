package domain

import "context"

type StudentClassModel struct {
}

type StudentClassRepository interface {
	Create(ctx context.Context, data StudentClassModel) (StudentClassModel, error)
	Read(ctx context.Context, classId uint, studentId uint) ([]StudentClassModel, error)
	Delete(ctx context.Context, classId uint, studentId uint) (int, error)
}

type StudentClassService interface {
	AssignStudent(ctx context.Context, classId uint, studentId uint) (uint, uint, error)
	UnsignStudent(ctx context.Context, classId uint, studentId uint) (bool, error)
}