package data

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type DISKstats struct {
	TotalMB     int     `json:"total_mb"`
	UsedMB      int     `json:"used_mb"`
	FreeMB      int     `json:"free_mb"`
	UsedPercent float64 `json:"used_percent"` // one decimal place
}

func DiskUsage(client *ssh.Client) (DISKstats, error) {

	mount := "/"

	total_mb, err := GetTotalMB(client, mount)
	if err != nil {
		fmt.Println("Getting total mb failed: ", err)
	}

	used_mb, err := GetUsedMB(client, mount)
	if err != nil {
		fmt.Println("Getting total mb failed: ", err)
	}

	free_mb, err := GetFreeMB(client, mount)
	if err != nil {
		fmt.Println("Getting total mb failed: ", err)
	}

	used_percent, err := GetUsedPercent(client, mount)
	if err != nil {
		fmt.Println("Getting total mb failed: ", err)
	}

	return DISKstats{
		TotalMB:     total_mb,
		UsedMB:      used_mb,
		FreeMB:      free_mb,
		UsedPercent: used_percent,
	}, nil
}
