// utils/csv_utils.go
package utils

import (
	"encoding/csv"
	"os"
	// "errors"
)

// ReadStoreData reads store data from a CSV file
func ReadStoreData(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	storeData := make(map[string]string)
	for _, record := range records[1:] { // Skipping header
		storeData[record[0]] = record[1] // store_id: store_name
	}

	return storeData, nil
}
