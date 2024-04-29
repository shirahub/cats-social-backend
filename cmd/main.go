package main

import (
	"app/repository"
	"app/api"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})
	// app.Use(cors.New())

	dbConn := repository.NewPostgresConnection()
	userRepo := repository.NewUserRepo(dbConn)

	router := api.NewRouter(api.NewUserHandler(userRepo))

	router.Setup(app)
	log.Fatal(app.Listen(":8080"))
}
