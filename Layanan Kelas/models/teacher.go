package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type Guru struct {
	gorm.Model
	Username  *string `json:"username" gorm:"unique;not null;type:varchar(255)"`
	Password  *string `json:"password" gorm:"not null;type:varchar(255)"`
	Nama      *string `json:"nama" gorm:"not null;type:varchar(255)"`
	ListKelas []Kelas `json:"list_kelas"`
}

func (k *Guru) GetID() uint {
	return k.ID
}
func (k *Guru) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *Guru) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}