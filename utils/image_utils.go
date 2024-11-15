// utils/image_utils.go
package utils

import (
	"errors"
	"math/rand"
)

// GetImageDimensions is a mock function to simulate fetching image dimensions
func GetImageDimensions(imageURL string) (int, int, error) {
	if imageURL == "" {
		return 0, 0, errors.New("invalid image URL")
	}

	// Simulating random dimensions
	width := 100 + rand.Intn(200)
	height := 100 + rand.Intn(200)
	return width, height, nil
}
