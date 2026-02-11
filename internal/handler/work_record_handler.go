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

	// Manual validation for updates to ensure non-empty fields if they are present
	requiredStringFields := []string{"trunk_model", "customer_name", "construction_site"}
	for _, field := range requiredStringFields {
		if val, ok := updates[field]; ok {
			if strVal, ok := val.(string); ok && strVal == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": field + " cannot be empty"})
				return
			}
		}
	}

	// Validate date, quantity, price, and charged presence if intended to be "non-empty" equivalent?
	// For numbers and booleans, "non-empty" usually means strictly they must be valid values.
	// Since it's a map binding, 0 or false are values. Checking for their existence is enough to know user tried to set them.
	// If user sends key "quantity": 0, validation might fail if we require > 0.
	if val, ok := updates["quantity"]; ok {
		// JSON numbers are often float64 in map[string]interface{}
		if fVal, ok := val.(float64); ok && fVal == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "quantity cannot be 0"})
			return
		}
	}
	if val, ok := updates["price"]; ok {
		if fVal, ok := val.(float64); ok && fVal == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "price cannot be 0"})
			return
		}
	}

	// 'charged' can be true or false, both are "non-empty" values.
	// If the key is present, it's fine. nil is not possible in JSON boolean unless explicit null.
	if val, ok := updates["charged"]; ok && val == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "charged cannot be null"})
		return
	}

	if err := h.service.UpdateRecord(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record updated"})
}

// ExportRecords godoc
// @Summary Export work records to Excel
// @Description Export work records based on filters
// @Tags records
// @Produce application/octet-stream
// @Param customer_name query string false "Customer Name"
// @Param trunk_model query string false "Trunk Model"
// @Param date query string false "Date (YYYY-MM-DD)"
// @Success 200 {file} file
// @Failure 500 {object} map[string]string
// @Router /api/records/export [get]
func (h *WorkRecordHandler) ExportRecords(c *gin.Context) {
	var filter model.WorkRecordFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buffer, err := h.service.ExportRecords(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileName := "work_records.xlsx"
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
}
