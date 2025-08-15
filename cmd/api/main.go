package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("app is running..ß.")

	app := fiber.New()

	inMemSlice := []Todo{} // Each element in the inMemSlice is a full type Todo struct{} , not a pointer.

	app.Use(logger.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		fmt.Println("pinging...")
		return c.JSON(fiber.Map{"message": "Server is healthy"})
	})
	// ÷ create
	app.Post("/api/create-todo", func(c *fiber.Ctx) error {
		//  creating a value of type Todo with default zero values
		todo := &Todo{} // address of the value
		// todo := &Todo{} // todo is a *pointer* to a Todo

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(inMemSlice) + 1
		inMemSlice = append(inMemSlice, *todo) // *todo is directly asking the value stored in the value
		// *todo // the actual Todo value stored at that address

		return c.Status(201).JSON(todo)

	})

	app.Get("/api/get-todo", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(inMemSlice)
	})

	//  update
	app.Patch("/api/update/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range inMemSlice {
			if fmt.Sprint(todo.ID) == id {
				inMemSlice[i].Completed = true
				return c.Status(200).JSON(inMemSlice[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/api/delete/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		for i, todo := range inMemSlice {
			if fmt.Sprint(todo.ID) == id {
				inMemSlice = append(inMemSlice[:i], inMemSlice[:i+1]...)
				return c.Status(200).JSON(fiber.Map{"success": "true"})

			}

		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

	})

	log.Fatal(app.Listen(":8000"))
}
