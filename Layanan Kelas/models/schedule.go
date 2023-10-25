package model

import (
	"context"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)


type Jadwal struct {
	gorm.Model
	Hari			string				`json:"hari" gorm:"type:varchar(255)"`
	JamMulai 		string				`json:"jam_mulai" gorm:"not null;type:varchar(255)"`
	JamAkhir 		string 				`json:"jam_akhir" gorm:"not null;type:varchar(255)"`
	KelasID 		*uint 				`json:"kelas_id" gorm:"not null"`
	Kelas 			Kelas 				`json:"kelas" gorm:"foreignKey:KelasID;references:ID"`
	TahunAjaranID 	*uint 				`json:"tahun_ajaran_id" gorm:"not null"`
	TahunAjaran 	TahunAjaran 		`json:"tahun_ajaran" gorm:"foreignKey:TahunAjaranID;references:ID"`
}

func (k *Jadwal) GetID() uint {
	return k.ID
}
func (k *Jadwal) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *Jadwal) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}