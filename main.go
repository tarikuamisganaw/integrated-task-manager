package main

import (
	"task-mongo/data"
	"task-mongo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB using the provided connection string
	data.ConnectDB("mongodb://localhost:27017/")

	r := gin.Default()
	router.SetupRouter(r)
	r.Run()
}
