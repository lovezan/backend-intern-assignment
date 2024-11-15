// services/job_service.go
package services

import (
	"errors"
	"math/rand"
	"sync"
	"time"
	"backend-intern-assignment/models"
	"backend-intern-assignment/utils"
)

var (
	jobs      = make(map[string]*models.Job)
	jobsMutex sync.Mutex
)

// CreateJob creates a new job to process images
func CreateJob(jobID string, imageURLs []string, storeID string) (string, error) {
	jobsMutex.Lock()
	defer jobsMutex.Unlock()

	// Check if job already exists
	if _, exists := jobs[jobID]; exists {
		return "", errors.New("job already exists")
	}

	// Create a new job and add it to the map
	job := &models.Job{
		JobID:     jobID,
		ImageURLs: imageURLs,
		StoreID:   storeID,
		Status:    "ongoing",
	}

	jobs[jobID] = job

	// Start processing the job asynchronously
	go processJob(job)

	return jobID, nil
}

// GetJobStatus returns the status of a job by its ID
func GetJobStatus(jobID string) (string, error) {
	jobsMutex.Lock()
	defer jobsMutex.Unlock()

	job, exists := jobs[jobID]
	if !exists {
		return "", errors.New("job not found")
	}

	return job.Status, nil
}

// processJob processes the job and calculates the perimeter of each image
func processJob(job *models.Job) {
	for _, imageURL := range job.ImageURLs {
		// Simulate downloading the image and calculating its perimeter
		width, height, err := utils.GetImageDimensions(imageURL)
		if err != nil {
			job.Status = "failed"
			return
		}

		perimeter := 2 * (width + height)
		time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond) // Random sleep between 0.1 to 0.4 seconds

		job.Results = append(job.Results, models.Result{
			ImageURL:  imageURL,
			Perimeter: perimeter,
		})
	}

	job.Status = "completed"
}
