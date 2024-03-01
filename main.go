package main

import (
	"context"
	"net/http"

	"module/db"
	"module/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// MAIN MENU
	component := templates.Index()
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/dados", func(c echo.Context) error {
		// Chama a função que acessa o banco de dados
		dados, err := db.Db_handler()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro interno do servidor"})
		}
		return c.JSON(http.StatusOK, dados)
	})

	e.Static("/css", "css")

	// Iniciar o servidor Echo
	e.Logger.Fatal(e.Start(":3000"))
}
