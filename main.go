// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"backend-intern-assignment/controllers"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/api/submit", controllers.SubmitJob)
	r.GET("/api/status", controllers.GetJobStatus)

	// Run the server
	r.Run(":8080")
}
