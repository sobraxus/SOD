package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sobraxus/SOD/internal/models"
	"github.com/sobraxus/SOD/internal/repositories"
)

type CaseHandler struct {
	repo *repositories.CaseRepository
}

func NewCaseHandler(repo *repositories.CaseRepository) *CaseHandler {
	return &CaseHandler{repo: repo}
}

// CreateCase godoc
// @Summary Create a new case
// @Accept json
// @Produce json
// @Param case body models.Case true "Case data"
// @Success 201 {object} models.Case
// @Router /cases [post]
func (h *CaseHandler) CreateCase(c *gin.Context) {
	var newCase models.Case
	if err := c.ShouldBindJSON(&newCase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Generate UUID and timestamps
	caseToCreate := models.NewCase(newCase.Title, newCase.Description)
	if err := h.repo.CreateCase(c.Request.Context(), caseToCreate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create case"})
		return
	}

	c.JSON(http.StatusCreated, caseToCreate)
}

// GetCaseByID godoc
// @Summary Get a case by ID
// @Produce json
// @Param id path string true "Case ID"
// @Success 200 {object} models.Case
// @Router /cases/{id} [get]
func (h *CaseHandler) GetCaseByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid case ID"})
		return
	}

	caseData, err := h.repo.GetCaseByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "case not found"})
		return
	}

	c.JSON(http.StatusOK, caseData)
}
