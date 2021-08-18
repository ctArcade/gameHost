package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New()

	// s.Static("/platformer-game", "./public/platformer-game")

	s.Get("/games/:author/:title", func(c *fiber.Ctx) error {
		filePath := fmt.Sprintf("./public/%s/%s/index.html", c.Params("author"), c.Params("title"))

		fmt.Println(filePath)

		return c.SendFile(filePath)
	})

	s.Listen(":3000")
}
