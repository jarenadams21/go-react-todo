package main

// c * fiber.Ctx
// = have fiber context point to c variable

import (
	"fmt"
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
	fmt.Print("Hello world")

	// Infer the type from whatever New() returns
	// It returns a pointer to an app instance of type App struct
	// Server creation
	app := fiber.New()

	// Avoid CORS issues
	// Define our type for the config
	// Allow origins (whitlisted urls)
	// Allow headers
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	// Could also return an error (error syntax)
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// Return errors
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		// give todo an ID
		todo.ID = len(todos) + 1

		// Append new todo to array of todos
		todos = append(todos, *todo)

		// Return slice of todos we have
		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		// i = index
		// t = todo
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}

		}
// c = Context 
		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
