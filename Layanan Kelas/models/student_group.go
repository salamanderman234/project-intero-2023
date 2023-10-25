package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type GrupKelas struct {
	gorm.Model
	Tingkat 	string 	`json:"tingkat" gorm:"type:varchar(255)"`
	Konsentrasi string 	`json:"konsentrasi" gorm:"type:varchar(255)"`
	Grup 		string 	`json:"grup" gorm:"type:varchar(255)"`
	ListKelas	[]Kelas `json:"list_kelas"`
	ListSiswa	[]Siswa	`json:"list_siswa"`
}

func (k *GrupKelas) GetID() uint {
	return k.ID
}
func (k *GrupKelas) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *GrupKelas) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}