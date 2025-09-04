package handlers

import (
	"net/http"

	"github.com/google/uuid"
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

func (h *Handler) WorkerListHandler(c echo.Context) error {
	var workers []models.Worker
	if err := h.DB.Find(&workers).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	response := schemas.WorkerListResponse{
		Workers: make([]schemas.WorkerStatus, len(workers)),
	}
	for i, worker := range workers {
		response.Workers[i] = schemas.WorkerStatus{
			WorkerUUID: worker.ID.String(),
			Name:      worker.Name,
			CreatedAt: worker.CreatedAt.Unix(),
			LastOnline: worker.LastOnline,
		}
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) WorkerGetBackupTasksHandler(c echo.Context) error {
	var tasks []models.WorkerCapability
	workerID := c.Param("id")
	
	parsedUUID, err := uuid.Parse(workerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid UUID")
	}
	if parsedUUID.Version() != 4 {
		return echo.NewHTTPError(http.StatusBadRequest, "UUID must be version 4")
	}

	if err := h.DB.Where("worker_id = ?", workerID).Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}
	
	tasksResp := make([]schemas.BackupTask, len(tasks))
	for i, task := range tasks {
		tasksResp[i] = schemas.BackupTask{
			Type:    task.Type,
			About:   task.About,
			Freq:    task.BackupInterval,
			LastBck: task.LastBck,
		}
	}

	return c.JSON(http.StatusOK, schemas.WorkerGetBackupTasksResponse{Tasks: tasksResp})
}