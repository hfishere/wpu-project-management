package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hfishere/wpu-project-management/config"
	"github.com/hfishere/wpu-project-management/controllers"
	"github.com/hfishere/wpu-project-management/database/seed"
	"github.com/hfishere/wpu-project-management/repositories"
	"github.com/hfishere/wpu-project-management/routes"
	"github.com/hfishere/wpu-project-management/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	boardRepo := repositories.NewBoardRepository()
	boardService := services.NewBoardService(boardRepo, userRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))
}
