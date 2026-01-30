package output

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSON(path string, v any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func AppendJSON(path string, v any) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	return enc.Encode(v) // writes one JSON object per line
}

func CheckIfFileExits(file_path string) {
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		file, err := os.Create(file_path)
		if err != nil {
			fmt.Println("Failed to create data file:", err)
			return
		}
		file.Close()
	}
}
