package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello world 🌈")
	// fiber instance
	app := fiber.New()

	// routes
	app.Get("/", func(c *fiber.Ctx) error { //string
		return c.SendString("hello world 🌈")
	})

	app.Get("/info", func(c *fiber.Ctx) error { // JSON
		return c.JSON(fiber.Map{
			"msg":     "hello world 🚀",
			"go":      "fiber 🥦",
			"boolean": true,
			"number":  1234,
		})
	})

	type Food struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}

	var foods = []Food{
		{ID: 1, Name: "ต้มยำกุ้ง", Price: 140},
		{ID: 2, Name: "ไก่ทอด", Price: 100},
		{ID: 3, Name: "ก๋วยเตี๋ยว", Price: 30},
		{ID: 4, Name: "เบอร์เกอร์", Price: 149},
	}

	type EditFood struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}

	app.Get("/foods", func(c *fiber.Ctx) error {
		return c.JSON(foods)
	})

	app.Get("/foods/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, food := range foods {
			if fmt.Sprint(food.ID) == id {
				return c.JSON(food)
			}
		}
		return nil
	})

	app.Post("/foods", func(c *fiber.Ctx) error {
		var food Food
		if err := c.BodyParser(&food); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		foods = append(foods, food)
		return c.JSON(food)
	})

	app.Put("/foods/:id", func(c *fiber.Ctx) error {
		var editFood EditFood
		if err := c.BodyParser(&editFood); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		id := c.Params("id")
		for i := range foods {
			if fmt.Sprint(foods[i].ID) == id {
				foods[i].Name = editFood.Name
				foods[i].Price = editFood.Price
				return c.JSON(foods[i])
			}
		}
		return nil
	})

	app.Delete("/foods/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, food := range foods {
			if fmt.Sprint(food.ID) == id {
				foods = append(foods[:i], foods[i+1:]...)
				return c.SendString("delete success")
			}
		}
		return nil
	})

	// app listening at PORT: 3000
	app.Listen(":3000")
}
