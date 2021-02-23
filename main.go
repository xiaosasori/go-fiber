package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/xiaosasori/go-fiber/db"
	"github.com/xiaosasori/go-fiber/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db.Connect()
	defer db.Repo.Client.Disconnect(nil)
	app := fiber.New()
	routes.Setup(app)
	fmt.Print(os.Getenv("PORT"))
	log.Fatal(app.Listen(":3000"))
	defer fmt.Print("end main")
	defer db.Repo.Client.Disconnect(context.Background())
}
