package data

import (
	"fmt"
	"math"
	server "mvp/backend/server_coms"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GetTotalMB(client *ssh.Client, mount string) (int, error) {
	cmd := fmt.Sprintf("df -m %s | awk 'NR==2 {print $2}'", mount)
	output, err := server.ExecuteCommand(client, cmd)
	if err != nil {
		return 0, fmt.Errorf("error executing total disk command: %w", err)
	}

	clean := strings.TrimSpace(output)
	totalMB, err := strconv.Atoi(clean)
	if err != nil {
		return 0, fmt.Errorf("failed to parse total disk MB: %w", err)
	}

	return totalMB, nil
}

func GetUsedMB(client *ssh.Client, mount string) (int, error) {
	cmd := fmt.Sprintf("df -m %s | awk 'NR==2 {print $3}'", mount)
	output, err := server.ExecuteCommand(client, cmd)
	if err != nil {
		return 0, fmt.Errorf("error executing used disk command: %w", err)
	}

	clean := strings.TrimSpace(output)
	usedMB, err := strconv.Atoi(clean)
	if err != nil {
		return 0, fmt.Errorf("failed to parse used disk MB: %w", err)
	}

	return usedMB, nil
}

func GetFreeMB(client *ssh.Client, mount string) (int, error) {
	cmd := fmt.Sprintf("df -m %s | awk 'NR==2 {print $4}'", mount)
	output, err := server.ExecuteCommand(client, cmd)
	if err != nil {
		return 0, fmt.Errorf("error executing free disk command: %w", err)
	}

	clean := strings.TrimSpace(output)
	freeMB, err := strconv.Atoi(clean)
	if err != nil {
		return 0, fmt.Errorf("failed to parse free disk MB: %w", err)
	}

	return freeMB, nil
}

func GetUsedPercent(client *ssh.Client, mount string) (float64, error) {
	cmd := fmt.Sprintf("df -m %s | awk 'NR==2 {gsub(\"%%\",\"\",$5); print $5}'", mount)
	output, err := server.ExecuteCommand(client, cmd)
	if err != nil {
		return 0, fmt.Errorf("error executing used percent command: %w", err)
	}

	clean := strings.TrimSpace(output)
	usedPercent, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse used percent: %w", err)
	}

	// round to 1 decimal place if you like
	usedPercent = math.Round(usedPercent*10) / 10

	return usedPercent, nil
}
