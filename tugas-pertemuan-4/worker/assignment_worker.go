package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"tugaspertemuan4/config"
	"tugaspertemuan4/models"

	"gorm.io/gorm"
)

var defaultAssignments []string = []string{
	"Tugas Pemrograman Goroutine",
	"Tugas Implementasi WaitGroup",
	"Tugas Implementasi Mutex",
	"Tugas Implementasi Channel",
}

func shuffleAssign() []string {
	tugas := make([]string, len(defaultAssignments))
	copy(tugas, defaultAssignments)

	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src).Shuffle(len(tugas), func(i, j int) {
		tugas[i], tugas[j] = tugas[j], tugas[i]
	})

	return tugas
}

func assignTugasSingle(db *gorm.DB, mhs models.MahasiswaModel) *models.TugasModel {
	allAssigns := shuffleAssign()

	for i := range len(allAssigns) {
		assignTitle := allAssigns[i]
		assignDb := models.TugasModel{
			MahasiswaID: mhs.ID,
			Judul:       assignTitle,
		}

		query := db.Where(&assignDb).Find(&assignDb)
		if err := query.Error; err != nil {
			panic(fmt.Errorf("Gagal mendapatkan tugas \"%s\" untuk mahasiswa ID %d\n", assignTitle, mhs.ID))
		}

		if query.RowsAffected == 1 {
			// Already assigned, assign another task
			continue
		}

		assignDb.Deskripsi = assignTitle
		db.Create(&assignDb)

		return &assignDb
	}

	return nil
}

func AssignTugas(mhs []models.MahasiswaModel, assignCh chan models.TugasModel, wgAssign *sync.WaitGroup) {
	defer wgAssign.Done()

	db := config.GetDB()

	for _, mahasiswa := range mhs {
		assignModel := assignTugasSingle(db, mahasiswa)
		if assignModel == nil {
			// It returns nil if and only if all tasks already assigned
			fmt.Printf("Tidak ada tugas yang bisa diberikan ke \"%s\"\n", mahasiswa.Nama)
			continue
		}

		fmt.Printf("Tugas \"%s\" diberikan ke \"%s\"\n", assignModel.Judul, mahasiswa.Nama)

		assignCh <- *assignModel
	}
}
