// used ram percent
// total ram
// avalible ram

package data

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type RAMstats struct {
	Active_Ram_Percent float64 `json:"active_ram_percent"`
	Total_Ram          int     `json:"total_ram"`
	Used_Ram           int     `json:"avalible_ram"`
}

func RamUsage(client *ssh.Client) (RAMstats, error) {

	totalRAM, err := GetTotalRam(client)
	if err != nil {
		fmt.Println("Error getting total Ram: ", err)
		return RAMstats{}, nil
	}

	usedRAM, err := GetUsedRam(client)
	if err != nil {
		fmt.Println("Error getting used Ram ", err)
		return RAMstats{}, nil
	}

	usedPercent, err := ActiveRamPercent(usedRAM, totalRAM)
	if err != nil {
		fmt.Println("Used Percent Calc error: ", err)
		return RAMstats{}, nil
	}

	return RAMstats{
		Active_Ram_Percent: usedPercent,
		Total_Ram:          totalRAM,
		Used_Ram:           usedRAM,
	}, nil
}
