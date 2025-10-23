package main

import (
	"REST_API/db"
	"REST_API/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8090")
	if err != nil {
		panic("An error occurred in server")
	}

}
