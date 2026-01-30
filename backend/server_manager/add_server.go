package servermanager

import (
	"encoding/json"
	"fmt"
	"mvp/backend/output"
	"os"
)

func AddServer(serversFilePath string, newServer Server) error {
	var servers []Server

	output.CheckIfFileExits(serversFilePath)

	data, err := os.ReadFile(serversFilePath)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &servers); err != nil {
			return err
		}
	}

	// 🔒 Uniqueness checks
	for _, s := range servers {
		if s.IP == newServer.IP {
			return fmt.Errorf("server with IP %s already exists", newServer.IP)
		}
		if s.Name == newServer.Name {
			return fmt.Errorf("server with name %s already exists", newServer.Name)
		}
	}

	servers = append(servers, newServer)

	return output.WriteJSON(serversFilePath, servers)
}
