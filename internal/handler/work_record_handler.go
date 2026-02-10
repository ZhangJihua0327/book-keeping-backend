package handler

import (
	"book-keeping-backend/internal/model"
	"book-keeping-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkRecordHandler struct {
	service *service.WorkRecordService
}

func NewWorkRecordHandler(s *service.WorkRecordService) *WorkRecordHandler {
	return &WorkRecordHandler{service: s}
}

// AddRecord godoc
// @Summary Add a new work record
// @Description Add a new work record to the database
// @Tags records
// @Accept json
// @Produce json
// @Param record body model.WorkRecord true "Work Record"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/records [post]
func (h *WorkRecordHandler) AddRecord(c *gin.Context) {
	var record model.WorkRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AddRecord(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record added", "id": record.ID})
}

// GetRecords godoc
// @Summary Get work records by date
// @Description Get all work records created on a specific date, ordered by ID descending
// @Tags records
// @Produce json
// @Param date query string true "Date (YYYY-MM-DD)"
// @Success 200 {array} model.WorkRecord
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/records [get]
func (h *WorkRecordHandler) GetRecords(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date parameter is required"})
		return
	}
	records, err := h.service.GetRecordsByDate(dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

// UpdateRecord godoc
// @Summary Update a work record
// @Description Update an existing work record by ID
// @Tags records
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Param updates body model.WorkRecord true "Updates (pass only fields to update)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/records/{id} [put]
func (h *WorkRecordHandler) UpdateRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateRecord(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record updated"})
}
