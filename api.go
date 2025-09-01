package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostData struct {
	Lastname string `json:"Lastname"`
	Name     string `json:"Name"`
	Phone    string `json:"Phone"`
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

/*package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

type result struct {
	Result struct {
		List []struct {
			Price  string `json:"price"`
			Symbol string `json:"symbol"`
		} `json:"list"`
	} `json:"result"`
}

func getPrice(db *leveldb.DB) error {
	prices := make([]float64, 20)
	for {

		rest, err := http.Get("https://api.bybit.com/spot/v3/public/quote/ticker/price")
		if err != nil {
			return err
		}
		defer rest.Body.Close()

		var responseData result
		if err := json.NewDecoder(rest.Body).Decode(&responseData); err != nil {
			return err
		}

		for _, v := range responseData.Result.List {

			if strings.HasSuffix(v.Symbol, "USDT") {

				//
				// if err := db.Put([]byte(v.Symbol), []byte(v.Price), nil); err != nil {
				// 	return err
				// }

				val, err := db.Get([]byte(v.Symbol), nil)
				if err != nil {
					return err
				}

				strVal := string(v.Price)

				floatVal, err := strconv.ParseFloat(strVal, 64)
				if err != nil {
					fmt.Println("Error:", err)
				}

				prices = append(prices, floatVal)

				byteArray, err := json.Marshal(val)
				if err != nil {
					fmt.Println("Error:", err)
				}

				if err := db.Put([]byte(v.Symbol), byteArray, nil); err != nil {
					return err
				}

				strData := string(val)

				fmt.Printf("Symbol: %s, New Value: %s\n", v.Symbol, strData)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
func main() {

	db, err := leveldb.OpenFile("example_db3", nil)
	if err != nil {
		log.Fatal(err)
	}
	getPrice(db)
	defer db.Close()

	// go GetinApi(db)
	// go SearchinData(db, "BTCUSDT")
}*/
