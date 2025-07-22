package mahasiswa

import "fmt"

func hitungRataRata(nilai ...int) float64 {
	sum := 0
	for _, n := range nilai {
		sum += n
	}

	return float64(sum) / float64(len(nilai))
}

func GetMaxNilai() int {
	return maxNilai
}

// Mahasiswa struct builder
func BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa {
	return Mahasiswa{
		Nama:  nama,
		umur:  umur,
		Nilai: nilai,
	}
}

// Implementasi struct Mahasiswa
func (m Mahasiswa) Info() string {
	return m.Nama
}

func (m *Mahasiswa) RataRata() float64 {
	// Mungkin lebih baik Mahasiswa.nilaiAvg baru terinisiasi
	// hanya saat Mahasiswa.RataRata() terpanggil
	m.nilaiAvg = hitungRataRata(m.Nilai...)

	return m.nilaiAvg
}

func (m Mahasiswa) GetUmur() int {
	return m.umur
}

// Implementasi interface Deskripsi
func PrintInfo(d Deskripsi) {
	fmt.Printf("Nama: %s, Umur: %d\n", d.Info(), d.GetUmur())
	fmt.Printf("Rata-rata nilai: %.2f\n", d.RataRata())
	fmt.Println("-----")
}
