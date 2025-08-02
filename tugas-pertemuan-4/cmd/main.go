package main

import (
	"fmt"
	"os"
	"sync"
	"tugaspertemuan4/config"
	"tugaspertemuan4/models"
	"tugaspertemuan4/utils"
	"tugaspertemuan4/worker"
)

var Versi string = "v1.0.0"

func exit(a ...any) {
	fmt.Println(a...)
	os.Exit(1)
}

func main() {
	db := config.GetDB()
	err := models.MigrateAll(db)
	if err != nil {
		exit("Gagal mengeksekusi MigrateAll:", err)
	}

	var count int64
	res := db.Model(&models.MahasiswaModel{}).Count(&count)
	if res.Error != nil {
		exit("Terjadi kesalahan saat menghitung jumlah mahasiswa:", res.Error)
	}
	if count == 0 {
		models.SeedMahasiswa(db)
	}

	var mhs []models.MahasiswaModel
	res = db.Find(&mhs)
	if res.Error != nil {
		exit("Gagal mengambil semua data mahasiswa:", err)
	}

	assignResult := []utils.AssignResult{}

	assignCh := make(chan models.TugasModel)
	resultCh := make(chan utils.AssignResult)

	var wgAssign sync.WaitGroup
	var wgResult sync.WaitGroup

	wgAssign.Add(1)
	wgResult.Add(1)

	// Run all goroutines
	go worker.AssignTugas(mhs, assignCh, &wgAssign)           // Assign task
	go worker.GradeAssignments(assignCh, resultCh, &wgResult) // Grade assigned task

	// Clean ups
	go func() { wgResult.Wait(); close(resultCh) }()
	go func() { wgAssign.Wait(); close(assignCh) }()

	// Append to assignResult over resultCh
	for res := range resultCh {
		assignResult = append(assignResult, res)
	}

	fmt.Println("\nHasil Tugas Mahasiswa:")

	for _, ar := range assignResult {
		fmt.Printf("%s - %s: %d\n", ar.MhsName, ar.AssignmentTitle, ar.AssignmentResult)
	}
}
