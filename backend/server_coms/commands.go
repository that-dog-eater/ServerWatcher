package server

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func ExecuteCommand(client *ssh.Client, command string) (string, error) {
	// Create a session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	// Run the command
	err = session.Run(command)
	if err != nil {
		return "", fmt.Errorf("command failed: %s, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
