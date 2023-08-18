package v1

import "github.com/gofiber/fiber/v2"

func Bind(router fiber.Router) {
	router.Get("/search", SearchReport)
}
