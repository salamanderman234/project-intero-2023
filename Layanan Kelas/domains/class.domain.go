package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ClassModel struct {
	gorm.Model
	FocusID uint   `json:"focus_id" gorm:"not null"`
	Group   string `json:"group" gorm:"varchar(255)"`
	GradeID uint   `json:"grade_id" gorm:"not null"`
}

func (ClassModel) TableName() string {
	return "classes"
}

type ClassEntity struct {
	ID      uint           `json:"id,omitempty"`
	FocusID uint           `json:"focus_id" gorm:"not null"`
	Focus   map[string]any `json:"focus"`
	Group   string         `json:"group" gorm:"varchar(255)"`
	GradeID uint           `json:"grade_id" gorm:"not null"`
	Grade   map[string]any `json:"grade"`
}

type ClassCreateForm struct {
	FocusID uint   `json:"focus_id" valid:"int~focus id field must be an integer,required~focus id field is required"`
	Group   string `json:"group" valid:"required~group is required"`
	GradeID uint   `json:"grade_id" valid:"int~grade id field must be an integer,required~grade id field is required"`
}

func (ClassCreateForm) GetModelName() string {
	return "class_subjects"
}

type ClassUpdateForm struct {
	FocusID uint   `json:"focus_id" valid:"int~focus id field must be an integer,required~focus id field is required"`
	Group   string `json:"group" valid:"required~group is required"`
	GradeID uint   `json:"grade_id" valid:"int~grade id field must be an integer,required~grade id field is required"`
}

func (ClassUpdateForm) GetModelName() string {
	return "class_subjects"
}

type ClassRepository interface {
	Create(ctx context.Context, data ClassModel) (ClassModel, error)
	Read(ctx context.Context, query string, id uint, page uint, orderBy string, orderWith string) ([]ClassModel, uint, error)
	Update(ctx context.Context, id uint, data ClassModel) (int, ClassModel, error)
	Delete(ctx context.Context, id uint) (int, error)
	GetAll(ctx context.Context) ([]ClassModel, error)
}

type ClassService interface {
	CreateClass(ctx context.Context, data ClassCreateForm) (ClassEntity, error)
	GetClassList(ctx context.Context, query string, page uint, orderBy string, orderWith string, withoutPagination bool) ([]ClassEntity, Pagination, error)
	GetClassInfo(ctx context.Context, id uint) (ClassEntity, error)
	UpdateClassInfo(ctx context.Context, id uint, updateData ClassUpdateForm) (int, ClassEntity, error)
	DeleteClass(ctx context.Context, id uint) (bool, error)
}

type ClassView interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	// Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
