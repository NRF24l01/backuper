package main

import (
	"github.com/NRF24l01/backuper/core"
	"github.com/NRF24l01/backuper/schemas"

	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)
func main() {
	PRODUCTION_ENV := os.Getenv("RUNTIME_PRODUCTION") == "true"
	if !PRODUCTION_ENV {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("failed to load .env: %v", err)
		}
	}
	
	config, err := core.BuildConfigFromEnv()

	if err != nil {
		log.Fatalf("failed to build config: %v", err)
	}

	e := echo.New()

	if !PRODUCTION_ENV {
		e.Use(echoMw.Logger())
	}
    e.Use(echoMw.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, schemas.Message{Status: "backuper is ok"})
	})

	e.Logger.Fatal(e.Start(config.APPHost))
}