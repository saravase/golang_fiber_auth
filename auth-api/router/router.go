package router

import (
	"golang_fiber_auth/auth-api/handler"
	"golang_fiber_auth/auth-api/middleware"

	"github.com/gofiber/fiber/v2"
)

// InitRoutes used to initialize the routes
func InitRoutes(app *fiber.App, auth *handler.Auth) {

	// auth
	api := app.Group("/api")
	api.Post("/signup", auth.SignupHandler)
	api.Post("/signin", auth.SigninHandler)
	api.Post("/signout", auth.SignoutHandler)

	// plant
	api.Get("/plants", middleware.JWTMiddleware, auth.GetPlantsHandler)
}
