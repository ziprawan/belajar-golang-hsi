package main

import (
	"fmt"
	"tugaspertemuan3/mahasiswa"
)

var Versi string = "v1.0.0"

// Ada baiknya menggunakan closure function
func hitungUmur(mhs ...mahasiswa.Mahasiswa) func() int {
	total_umur := 0

	for _, m := range mhs {
		total_umur += m.GetUmur()
		mahasiswa.PrintInfo(&m)
	}

	return func() int {
		return total_umur
	}
}

func main() {
	mhs := []mahasiswa.Mahasiswa{}

	// Data dummy mahasiswa
	mhs = append(mhs, mahasiswa.BuatMahasiswa("Ali", 20, 100, 80, 75, 89, 76))
	mhs = append(mhs, mahasiswa.BuatMahasiswa("Budi", 18, 78, 56, 85, 75, 90))
	mhs = append(mhs, mahasiswa.BuatMahasiswa("Cecep", 23, 67, 57, 51, 76, 82))

	// Inisiasi penghitung umur
	getTotalUmur := hitungUmur(mhs...)

	// Berikan output terakhir
	fmt.Printf("Versi Package: %s\n", Versi)
	fmt.Printf("Nilai Maksimum: %d\n", mahasiswa.GetMaxNilai())
	fmt.Printf("Total Umur Mahasiswa: %d\n", getTotalUmur())
}
