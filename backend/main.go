package main

import (
	"github.com/NRF24l01/backuper/core"
	"github.com/NRF24l01/backuper/handlers"
	"github.com/NRF24l01/backuper/models"
	"github.com/NRF24l01/backuper/routes"
	"github.com/NRF24l01/backuper/schemas"
	"github.com/NRF24l01/go-web-utils/s3util"

	"log"
	"os"
	"strings"

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

	// Data providers
	db := models.RegisterPostgres(config)
	// create s3util client from config
	s3cfg := s3util.S3Config{
		Endpoint:  config.MinioEndpoint,
		AccessKey: config.MinioUser,
		SecretKey: config.MinioPass,
		UseSSL:    strings.HasPrefix(strings.ToLower(config.MinioBaseUrl), "https"),
		BaseURL:   config.MinioBaseUrl,
	}
	s3client, err := s3util.New(s3cfg)
	if err != nil {
		log.Fatalf("failed to create S3 client: %v", err)
	}

	e := echo.New()

	if !PRODUCTION_ENV {
		e.Use(echoMw.Logger())
	}
    e.Use(echoMw.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, schemas.Message{Status: "backuper is ok"})
	})

	handler := &handlers.Handler{DB: db, MinIOClient: s3client, Config: config}
	routes.RegisterRoutes(e, handler)

	e.Logger.Fatal(e.Start(config.APPHost))
}