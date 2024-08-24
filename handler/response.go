package handler

import (
	"fmt"
	"net/http"

	"github.com/Luiz0311/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, status int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(status, gin.H{
		"message": msg,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"data"`
}

type OpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type OpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type CreateOpeningResponse OpeningResponse

type DeleteOpeningResponse OpeningResponse

type ShowOpeningResponse OpeningResponse

type UpdateOpeningResponse OpeningResponse

type ListOpeningsResponse OpeningsResponse
