package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type KelasSiswa struct {
	gorm.Model
	MateriDibaca	uint 	`json:"materi_dibaca" gorm:"default:0"`
	KelasID   		*uint 	`json:"kelas_id" gorm:"not null"`
	Kelas     		Kelas   `json:"kelas" gorm:"foreignKey:KelasID;references:ID"`
	SiswaID   		*uint   `json:"siswa_id" gorm:"not null"`
	Siswa     		Siswa   `json:"siswa" gorm:"foreignKey:SiswaID;references:ID"`
}

func (k *KelasSiswa) GetID() uint {
	return k.ID
}
func (k *KelasSiswa) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *KelasSiswa) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}