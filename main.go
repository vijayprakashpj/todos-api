package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	todos := []Todo{}

	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if error := c.BodyParser(todo); error != nil {
			return error
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Patch("/todos/:id/done", func(c *fiber.Ctx) error {
		id, error := c.ParamsInt("id")

		if error != nil {
			return c.Status(400).SendString("invalid id")
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		id, error := c.ParamsInt("id")

		if error != nil {
			return c.Status(400).SendString("invalid id")
		}

		for i, todo := range todos {
			if todo.ID == id {
				return c.JSON(todos[i])
			}
		}

		return c.Status(404).SendString("todo not found")
	})

	log.Fatal(app.Listen(":5050"))
}
