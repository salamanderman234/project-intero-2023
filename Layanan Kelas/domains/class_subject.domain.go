package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ClassSubjectModel struct {
	gorm.Model
	SubjectID 	*uint `json:"subject_id" gorm:"not null" valid:"int~subject id field must be an integer,required~subject id field is required"`
	TeacherID 	*uint `json:"teacher_id" gorm:"not null" valid:"int~teacher id field must be an integer,required~teacher id field is required"`
	ClassID 	*uint `json:"class_id" gorm:"not null" valid:"int~class id field must be an integer,required~class id field is required"`
	Year 		uint  `json:"year" valid:"int~year field must be an integer,required~year field is required"`
}

func(ClassSubjectModel) TableName() string {
	return "class_subjects"
}

type ClassSubjectEntity struct {
	ID 			uint  			`json:"id,omitempty"`
	SubjectID 	*uint 			`json:"subject_id,omitempty" gorm:"not null" valid:"int~subject id field must be an integer,required~subject id field is required"`
	Subject		any 			`json:"subject,omitempty"`
	TeacherID 	*uint 			`json:"teacher_id,omitempty" gorm:"not null" valid:"int~teacher id field must be an integer,required~teacher id field is required"`
	Teacher 	any 			`json:"teacher,omitempty"`
	Class		any 			`json:"class,omitempty"`
	Year 		uint  			`json:"year" valid:"int~year field must be an integer,required~year field is required"`
}

type ClassSubjectCreateForm struct {
	SubjectID 	uint `json:"subject_id" gorm:"not null" valid:"int~subject id field must be an integer,required~subject id field is required"`
	TeacherID 	uint `json:"teacher_id" gorm:"not null" valid:"int~teacher id field must be an integer,required~teacher id field is required"`
	ClassID 	uint `json:"class_id" gorm:"not null" valid:"int~class id field must be an integer,required~class id field is required"`
	Year 		uint  `json:"year" valid:"int~year field must be an integer,required~year field is required"`
}

func(ClassSubjectCreateForm) GetModelName() string {
	return "class_subjects"
}

type ClassSubjectUpdateForm struct {
	SubjectID 	uint `json:"subject_id" gorm:"not null" valid:"int~subject id field must be an integer,required~subject id field is required"`
	TeacherID 	uint `json:"teacher_id" gorm:"not null" valid:"int~teacher id field must be an integer,required~teacher id field is required"`
	Year 		uint  `json:"year" valid:"int~year field must be an integer,required~year field is required"`
}

func(ClassSubjectUpdateForm) GetModelName() string {
	return "class_subjects"
}

type ClassSubjectRepository interface {
	Create(c context.Context, data ClassSubjectModel) (ClassSubjectModel, error)
	Read(c context.Context, id uint, studentID uint, teacherID uint, classID uint, year uint, page uint, orderBy string, orderWith string) ([]ClassSubjectModel, uint, error)
	Update(c context.Context, id uint, data ClassSubjectModel) (int64, ClassSubjectModel, error)
	Delete(c context.Context, id uint) (int64, error)
}

type ClassSubjectService interface {
	CreateClassSubject(c context.Context, data ClassSubjectCreateForm) (ClassSubjectEntity, error)
	// GetStudentClassSubject(c context.Context, studentID uint, year uint, page uint, orderBy string, orderWith string) ([]ClassSubjectEntity, Pagination, error)
	// GetClassSubject(c context.Context, classID uint, year uint, page uint, orderBy string, orderWith string) ([]ClassSubjectEntity, Pagination, error)
	// GetTeacherClassSubject(c context.Context, teacherID uint, year uint, page uint, orderBy string, orderWith string) ([]ClassSubjectEntity, Pagination, error)
	// FindClassSubject(c context.Context, id uint) (ClassSubjectEntity, error)
	GetClassSubject(c context.Context, id uint, classId uint, studentId uint, teacherId uint, year uint, page uint, orderBy string, orderWith string) ([]ClassSubjectEntity, Pagination, error)
	UpdateClassSubject(c context.Context, id uint, data ClassSubjectUpdateForm) (int64, ClassSubjectEntity, error)
	DeleteClassSubject(c context.Context, id uint) (error)
}

type ClassSubjectView interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}