package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetLatestServerMetrics reads the last line of P{serverName}.jsonl and returns it as a map
func (a *API) GetLatestServerMetrics(serverName string) (map[string]any, error) {
	if serverName == "" {
		return nil, fmt.Errorf("server name is required")
	}

	// Build file path: OutputDir/P{serverName}.jsonl
	filename := fmt.Sprintf("%s.jsonl", serverName)
	filePath := filepath.Join(a.OutputDir, filename)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("metrics file not found for server %s", serverName)
		}
		return nil, fmt.Errorf("failed to open metrics file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastLine string
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lastLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading metrics file: %w", err)
	}

	if lastLine == "" {
		return nil, fmt.Errorf("metrics file is empty for server %s", serverName)
	}

	var result map[string]any
	if err := json.Unmarshal([]byte(lastLine), &result); err != nil {
		return nil, fmt.Errorf("failed to parse metrics JSON: %w", err)
	}

	return result, nil
}
