package data

import (
	"fmt"
	"math"
	server "mvp/backend/server_coms"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

type CPUstats struct {
	Usage float64 `json:"usage"`
}

func CpuUsage(client *ssh.Client) (CPUstats, error) {

	bash := `top -bn1 | grep "Cpu(s)" | awk -F',' '{for(i=1;i<=NF;i++){if($i ~ /id/){print $i}}}' | awk '{print $1}'`

	output, err := server.ExecuteCommand(client, bash)
	if err != nil {
		fmt.Println("Server Output Error: ", err)
	}

	clean_value := strings.TrimSpace(output) // removes all leading/trailing spaces and newlines

	idle, err := strconv.ParseFloat(clean_value, 64)
	if err != nil {
		return CPUstats{}, fmt.Errorf("parsing CPU usage failed: %w", err)
	}

	cpu_usage_percent := 100 - idle
	cpu_usage_percent = math.Round(cpu_usage_percent*10) / 10

	return CPUstats{
		Usage: cpu_usage_percent,
	}, nil
}
