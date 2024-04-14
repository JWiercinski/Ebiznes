package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func kill(c echo.Context) error {
	err := c.Echo().Shutdown(context.Background())
	if err != nil {
		if err != http.ErrServerClosed {
			c.Echo().Logger.Fatal("Serwer wyłączony")
		}
	}
	return nil
}
