package data

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

type Metrics struct {
	CPU  CPUstats  `json:"cpu"`
	RAM  RAMstats  `json:"ram"`
	DISK DISKstats `json:"disk"`
}

type Snapshot struct {
	Host      string    `json:"host"`
	Timestamp time.Time `json:"timestamp"`
	Metrics   Metrics   `json:"metrics"`
}

func BuildSnapshot(client *ssh.Client, ServerIP string) (Snapshot, error) {

	metrics, err := getMetrics(client)
	if err != nil {
		return Snapshot{}, fmt.Errorf("getting metrics failed: %w", err)
	}

	return Snapshot{
		Host:      ServerIP,
		Timestamp: time.Now(),
		Metrics:   metrics,
	}, nil
}

func getMetrics(client *ssh.Client) (Metrics, error) {
	cpu_stats, err := CpuUsage(client)
	if err != nil {
		fmt.Println("[-] CPU usage err: ", err)
	} else {
		fmt.Println("[+] CPU Scan Complete")
	}

	ram_stats, err := RamUsage(client)
	if err != nil {
		fmt.Println("[-] RAM usage err: ", err)
	} else {
		fmt.Println("[+] RAM Scan Complete")
	}

	disk_stats, err := DiskUsage(client)
	if err != nil {
		fmt.Println("[-] Disk usage err: ", err)
	} else {
		fmt.Println("[+] DISK Scan Complete")
	}

	return Metrics{
		CPU:  cpu_stats,
		RAM:  ram_stats,
		DISK: disk_stats,
	}, nil

}
