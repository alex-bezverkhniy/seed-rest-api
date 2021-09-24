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
	userRoute.Get("/:userID", handler.getUsers)
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

// Gets get user by ID from database
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=user.User}
// @Success 404 {object} ResponseHTTP{}
// @Success 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users/{userId} [get]
func (h *UserHandler) getUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Status:  "fail",
			Message: "Please specify a valid user ID",
		})
	}

	user, err := h.userService.GetUser(customContext, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	if user == nil || user.ID == -1 {
		return c.Status(fiber.StatusNotFound).JSON(ResponseHTTP{
			Status:  "fail",
			Message: "User with specified ID is not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Status: "success",
		Data:   user,
	})
}
