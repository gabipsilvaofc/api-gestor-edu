package main

import (
	_ "git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/docs"

	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/routes"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title API GestorEdu
// @version 2.0
// @description Realiza CRUD de usuarios
// @host localhost:3000
// @BasePath / 
// @schemes http

func main() {
	db := database.Db()
	app := fiber.New()

	app.Use(cors.New())
	routes.AppRoutes(app, db)
	

	app.Listen(":3000")
}
