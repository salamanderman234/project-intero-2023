package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type MataPelajaran struct {
	gorm.Model
	Nama	string 		`json:"nama" gorm:"type:varchar(255)"`
	Deskripsi	string 	`json:"deskripsi" gorm:"type:varchar(255)"`
	ListKelas	[]Kelas `json:"list_kelas"`
}

func (k *MataPelajaran) GetID() uint {
	return k.ID
}
func (k *MataPelajaran) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *MataPelajaran) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}