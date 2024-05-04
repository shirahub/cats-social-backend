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
		ServerHeader:  "Fiber",
		AppName:       "Cats Social @codingsa",
	})
	// app.Use(cors.New())

	dbConn := repository.NewPostgresConnection()
	defer repository.CloseDbConn(dbConn)

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
