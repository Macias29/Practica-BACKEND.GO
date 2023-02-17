package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type user struct {
	id        string
	Firstname string
	lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := user{

		Firstname: "Ricardo",
		lastname:  "Macias",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
func handleCreateUser(c *fiber.Ctx) error {
	user := user{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Desarrollo de Sistema Informatico")
	})

	app.Get("/user", handleUser)
	app.Post("/user", handleCreateUser)
	app.Listen(":4000")
}
