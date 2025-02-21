package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/javed2214/arctic-wolf-assignment/config"
	"github.com/javed2214/arctic-wolf-assignment/routes"
)

func InitServer() {

	config.InitConfig()
	conf := config.GetConfig()
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(conf.AppConfig.Port)
	if err != nil {
		log.Fatalf("Error initializing server. Error: %s", err.Error())
	}
}
