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
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Create a new for User
func NewUserHandler(userRoute fiber.Router, us UserService) {
	handler := &UserHandler{
		userService: us,
	}

	userRoute.Get("", handler.getUsers)
	userRoute.Post("", handler.createUser)
	userRoute.Get("/:userID", handler.getUser)
	userRoute.Put("/:userID", handler.updateUser)
	userRoute.Delete("/:userID", handler.deleteUser)
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
// @Param userID path int true "User ID"
// @Success 200 {object} ResponseHTTP{data=user.User}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users/{userID} [get]
func (h *UserHandler) getUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Status:  "fail",
			Message: "Please specify a valid user ID",
		})
	}

	user, err := h.userService.GetUser(customContext, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	if user == nil || user.ID != userID {
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

// Creates new user in the database
// @Summary Creates new user
// @Description Creates new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user.User true "Create user"
// @Success 202 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users [post]
func (h *UserHandler) createUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err := h.userService.CreateUser(customContext, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(ResponseHTTP{
		Status:  "success",
		Message: "User created",
	})
}

// Updates user in the database
// @Summary Updates user
// @Description Updates user
// @Tags user
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param user body user.User true "Update user"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users/{userID} [put]
func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Status:  "fail",
			Message: "Please specify a valid user ID",
		})
	}

	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	err = h.userService.UpdateUser(customContext, userID, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Status:  "success",
		Message: "User updated",
	})
}

// Deletes user in the database
// @Summary Deletes user
// @Description Deletes user
// @Tags user
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 204
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/users/{userID} [delete]
func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Status:  "fail",
			Message: "Please specify a valid user ID",
		})
	}

	err = h.userService.DeleteUser(customContext, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Status:  "fail",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
