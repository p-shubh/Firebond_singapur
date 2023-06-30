package router

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/{cryptocurrency}/{fiat}")
		api.GET("/{cryptocurrency}")
		api.GET("")
		api.GET("/history/{cryptocurrency}/{fiat}")
	}
}
