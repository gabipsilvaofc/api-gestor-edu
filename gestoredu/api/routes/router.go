package routes

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func AppRoutes(app *fiber.App, db *gorm.DB) {
	usersController := controller.NewUsersController(db)

	// v1 := app.Group("/v1")
	app.Post("/user", usersController.CreateUser)
	app.Patch("/user/:id", usersController.UpdateUser)
	app.Get("/user/:id", usersController.GetUser)
	app.Delete("/user/:id", usersController.DeleteUser)
	app.Get("/swagger/*", swagger.HandlerDefault)

}
