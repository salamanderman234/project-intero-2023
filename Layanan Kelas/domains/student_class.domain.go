package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentClassModel struct {
	gorm.Model
	KelasID *uint `json:"kelas_id" gorm:"not null" valid:"int,required~field kelas id diperlukan"`
	SiswaID *uint `json:"siswa_id" gorm:"not null" valid:"int,required~field siswa id diperlukan"`
}

type AssignStudentForm struct {
	KelasID *uint `json:"kelas_id" valid:"int,required~field kelas id diperlukan"`
	SiswaID *uint `json:"siswa_id" valid:"int,required~field siswa id diperlukan"`
}

func (StudentClassModel) TableName() string {
	return "kelas_siswa"
}

type StudentClassRepository interface {
	Create(ctx context.Context, data StudentClassModel) (StudentClassModel, error)
	Read(ctx context.Context, classId uint, studentId uint) ([]StudentClassModel, error)
	Delete(ctx context.Context, classId uint, studentId uint) (int, error)
}

type StudentClassService interface {
	AssignStudent(ctx context.Context, classId uint, studentId uint) (uint, uint, error)
	UnasssignStudent(ctx context.Context, classId uint, studentId uint) (bool, error)
	GetStudentClassList(ctx context.Context, studentId uint) ([]ClassEntity, error)
}

type StudentClassView interface {
	AssignStudentToAClass(c echo.Context) error
	UnsignStudentFromAClass(c echo.Context) error
	GetStudentClassList(c echo.Context) error
}