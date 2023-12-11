package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentClassModel struct {
	gorm.Model
	ClassID   *uint `json:"class_id" gorm:"not null" valid:"int~class id field must be an integer,required~class id field is required"`
	StudentID *uint `json:"student_id" gorm:"not null" valid:"int~student id field must be an integer,required~student id field is required"`
	Year      *uint `json:"year" gorm:"int~year field must be an integer,required~year field is required"`
}
type StudentClassEntity struct {
	ID uint `json:"id"`
	// ClassID 	*uint `json:"class_id" gorm:"not null" valid:"int~class id field must be an integer,required~class id field is required"`
	Class     any            `json:"class"`
	StudentID *uint          `json:"student_id" gorm:"not null" valid:"int~student id field must be an integer,required~student id field is required"`
	Student   map[string]any `json:"student"`
	Year      *uint          `json:"year" gorm:"int~year field must be an integer,required~year field is required"`
}

type AssignStudentForm struct {
	ClassID   uint `json:"class_id" gorm:"not null" valid:"int,required~class id field is required"`
	StudentID uint `json:"student_id" gorm:"not null" valid:"int,required~student id field is required"`
	Year      uint `json:"year" valid:"int~year field must be an integer,required~year field is required"`
}

func (AssignStudentForm) GetModelName() string {
	return "student_classes"
}

func (StudentClassModel) TableName() string {
	return "student_classes"
}

type StudentClassRepository interface {
	Create(ctx context.Context, data StudentClassModel) (StudentClassModel, error)
	Read(ctx context.Context, classId uint, studentId uint, year uint) ([]StudentClassModel, error)
	Delete(ctx context.Context, classId uint, studentId uint, year uint) (int, error)
}

type StudentClassService interface {
	AssignStudent(ctx context.Context, assignForm AssignStudentForm) (uint, uint, error)
	UnasssignStudent(ctx context.Context, assignForm AssignStudentForm) (bool, error)
	GetStudentClassList(ctx context.Context, studentId uint, classId uint, year uint) ([]StudentClassEntity, error)
}

type StudentClassView interface {
	AssignStudentToAClass(c echo.Context) error
	UnsignStudentFromAClass(c echo.Context) error
	GetStudentClassList(c echo.Context) error
}
