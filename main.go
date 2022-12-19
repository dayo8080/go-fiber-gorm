package main

import (
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)


func main() {

	app := fiber.New()

	// INITIAL ROUTE

	route.RouteInit(app)
	
app.Listen(":3000")

}