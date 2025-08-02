package mahasiswa

type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int
	nilaiAvg float64
}

type Deskripsi interface {
	Info() string
	RataRata() float64
	GetUmur() int
}

var maxNilai int = 100
