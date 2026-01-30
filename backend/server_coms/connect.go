package server

import (
	"errors"
	"os"

	"golang.org/x/crypto/ssh"
)

func ConnectToServer(PrivKeyPath string, ServerIP string) (*ssh.Client, error) {

	if !check_key(PrivKeyPath) {
		return nil, errors.New("key is not formatted correctly")
	}

	key, err := os.ReadFile(PrivKeyPath)
	if err != nil {
		return nil, err
	}

	parsed_key, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(parsed_key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", ServerIP+":22", config)
	if err != nil {
		return nil, err
	}

	return client, nil

}
