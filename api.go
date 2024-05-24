package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostData struct {
	Message string `json:"message"`
	Name    string `json:"Name"`
}

func main() {
	e := echo.New()

	data := &PostData{
		Message: "Hello",
		Name:    "World",
	}

	e.GET("mohsen/get", func(c echo.Context) error {
		return c.String(http.StatusOK, data.Message)
	})

	e.POST("mohsen/post", func(c echo.Context) error {

		var postData PostData
		if err := c.Bind(&postData); err != nil {
			return err
		}

		data = &postData

		return c.JSON(http.StatusOK, data)

	})

	e.PUT("moshen/put", func(c echo.Context) error {
		var putData PostData
		if err := c.Bind(&putData); err != nil {
			return err
		}
		data.Name = putData.Name
		data.Message = putData.Message

		return c.JSON(http.StatusOK, data)
	})

	e.Start(":8080")
}
