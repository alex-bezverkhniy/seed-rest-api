package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService UserService
}

type ResponseHTTP struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Create a new for User
func NewUserHandler(userRoute fiber.Router, us UserService) {
	handler := &UserHandler{
		userService: us,
	}

	userRoute.Get("", handler.getUsers)
}

// Gets  get all users from database
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]user.User}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users [get]
func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	users, err := h.userService.GetUsers(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Status: "success",
		Data:   users,
	})
}
