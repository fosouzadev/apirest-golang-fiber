package routes

import (
	"github.com/fosouzadev/apirest-golang-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
