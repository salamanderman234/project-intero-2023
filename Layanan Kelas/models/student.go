package model

import (
	"context"
	"time"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type Siswa struct {
	gorm.Model
	Username 		*string 				`json:"username" gorm:"unique;not null;type:varchar(255)"`
	Password 		*string 				`json:"password" gorm:"not null;type:varchar(255)"`
	Nama 			*string 				`json:"nama" gorm:"not null;type:varchar(255)"`
	NIS 			*string 				`json:"nis" gorm:"unique;not null;type:varchar(255)"`
	TahunMasuk 		int 					`json:"tahun_masuk"`
	TanggalLahir 	time.Time 				`json:"tanggal_lahir" `
	TempatLahir 	string 					`json:"tempat_lahir" gorm:"default:'';type:varchar(255)"`
	Alamat 			string 					`json:"alamat" gorm:"default:'';type:varchar(255)"`
	JenisKelamin 	string					`json:"jenis_kelamin" gorm:"varchar(255)"`
	NoHandphone 	string 					`json:"ho_handphone" gorm:"default:'';type:varchar(255)"`
	GrupKelasID		*uint 					`json:"grup_kelas_id" gorm:"not null"`
	GrupKelas 		GrupKelas				`json:"grup_kelas" gorm:"foreignKey:GrupKelasID;references:ID"`
	ListKelas		[]Kelas		  			`json:"list_kelas" gorm:"many2many:kelas_siswas"`
}

func (k *Siswa) GetID() uint {
	return k.ID
}
func (k *Siswa) GetPreloadStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}
func (k *Siswa) GetAdditionalStatement() domain.QueryFunc {
	return func(ctx context.Context, query *gorm.DB) *gorm.DB {
		return nil
	}
}