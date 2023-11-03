package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ClassModel struct {
	gorm.Model
	GrupKelasID     *uint         		`json:"grup_kelas_id" gorm:"not null" valid:"int,required~field grup kelas id diperlukan"`
	GrupKelas       map[string]string   `json:"grup_kelas" gorm:"foreignKey:GrupKelasID;references:ID"`
	GuruID          *uint         		`json:"guru_id" gorm:"not null" valid:"int,required~field guru id diperlukan"`
	Guru            map[string]string   `json:"guru,omitempty" gorm:"foreignKey:GuruID;references:ID"`
	MataPelajaranID *uint         		`json:"mata_pelajaran_id" gorm:"not null" valid:"int,required~mata pelajaran id diperlukan"`
	MataPelajaran   map[string]string   `json:"mata_pelajaran,omitempty" gorm:"foreignKey:MataPelajaranID;references:ID"`
	TahunAjaranID   *uint         		`json:"tahun_ajaran_id" gorm:"not null" valid:"int,required~field mata pelajaran id diperlukan"`
	TahunAjaran     map[string]string   `json:"tahun_ajaran,omitempty" gorm:"foreignKey:TahunAjaranID;references:ID"`
	ListSiswa       []map[string]string `json:"list_siswa,omitempty"`
	Deskripsi       string        		`json:"deskripsi" gorm:"type:varchar(255);default:'-'"`
}

type ClassEntity struct {
	ID 				uint 				`json:"id"`
	GrupKelasID     *uint         		`json:"grup_kelas_id,omitempty"`
	GrupKelas       map[string]string   `json:"grup_kelas"`
	GuruID          *uint         		`json:"guru_id,omitempty"`
	Guru            map[string]string   `json:"guru"`
	MataPelajaranID *uint         		`json:"mata_pelajaran_id,omitempty"`
	MataPelajaran   map[string]string   `json:"mata_pelajaran"`
	TahunAjaranID   *uint         		`json:"tahun_ajaran_id,omitempty"`
	TahunAjaran     map[string]string   `json:"tahun_ajaran"`
	ListSiswa       []map[string]string `json:"list_siswa"`
	Deskripsi       string        		`json:"deskripsi"`
}

type ClassCreateForm struct {
	GrupKelasID     *uint         		`json:"grup_kelas_id" valid:"required~grup kelas diperlukan"`
	GuruID          *uint         		`json:"guru_id" valid:"required~guru diperlukan"`
	MataPelajaranID *uint         		`json:"mata_pelajaran_id" valid:"required~mata pelajaran diperlukan"`
	TahunAjaranID   *uint         		`json:"tahun_ajaran_id" valid:"tahun ajaran diperlukan"`
	Deskripsi       string        		`json:"deskripsi"`
}

type ClassUpdateForm struct {
	GuruID          *uint         		`json:"guru_id" valid:"required~guru diperlukan"`
	Deskripsi       string        		`json:"deskripsi"`
}

type ClassRepository interface {
	Create(ctx context.Context, data ClassModel) (ClassModel, error)
	Read(ctx context.Context, query string, id uint,page uint) ([]ClassModel, uint, error)
	Update(ctx context.Context, id uint, data ClassModel) (int, ClassModel, error)
	Delete(ctx context.Context, id uint) (int, error)
}

type ClassService interface {
	CreateClass(ctx context.Context, data ClassCreateForm) (ClassEntity, error)
	GetClassList(ctx context.Context, query string, page uint) ([]ClassEntity, Pagination,error)
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