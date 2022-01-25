package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Ticker struct {
	Symbol string `json:"symbol"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice string `json:"lastPrice"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Start routing with /ticker/:symbol")
	})

	// Get single ticker
	e.GET("/ticker/:symbol", func(c echo.Context) error {
		symbol := c.Param("symbol")
		ticker, err := getTicker(symbol)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ticker)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func getTicker(symbol string) (Ticker, error) {
	resp, err := http.Get("https://api.binance.com/api/v1/ticker/24hr?symbol=" + symbol + "USDT")
	if err != nil {
		return Ticker{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Ticker{}, err
	}
	var ticker Ticker
	err = json.Unmarshal(body, &ticker)
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}