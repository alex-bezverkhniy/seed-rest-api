package swagger

import (
	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
)

// Create a new for Swagger
func NewSwaggerHandler(userRoute fiber.Router) {
	userRoute.Get("", swagger.Handler)
}
