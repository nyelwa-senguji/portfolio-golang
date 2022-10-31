package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Create a new engine
	engine := html.New("./template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Serving static files
	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":"nyelwa-senguji portfolio",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
