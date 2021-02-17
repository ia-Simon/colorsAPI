package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/ia-Simon/colorsAPI/database"
	"github.com/ia-Simon/colorsAPI/models"
)

//GetColors returns a list of colors from database.
func GetColors(c *fiber.Ctx) error {
	db := database.DBConn
	var colors []models.Color
	db.Find(&colors)
	return c.JSON(colors)
}

//GetColor retrieves a single color from database.
func GetColor(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var color models.Color
	db.First(&color, id)

	if color.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"detail": "color not found",
		})
	}
	return c.JSON(color)
}

//CreateColor creates a new color on database.
func CreateColor(c *fiber.Ctx) error {
	db := database.DBConn
	var color models.Color

	if err := c.BodyParser(&color); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"detail": fmt.Sprintln(err),
		})
	}

	db.Create(&color)
	return c.Status(201).JSON(color)
}

//UpdateColor updates a color on database.
func UpdateColor(c *fiber.Ctx) error {
	return c.JSON("Update color")
}

//DeleteColor deletes a color from database.
func DeleteColor(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var color models.Color
	db.First(&color, id)
	if color.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"detail": "color not found",
		})
	}
	db.Delete(&color)
	return c.SendStatus(204)
}

//RandomColor randomly creates an RGB color according to the db model.
func RandomColor(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())

	red := uint8(rand.Intn(256))
	green := uint8(rand.Intn(256))
	blue := uint8(rand.Intn(256))

	hex := fmt.Sprintf("#%X%X%X", red, green, blue)
	return c.JSON(fiber.Map{
		"hex": hex,
		"r":   red,
		"g":   green,
		"b":   blue,
	})
}
