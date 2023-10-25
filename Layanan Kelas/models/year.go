package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type TahunAjaran struct {
	gorm.Model
	Semester string `json:"semester" gorm:"type:varchar(255)"`
	TahunMulai uint `json:"tahun_mulai"`
	TahunBerakhir uint `json:"tahun_berakhir"`
	ListJadwal []Jadwal `json:"list_jadwal"`
	ListKelas  []Kelas 	`json:"list_kelas"`
}

func (k *TahunAjaran) GetID() uint {
	return k.ID
}
func (k *TahunAjaran) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *TahunAjaran) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}