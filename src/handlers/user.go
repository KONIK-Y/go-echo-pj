package handlers

import (
	"net/http"
	"training-pj/src/models"
	"training-pj/src/repos"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	repo repos.UserRepository
}

func NewUsersHandler(repo repos.UserRepository) *UsersHandler {
	return &UsersHandler{repo: repo}
}



func (h *UsersHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	user.CreatedAt = user.CreatedAt.UTC()
	if err := h.repo.CreateUser(c.Request().Context(), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UsersHandler) GetUser(c echo.Context) error {
	id:= c.Param("id")
	if id != "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}
	user, err := h.repo.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UsersHandler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UsersHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.repo.UpdateUser(c.Request().Context(), id, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "User updated"})
}

func (h *UsersHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}
	if err := h.repo.DeleteUser(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "User deleted"})
}