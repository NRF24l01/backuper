package core

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	APPHost                 string

	PGHost                  string
	PGPort                  string
	PGUser                  string
	PGPassword              string
	PGDatabase              string
	PGSSLMode               string
	PGTimeZone              string

	MinioEndpoint           string
	MinioUser               string
	MinioPass               string
	MinioBaseUrl            string
	MinioPresignedLifetime  int

	PasswordSalt            string
	JWTAccessSecret         string
	JWTRefreshSecret        string
	
	AllowOrigins            string
}

func BuildConfigFromEnv() (*Config, error) {
	minioPresignedLifetime, err := strconv.Atoi(os.Getenv("MINIO_PRESIGNED_LIFETIME"))
	if err != nil {
		return nil, fmt.Errorf("invalid MINIO_PRESIGNED_LIFETIME: %v", err)
	}

	config := &Config{
		APPHost:                os.Getenv("APP_HOST"),
		PGHost:                 os.Getenv("POSTGRES_HOST"),
		PGPort:                 os.Getenv("POSTGRES_PORT"),
		PGUser:                 os.Getenv("POSTGRES_USER"),
		PGPassword:             os.Getenv("POSTGRES_PASSWORD"),
		PGDatabase:             os.Getenv("POSTGRES_DB"),
		PGSSLMode:              os.Getenv("POSTGRES_SSLMODE"),
		PGTimeZone:             os.Getenv("POSTGRES_TIMEZONE"),
		MinioEndpoint:          os.Getenv("MINIO_ENDPOINT"),
		MinioUser:              os.Getenv("MINIO_USERNAME"),
		MinioPass:              os.Getenv("MINIO_PASSWORD"),
		MinioBaseUrl:           os.Getenv("MINIO_BASE_URL"),
		MinioPresignedLifetime: minioPresignedLifetime,
		PasswordSalt:           os.Getenv("PASSWORD_SALT"),
		JWTAccessSecret:        os.Getenv("JWT_ACCESS_SECRET"),
		JWTRefreshSecret:       os.Getenv("JWT_REFRESH_SECRET"),
		AllowOrigins:           os.Getenv("ALLOWED_ORIGINS"),
	}

	return config, nil
}