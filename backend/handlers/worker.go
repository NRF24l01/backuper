package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrf24l01/backuper/models"
	"github.com/nrf24l01/backuper/schemas"
)

func (h *Handler) WorkerCreateHandler(c echo.Context) error {
	worker_data := c.Get("validatedBody").(*schemas.WorkerCreateRequest)

	var worker models.Worker
	if err := h.DB.Where("name = ?", worker_data.Name).First(&worker).Error; err == nil {
		return echo.NewHTTPError(http.StatusConflict, "Worker already exists")
	}

	worker.Name = worker_data.Name
	if err := worker.NewToken(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	if err := h.DB.Create(&worker).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	worker_response := schemas.WorkerCreateResponse{
		WorkerUUID:   worker.ID.String(),
		Token: worker.Token,
	}

	return c.JSON(http.StatusCreated, worker_response)
}