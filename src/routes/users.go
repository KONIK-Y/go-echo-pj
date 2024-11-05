package routes

import (
	"database/sql"
	"training-pj/src/handlers"
	"training-pj/src/repos"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, db *sql.DB) {
    userGroup := e.Group("/users")
	ur := repos.NewUserRepository(db)
	uh := handlers.NewUsersHandler(ur)
	
    userGroup.GET("/", uh.GetUsers)
    userGroup.POST("/", uh.CreateUser)
    userGroup.GET("/:id", uh.GetUser)
    userGroup.PUT("/:id", uh.UpdateUser)
    userGroup.DELETE("/:id", uh.DeleteUser)
}