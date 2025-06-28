package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main(){
	fmt.Println("Hello, Simon")
	app := fiber.New()

	//endpoint
	app.Get("/",func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg":"Hello Simon, this is your first endpoint using Go."})
	})
	//Puerto del servidor local
	log.Fatal(app.Listen(":4000"))
}