package main

import (
	"github.com/gofiber/fiber/v2"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/ia-Simon/colorsAPI/database"
	"github.com/ia-Simon/colorsAPI/routes"
)

func main() {
	database.InitDBConn()
	defer database.DBConn.Close()

	app := fiber.New()
	app.Use(recovery.New())

	app.Mount("/colors", routes.Colors)

	app.Listen("localhost:4000")
}
