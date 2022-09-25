package routes

import (
	"coditeach/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//Authorization routes
	app.Post("api/v1/auth/sign_up", controllers.SignUp)
	app.Post("api/v1/auth/sign_in", controllers.SignIn)
	app.Post("api/v1/auth/refresh", controllers.Refresh)
}
