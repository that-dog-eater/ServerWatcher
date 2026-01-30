package api

import (
	"encoding/json"
	"fmt"
	servermanager "mvp/backend/server_manager"
	"net"
	"os"
)

func (a *API) AddServer(name, ip, pemKey string) error {

	err := validateInput(name, ip, pemKey, a.ServersFilePath)
	if err != nil {
		return err
	}

	newServer := servermanager.Server{
		Name:       name,
		IP:         ip,
		PemKeyPath: pemKey,
	}

	err = addServerAndStart(newServer, a.ServersFilePath, a.OutputDir)
	if err != nil {
		fmt.Println("Adding Server Error: ", err)
	}

	fmt.Println("Server added:", name)
	return nil
}

func addServerAndStart(
	newServer servermanager.Server,
	serversFilePath string,
	outputDir string,
) error {

	err := servermanager.AddServer(serversFilePath, newServer)
	if err != nil {
		fmt.Println("Error Adding server: ", err)
		return err
	}

	go servermanager.PerServerTask(newServer, outputDir)

	return nil
}

func validateInput(name, ip, pemKey string, serversFilePath string) error {
	// Required fields
	if name == "" || ip == "" || pemKey == "" {
		return fmt.Errorf("all fields are required")
	}

	// Validate IP address (IPv4 or IPv6)
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("invalid IP address format")
	}

	// Check PEM file exists
	info, err := os.Stat(pemKey)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("PEM key file does not exist")
		}
		return fmt.Errorf("unable to access PEM key file: %v", err)
	}

	// Ensure it's a file, not a directory
	if info.IsDir() {
		return fmt.Errorf("PEM key path must be a file, not a directory")
	}

	data, err := os.ReadFile(serversFilePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read servers file: %v", err)
	}

	if len(data) > 0 {
		var servers []servermanager.Server
		if err := json.Unmarshal(data, &servers); err != nil {
			return fmt.Errorf("failed to parse servers file: %v", err)
		}

		for _, s := range servers {
			if s.Name == name {
				return fmt.Errorf("server name already exists")
			}
			if s.IP == ip {
				return fmt.Errorf("server with this IP already exists")
			}
		}

	}
	return nil

}
