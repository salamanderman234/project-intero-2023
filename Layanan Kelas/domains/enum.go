package domain

type TingkatEnum string

const (
	SATU TingkatEnum = "1"
	DUA  TingkatEnum = "2"
	TIGA TingkatEnum = "3"
)

type KonsentrasiEnum string

const (
	SOSIAL KonsentrasiEnum = "sosial"
	MIPA   KonsentrasiEnum = "mipa"
	BUDAYA KonsentrasiEnum = "budaya"
)

type HariEnum string

const (
	SENIN  HariEnum = "SENIN"
	SELASA HariEnum = "SELASA"
	RABU   HariEnum = "RABU"
	KAMIS  HariEnum = "KAMIS"
	JUMAT  HariEnum = "JUMAT"
	SABTU  HariEnum = "SABTU"
	MINGGU HariEnum = "MINGGU"
)

type JenisKelaminEnum string

const (
	LAKI      JenisKelaminEnum = "L"
	PEREMPUAN JenisKelaminEnum = "P"
)