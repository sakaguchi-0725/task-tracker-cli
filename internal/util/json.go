package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func IsJsonExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func CreateEmptyJSON(filePath string) error {
	var emptyData []map[string]interface{}

	data, err := json.MarshalIndent(emptyData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal empty JSON: %v", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write empty JSON to file: %v", err)
	}

	return nil
}

func DeleteJSONFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete JSON file: %v", err)
	}
	return nil
}

func ReadJSON[T any](filePath string) ([]T, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []T{}, nil
		}
		return nil, fmt.Errorf("failed to read file (%s): %v", filePath, err)
	}

	var vals []T
	if len(data) > 0 {
		if err := json.Unmarshal(data, &vals); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
	}

	return vals, nil
}

func WriteJSON[T any](filePath string, vals []T) error {
	results, err := json.MarshalIndent(vals, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(filePath, results, 0644); err != nil {
		return fmt.Errorf("failed to write file (%s): %v", filePath, err)
	}

	return nil
}
