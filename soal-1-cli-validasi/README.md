# CLI Sederhana: Validasi Umur

Program CLI interaktif sederhana yang meminta input (masukan) dari pengguna untuk nama dan umur, lalu memvalidasi nilai umur harus lebih dari atau sama dengan 18. Jika tidak valid akan menampilkan error dengan pesan yang jelas dan log kesalahan, jika valid akan menampilkan ucapan selamat datang yang disertai dengan nama yang dimasukkan tadi.

## Menjalankan program

Ada dua cara menjalankan program. Pertama menjalankan langsung menggunakan "go run"

`go run main.go`

Kedua, mem-build source code menjadi file yang bisa dieksekusi (misalnya \*.exe)

```bash
go build -o build/main main.go  # Build
./build/main                    # Menjalankan
```
