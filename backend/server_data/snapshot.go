package data

import (
	"fmt"
	"mvp/backend/output"
	"time"

	"golang.org/x/crypto/ssh"
)

func AutoSnapshots(client *ssh.Client, serverIP string, file_path string, interval time.Duration) {

	output.CheckIfFileExits(file_path)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		takeSnapshot(client, serverIP, file_path)
		<-ticker.C // wait for the next tick
	}
}

func takeSnapshot(client *ssh.Client, serverIP string, file_path string) {
	Snapshot, err := BuildSnapshot(client, serverIP)
	if err != nil {
		fmt.Println("CPU usage err: ", err)
	}
	err = output.AppendJSON(file_path, Snapshot)
	if err != nil {
		fmt.Println("WriteJson failed: ", err)
	}
}
