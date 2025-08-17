package main

import (
	"log"
	"pertemuan6/handlers"
	"pertemuan6/middlewares"
	"pertemuan6/models"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	// Database migration and seeding
	models.MigrateAll()
	models.SeedUserAdmin()

	// Initialize fiber and use needed middlewares
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/swagger/*", func(c fiber.Ctx) error { return nil })

	api := app.Group("/api")

	// 1. Authentication Endpoints
	auth := api.Group("/auth")
	auth.Post("/login", handlers.LoginHandler)
	auth.Post("/register", handlers.RegisterHandler)

	protectedApi := api.Use(middlewares.JwtMiddleware)

	// 2. Student Management Endpoints (Protected)
	student := protectedApi.Group("/students/")
	student.Get("/", handlers.GetAllStudentHandler)
	student.Get("/:id", handlers.GetStudentByIdHandler)
	student.Post("/", middlewares.AdminMiddleware, handlers.CreateStudentHandler)
	student.Put("/:id", middlewares.AdminMiddleware, handlers.UpdateStudentHandler)
	student.Delete("/:id", middlewares.AdminMiddleware, handlers.DeleteStudentHandler)

	// 3. Profile Endpoint (Protected)
	protectedApi.Get("/profile", handlers.GetProfileHandler)

	log.Fatal(app.Listen(":8082"))
}
