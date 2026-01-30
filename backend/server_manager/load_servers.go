package servermanager

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadServers(filePath string) ([]Server, error) {
	var servers []Server

	// Make sure the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("servers file does not exist: %s", filePath)
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read servers file: %v", err)
	}

	// Unmarshal JSON into slice of Server structs
	if err := json.Unmarshal(data, &servers); err != nil {
		return nil, fmt.Errorf("failed to parse servers JSON: %v", err)
	}

	return servers, nil
}
