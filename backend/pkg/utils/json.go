package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonFile(filename string, data any) error {
	jsonFile, err := os.ReadFile(fmt.Sprintf("./data/%s.json", filename))
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonFile, &data)
}

func WriteJsonFile(filename string, data any) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("./data/%s.json", filename), file, 0644)
}
