package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javed2214/arctic-wolf-assignment/internal/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/v1/risks", handlers.GetRisks)
	app.Post("/v1/risks", handlers.CreateRisk)
	app.Get("/v1/risks/:id", handlers.GetRiskByID)
}
