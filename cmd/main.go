package main

import (
	"app/config"
	"app/repository"
	"app/service"
	"app/api"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadConfig(".env")
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
	catRepo := repository.NewCatRepo(dbConn)
	matchRepo := repository.NewCatMatchRepo(dbConn)

	catSvc := service.NewCatManagementService(catRepo, matchRepo)
	catMatchSvc := service.NewCatMatchService(catRepo, matchRepo)

	userHandler := api.NewUserHandler(userRepo)
	catMgtHandler := api.NewCatManagementHandler(catSvc)
	catMatchHandler := api.NewCatMatchHandler(catMatchSvc)

	router := api.NewRouter(userHandler, catMgtHandler, catMatchHandler)

	router.Setup(app)
	log.Fatal(app.Listen(":8080"))
}
