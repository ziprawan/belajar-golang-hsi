package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"pertemuan6/config"
	"pertemuan6/models"
	"pertemuan6/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllStudentHandler(c *fiber.Ctx) error {
	db := config.GetDB()

	// Defaults
	var limit uint = 50
	var page uint = 1

	q := c.Queries()
	qLimit, ok := q["page_size"]
	if ok {
		limit = utils.SafeParseUint(qLimit)
	}

	qPage, ok := q["page"]
	if ok {
		page = utils.SafeParseUint(qPage)
	}

	if limit < 1 || limit > 100 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "page_size query must be a number between 1-100"})
	}
	if page < 1 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "page query must be a number that greater than 0"})
	}

	offset := limit * (page - 1)

	var students []models.StudentModel
	if tx := db.Limit(int(limit)).Offset(int(offset)).Find(&students); tx.Error != nil {
		fmt.Println(tx.Error)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	mappedStudents := []models.Student{}
	for _, student := range students {
		mappedStudents = append(mappedStudents, student.Student)
	}

	var count int64
	db.Model(&models.StudentModel{}).Count(&count)

	return c.JSON(models.GetAllStudentResponse{
		Students: mappedStudents,
		Pagination: models.Pagination{
			CurrentPage: page,
			PageSize:    limit,
			TotalItems:  uint(count),
			TotalPages:  uint(math.Ceil(float64(count) / float64(limit))),
		},
	})
}

func GetStudentByIdHandler(c *fiber.Ctx) error {
	db := config.GetDB()

	idParam := c.Params("id", "0")
	studentId, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Given student ID is not a positive number"})
	}

	if studentId == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid student ID 0"})
	}

	student := models.StudentModel{Student: models.Student{ID: uint(studentId)}}
	tx := db.Find(&student)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	if tx.RowsAffected == 0 {
		return c.Status(404).JSON(models.NotFoundErrorResponse)
	}

	return c.JSON(models.GetStudentByIdResponse{Student: student.Student})
}

func CreateStudentHandler(c *fiber.Ctx) error {
	db := config.GetDB()
	body := c.Body()

	req := models.CreateStudentRequest{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
	}

	var student models.StudentModel
	if len(req.NIM) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "nim field is empty"})
	}
	if len(req.Name) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "name field is empty"})
	}
	if len(req.Email) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "email field is empty"})
	}
	if len(req.Major) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "major field is empty"})
	}
	if req.Semester < 1 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "semester field is less than 1"})
	}

	student.NIM = req.NIM
	student.Name = req.Name
	student.Email = req.Email
	student.Major = req.Major
	student.Semester = req.Semester

	if err := db.Create(&student).Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	return c.Status(201).JSON(models.CreateStudentResponse{Student: student.Student})
}

func UpdateStudentHandler(c *fiber.Ctx) error {
	db := config.GetDB()
	body := c.Body()

	idParam := c.Params("id", "0")
	studentId, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Given student ID is not a positive number"})
	}

	if studentId == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid student ID 0"})
	}

	req := models.CreateStudentRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: err.Error()})
	}

	var student models.StudentModel
	if len(req.NIM) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "nim field is empty"})
	}
	if len(req.Name) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "name field is empty"})
	}
	if len(req.Email) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "email field is empty"})
	}
	if len(req.Major) == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "major field is empty"})
	}
	if req.Semester < 1 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "semester field is less than 1"})
	}

	student.ID = uint(studentId)

	tx := db.Model(&student).Update("nim", req.NIM).Update("name", req.Name).Update("email", req.Email).Update("major", req.Major).Update("semester", req.Semester)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}

	return c.Status(200).JSON(models.CreateStudentResponse{Student: student.Student})
}

func DeleteStudentHandler(c *fiber.Ctx) error {
	db := config.GetDB()

	idParam := c.Params("id", "0")
	studentId, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Given student ID is not a positive number"})
	}

	if studentId == 0 {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid student ID 0"})
	}

	student := models.StudentModel{Student: models.Student{ID: uint(studentId)}}

	tx := db.Delete(&student)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return c.Status(500).JSON(models.InternalServerErrorResponse)
	}
	if tx.RowsAffected == 0 {
		return c.Status(404).JSON(models.NotFoundErrorResponse)
	}

	return c.Status(204).Send([]byte{})
}
