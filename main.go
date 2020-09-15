package main

import (
	"net/http"

	"github.com/chonlatee/simpleredis/cache"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cacheData := cache.NewCache()
	e.GET("/", func(c echo.Context) error {
		data, err := cacheData.Get("data")
		if data == "" && err != nil {
			cacheValue := "just value"
			err = cacheData.Set("data", cacheValue)
			return c.String(http.StatusOK, cacheValue)
		}

		return c.String(http.StatusOK, data+" from cache")

	})
	e.Logger.Fatal(e.Start(":1323"))
}
