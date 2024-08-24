package handler

import (
	"net/http"

	"github.com/Luiz0311/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opening")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
