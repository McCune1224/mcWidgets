package main

import (
	"mcwidgets/templates/pages"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5173"
	}
	return port
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.IndexPage())
}

func main() {
	app := echo.New()
	app.GET("/", HomeHandler)
	// load static files
	app.Static("/", "static")
	port := getPort()
	app.Logger.Fatal(app.Start(":" + port))
}
