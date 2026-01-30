package data

import (
	"fmt"
	"math"
	server "mvp/backend/server_coms"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GetTotalRam(client *ssh.Client) (int, error) {

	total_ram_bash := "free -m | awk '/^Mem:/ {print $2}'"

	raw_total_ram, err := server.ExecuteCommand(client, total_ram_bash)
	if err != nil {
		fmt.Println("Err executing command: ", err)
		return 0, nil
	}

	clean_total_ram := strings.TrimSpace(raw_total_ram)

	totalRAM, err := strconv.Atoi(clean_total_ram)
	if err != nil {
		return 0, fmt.Errorf("failed to parse total RAM: %w", err)
	}

	return totalRAM, nil
}

func GetUsedRam(client *ssh.Client) (int, error) {

	used_ram_bash := "free -m | awk '/^Mem:/ {print $3}'"

	output, err := server.ExecuteCommand(client, used_ram_bash)
	if err != nil {
		fmt.Println("Err executing command: ", err)
		return 0, nil
	}

	clean_used_ram := strings.TrimSpace(output)

	usedRAM, err := strconv.Atoi(clean_used_ram)
	if err != nil {
		return 0, fmt.Errorf("failed to parse total RAM: %w", err)
	}

	return usedRAM, nil
}

func ActiveRamPercent(used_ram int, total_ram int) (float64, error) {
	percent := (float64(used_ram) / float64(total_ram)) * 100
	// Round to 1 decimal place
	percent = math.Round(percent*10) / 10
	return percent, nil
}
