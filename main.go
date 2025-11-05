package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// CORS middleware, set for DEV ONLY!!
	app.Use(cors.New())

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	listPositionRepo := repositories.NewListPositionRepository()
	listRepo := repositories.NewListRepository()
	listService := services.NewListService(listRepo, boardRepo, listPositionRepo)
	listController := controllers.NewListController(listService)

	routes.Setup(app, userController, boardController, listController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))
}
