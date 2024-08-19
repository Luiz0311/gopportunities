package router

import (
	"github.com/Luiz0311/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/opening", handler.ShowOpeningHandler)
		api.POST("/opening", handler.CreateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
		api.PUT("/opening", handler.UpdateOpeningHandler)
		api.GET("/openings", handler.ListOpeningsHandler)
	}
}
