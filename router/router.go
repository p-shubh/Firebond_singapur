package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RouteSetUp() {
	router := gin.Default()
	gin.SetMode(os.Getenv("APP_MODE"))
	Routes(router)
	router.Run(`:` + os.Getenv("PORT"))

}

func Routes(c *gin.Engine) {
	ApplyRoutes(c)
}
