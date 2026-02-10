package handler

import (
	"book-keeping-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrunkModelHandler struct {
	service *service.TrunkModelService
}

func NewTrunkModelHandler(s *service.TrunkModelService) *TrunkModelHandler {
	return &TrunkModelHandler{service: s}
}

type CreateTrunkModelRequest struct {
	TrunkModel string `json:"trunk_model" binding:"required"`
}

// AddTrunkModel godoc
// @Summary Add a new trunk model
// @Description Add a new trunk model to the database
// @Tags models
// @Accept json
// @Produce json
// @Param model body CreateTrunkModelRequest true "Trunk Model"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/models [post]
func (h *TrunkModelHandler) AddTrunkModel(c *gin.Context) {
	var req CreateTrunkModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AddTrunkModel(req.TrunkModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Trunk model added"})
}

// GetAllTrunkModels godoc
// @Summary Get all trunk models
// @Description Get a list of all trunk models
// @Tags models
// @Produce json
// @Success 200 {array} model.TrunkModel
// @Failure 500 {object} map[string]string
// @Router /api/models [get]
func (h *TrunkModelHandler) GetAllTrunkModels(c *gin.Context) {
	models, err := h.service.GetAllTrunkModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models)
}
