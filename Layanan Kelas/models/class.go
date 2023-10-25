package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type Kelas struct {
	gorm.Model
	Deskripsi       string        `json:"deskripsi" gorm:"type:varchar(255);default:'-'"`
	GrupKelasID     *uint         `json:"grup_kelas_id" gorm:"not null"`
	GrupKelas       GrupKelas     `json:"grup_kelas" gorm:"foreignKey:GrupKelasID;references:ID"`
	GuruID          *uint         `json:"guru_id" gorm:"not null"`
	Guru            Guru          `json:"guru" gorm:"foreignKey:GuruID;references:ID"`
	MataPelajaranID *uint         `json:"mata_pelajaran_id" gorm:"not null"`
	MataPelajaran   MataPelajaran `json:"mata_pelajaran" gorm:"foreignKey:MataPelajaranID;references:ID"`
	TahunAjaranID   *uint         `json:"tahun_ajaran_id" gorm:"not null"`
	TahunAjaran     TahunAjaran   `json:"tahun_ajaran" gorm:"foreignKey:TahunAjaranID;references:ID"`
	ListSiswa       []Siswa       `json:"list_siswa" gorm:"many2many:kelas_siswas"`
}

func (k *Kelas) GetID() uint {
	return k.ID
}
func (k *Kelas) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *Kelas) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}