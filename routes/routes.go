package routes

import (
	"coditeach/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	{
		//Authorization routes
		auth.Post("/sign_up", controllers.SignUp)
		auth.Post("/sign_in", controllers.SignIn)
		auth.Post("/refresh", controllers.Refresh)
	}

}
