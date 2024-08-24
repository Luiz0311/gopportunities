package router

import (
	docs "github.com/Luiz0311/gopportunities/docs"
	"github.com/Luiz0311/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api"
	handler.InitializeHandler()
	docs.SwaggerInfo.BasePath = basePath

	api := r.Group(basePath)
	{
		api.GET("/opening", handler.ShowOpeningHandler)
		api.POST("/opening", handler.CreateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
		api.PUT("/opening", handler.UpdateOpeningHandler)
		api.GET("/openings", handler.ListOpeningsHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
