package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* func RouteSetUp() {
	router := gin.Default()
	gin.SetMode(os.Getenv("APP_MODE"))
	Routes(router)
	router.Run(`:` + os.Getenv("PORT"))

}

func Routes(c *gin.Engine) {
	ApplyRoutes(c)
} */

func RouteSetUp() {
	router := mux.NewRouter()
	Routes(router)
	err := http.ListenAndServe(":8080", router)
	if err!= nil {
		log.Fatal("Failed To Start The Server")
	}else{
		log.Println("Server Started Successfully On Port 8080...")
	}

}

func Routes(c *mux.Router) {
	ApplyRoutes(c)
}
