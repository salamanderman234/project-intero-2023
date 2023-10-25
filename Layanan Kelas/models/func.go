package model

import domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"

func GetAllModel() []domain.Model {
	return []domain.Model{
		&Kelas{},
		&Jadwal{},
		&KelasSiswa{},
		&GrupKelas{},
		&Siswa{},
		&MataPelajaran{},
		&Guru{},
		&TahunAjaran{},
	}
}