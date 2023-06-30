package main

import (
	"firebond/database"
	"firebond/response"
	"firebond/router"

	"github.com/joho/godotenv"
	logs "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logs.WithFields(response.StandardFields).Fatalf("leError in reading the config fi: %v", err)
	}
}

func main()  {
	
	/* CALL DATABASE */
	database.DatabaseConnection()

	/* CALL ROUTES */
	router.RouteSetUp()
}
