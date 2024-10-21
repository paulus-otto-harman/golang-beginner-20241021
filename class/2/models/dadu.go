package models

type Dadu struct {
	angka int
}

func InitDadu(angka int) *Dadu {
	return &Dadu{angka}
}

func Cetak(dadu Dadu) int {
	return dadu.angka
}
