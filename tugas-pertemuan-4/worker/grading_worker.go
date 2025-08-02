package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"tugaspertemuan4/config"
	"tugaspertemuan4/models"
	"tugaspertemuan4/utils"

	"gorm.io/gorm"
)

var cache map[uint]string = map[uint]string{}

func getMhsName(db *gorm.DB, id uint) string {
	if name, ok := cache[id]; ok {
		return name
	}

	// Not in cache, fetch them all
	var mhs []models.MahasiswaModel
	db.Find(&mhs)

	// Save name to cache
	for _, m := range mhs {
		cache[m.ID] = m.Nama
	}

	return cache[id]
}

func GradeAssignments(assignCh chan models.TugasModel, resultCh chan utils.AssignResult, wgResult *sync.WaitGroup) {
	defer wgResult.Done()

	db := config.GetDB()

	// Loop over channel
	for assign := range assignCh {
		src := rand.NewSource(time.Now().UnixNano())
		value := rand.New(src).Intn(101)

		result := models.HasilModel{
			TugasID: assign.ID,
			Nilai:   uint(value),
		}

		db.Create(&result)

		name := getMhsName(db, assign.MahasiswaID)
		fmt.Printf("Nilai %d diberikan ke \"%s\" untuk tugas \"%s\"\n", value, assign.Judul, name)

		resultCh <- utils.AssignResult{
			MhsName:          name,
			AssignmentTitle:  assign.Judul,
			AssignmentResult: result.Nilai,
		}
	}

	// close(resultCh)
	fmt.Println("Niggers 1")
}
