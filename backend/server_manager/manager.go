package servermanager

import (
	"fmt"
	server "mvp/backend/server_coms"
	data "mvp/backend/server_data"
	"path/filepath"
	"time"
)

type Server struct {
	Name       string `json:"name"`
	IP         string `json:"ip"`
	PemKeyPath string `json:"pemkeypath"`
}

func PerServerTask(srv Server, outputDir string) {

	dataFilePath := filepath.Join(outputDir, srv.Name+".jsonl")

	sshClient, err := server.ConnectToServer(srv.PemKeyPath, srv.IP)
	if err != nil {
		fmt.Println("SSH connection failed:", err)
		return
	}

	if sshClient != nil {
		fmt.Println("Connected to server")
	}

	go data.AutoSnapshots(sshClient, srv.IP, dataFilePath, 1*time.Minute)

}
