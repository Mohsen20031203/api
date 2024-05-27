package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostData struct {
	Lastname string `json:"Lastname"`
	Name     string `json:"Name"`
}

func main() {
	e := echo.New()

	data := &PostData{
		Lastname: "Hello",
		Name:     "World",
	}

	e.GET("mohsen/get", func(c echo.Context) error {
		c.Response().Header().Set("hi", "hello")
		return c.String(http.StatusOK, data.Lastname)
	})

	e.POST("mohsen/post", func(c echo.Context) error {
		contentType := c.Request().Header.Get("Authorization")
		if contentType != "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error 400": "Your token is not valid"})
		}

		c.Response().Header().Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

		var postData PostData
		if err := c.Bind(&postData); err != nil {
			return err
		}

		data = &postData

		return c.JSON(http.StatusOK, data)

	})

	e.PUT("mohsen/put", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/jsonnnnn")
		c.Response().Header().Set("Cache-Control", "max-age=3600, public")
		var putData PostData
		if err := c.Bind(&putData); err != nil {
			return err
		}
		data.Name = putData.Name
		data.Lastname = putData.Lastname

		return c.JSON(http.StatusOK, data)
	})

	e.Start(":8080")
}
