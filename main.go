package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Welcome)
	e.GET("/get-user", GetUser)
	e.GET("/get-order", GetOrder)
	e.GET("/get-product", GetProduct)
	e.GET("/get-city", GetCity)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

type jsonResponse struct {
	Data   string `json:"data"`
	Status bool   `json:"status"`
}

// Handler
func Welcome(c echo.Context) error {
	welcome := fmt.Sprintln("Welcome To Website Test API \n 1. /get-user \n 2. /get-order \n 3. /get-product")

	return c.String(http.StatusOK, welcome)
}

func GetUser(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data User Berhasil di Get",
		Status: true,
	}

	return c.JSON(http.StatusOK, response)
}

func GetOrder(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data Order Berhasil di Get",
		Status: true,
	}

	return c.JSON(http.StatusOK, response)
}

func GetProduct(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data Product Berhasil di Get",
		Status: true,
	}

	return c.JSON(http.StatusOK, response)
}

func GetCity(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data City Berhasil di Get",
		Status: true,
	}

	return c.JSON(http.StatusOK, response)
}
