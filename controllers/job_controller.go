// controllers/job_controller.go
package controllers

import (
	"net/http"
	"backend-intern-assignment/services"
	"github.com/gin-gonic/gin"
)

// SubmitJob handles the creation of a new job
func SubmitJob(c *gin.Context) {
	var request struct {
		JobID     string   `json:"job_id"`
		ImageURLs []string `json:"image_urls"`
		StoreID   string   `json:"store_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	jobID, err := services.CreateJob(request.JobID, request.ImageURLs, request.StoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job created successfully", "job_id": jobID})
}

// GetJobStatus handles fetching the status of a job
func GetJobStatus(c *gin.Context) {
	jobID := c.Query("jobid")
	if jobID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job ID is required"})
		return
	}

	status, err := services.GetJobStatus(jobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job_status": status})
}
