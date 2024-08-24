package handler

import (
	"net/http"

	"github.com/Luiz0311/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opening")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
