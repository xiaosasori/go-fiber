package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xiaosasori/go-fiber/db"
	"github.com/xiaosasori/go-fiber/routes"
)

func main() {
	db.Connect()
	defer db.Repo.Client.Disconnect(nil)
	app := fiber.New()
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
	defer fmt.Print("end main")
}
