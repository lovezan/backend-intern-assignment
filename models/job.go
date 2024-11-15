// models/job.go
package models

type Job struct {
	JobID     string   `json:"job_id"`
	ImageURLs []string `json:"image_urls"`
	StoreID   string   `json:"store_id"`
	Status    string   `json:"status"`
	Results   []Result `json:"results"`
}

type Result struct {
	ImageURL  string `json:"image_url"`
	Perimeter int    `json:"perimeter"`
}
