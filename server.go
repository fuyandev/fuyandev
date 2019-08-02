package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	e.AutoTLSManager.Cache = autocert.DirCache(".cache")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.StartAutoTLS(":443"))
}

// Handler
func hello(c echo.Context) error {
	return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
	`)
}
