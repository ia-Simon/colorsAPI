package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ia-Simon/colorsAPI/controllers"
)

//Colors is a sub-router for color endpoints.
var Colors *fiber.App

func init() {
	Colors = fiber.New()

	Colors.Get("/", controllers.GetColors)
	Colors.Get("/random", controllers.RandomColor)
	Colors.Get("/:id", controllers.GetColor)
	Colors.Post("/", controllers.CreateColor)
	Colors.Put(":id", controllers.UpdateColor)
	Colors.Delete("/:id", controllers.DeleteColor)

}
