# Tugas Pertemuan 3

## Sistem Informasi Mahasiswa Modular

### Tujuan

Melatih pemahaman mengenai:

- Penggunaan function dari variadic, closure, dan chaining
- Penerapan pointer dan pengaruhnya terhadap data
- Struct dan inetrface (beserta polymorphism)
- Modular programming menggunakan package
- Scopre dan pengaturan akses (private/public)

### Struktur

#### mahasiswa/model.go

Berkas ini berisi struct `Mahasiswa` yang bisa menampung data data dari satu mahasiswa dan interface `Deskripsi` yang menampung fungsi fungsi pembantu untuk mengimplementasikan konsep _Polymorphism_.

Berkas ini juga terdapar variabel privat `maxNilai` yang bernilai 100.

#### mahasiswa/utils.go

Berisi fungsi privat:

1. `hitungRataRata(...) float64` - Menerima parameter variadic int dan mengembalikan rata rata dari semua parameter yang dimasukkan dengan tipe float64.

Berisi fungsi publik:

1. `GetMaxNilai() int` - Mengembalikan variabel maxNilai yang sudah disebutkan pada berkas mahasiswa/model.go.
2. `PrintInfo(d Deskripsi)` - Menerima parameter d dengan tipe Deskripsi dan prosedur ini akan mengeluarkan output info mahasiswa dengan memanfaatkan semua fungsi yang ada di interface `Deskripsi`.

Implementasi struct `Mahasiswa`:

1. `(m Mahasiswa) Info() string` - Menerima parameter `m` dengan tipe `Mahasiswa` dan mengembalikan nilai dari `m.Nama`.
2. `(m *Mahasiswa) RataRata() float64` - Menerima parameter `m` dengan tipe pointer to `Mahasiswa` (`*Mahasiswa`), lalu menghitung rata rata nilai berdasarkan dari m.Nilai menggunakan bantuan fungsi `hitungRataRata(...) float64` dan menyimpannya kembali ke variabel `m.nilaiAvg`. Mengembalikan nilai dari `m.nilaiAvg`.
3. `(m Mahasiswa) GetUmur() int` - Menerima parameter `m` dengan tipe `Mahasiswa` dan mengembalikan nilai privat dari `m.umur`.

#### main.go

Berkas ini adalah akses utama untuk program ini yang berisikan 2 fungsi.

- `hitungUmur(...) func () int` -
  Fungsi ini berupa variadic function yang menerima parameter bertipe mahasiswa.Mahasiswa. Fungsi ini akan mengeluarkan info semua input parameter mahasiswa dengan memanggil fungsi `mahasiswa.PrintInfo(d Deskripsi)` dan menambahkan umur mahasiswa ke `total_umur`. Mengembalikan fungsi yang mengembalikan nilai `total_umur`.
- `main()` - Fungsi utama agar program bisa dieksekusi. Berisi data dummy mahasiswa (_hard-coded_) lalu menginisiasi fungsi `hitungUmur(...) func () int` dan diakhiri dengan mengeluarkan output info versi package, nilai maksimum, dan total keseluruhan umur mahasiswa dengan menggunakan fungsi-funsgi yang sudah disebutkan sebelumnya.
