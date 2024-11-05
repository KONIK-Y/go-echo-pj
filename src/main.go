package main

import (
	"log"
	"net/http"
	cnf "training-pj/src/config"
	"training-pj/src/db"
	"training-pj/src/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	cnf.InitEnv()
	// DB connection
	db, err := db.NewSQLDB("postgres")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	// Include routes
	routes.RegisterUserRoutes(e, db)

	// Index route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	// Start server
	if err := e.Start(":8080"); err != nil {
        log.Fatal("Error starting server: ", err)
    }

	e.Logger.Fatal(e.Start(":8080"))
}

	
