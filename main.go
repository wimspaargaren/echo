// Package main bootstraps the api.
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/wimspaargaren/echo/api"
)

func main() {
	// Echo instance
	e := echo.New()
	api.RegisterHandlers(e, newServer())

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// Start server
		err := e.Start(":3000")
		if errors.Is(http.ErrServerClosed, err) {
			fmt.Println("DONE")
		} else if err != nil {
			fmt.Println("Err")
		}
	}()

	<-c
	fmt.Println("Shutting down server")
	err := e.Shutdown(context.Background())
	if err != nil {
		fmt.Println("unable to shutdown server")
	}
}

type server struct {
	*userHandler
}

func newServer() *server {
	return &server{
		userHandler: newUserHandler(),
	}
}
