package auth

import "github.com/gofiber/fiber/v2"

// Create auth handler
type AuthHandler struct{}

type HttpResponse struct {
	status  string `json:"status"`
	message string `json:"message"`
}

func NewAuthHandler(authRouter fiber.Router) {
	handler := &AuthHandler{}

	authRouter.Post("/login", handler.signInUser)
}

// Sings user and provide jwt token
// @Summary Sings user
// @Description Sings user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body user.User true "JWT token"
// @Success 202 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users [post]
func (h *AuthHandler) signInUser(c *fiber.Ctx) error {

	// Structure for login requests
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HttpResponse{
			Message: err.Error(),
			Status:  "fail",
		})
	}

}
