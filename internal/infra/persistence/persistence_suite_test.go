package persistence_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/sakaguchi-0725/task-tracker/internal/util"
)

func TestMain(m *testing.M) {
	err := util.CreateEmptyJSON("test_task.json")
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()

	err = util.DeleteJSONFile("test_task.json")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(code)
}

func clearJSON() {
	var emptyData []map[string]interface{}

	data, err := json.MarshalIndent(emptyData, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal empty JSON: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("test_task.json", data, 0644)
	if err != nil {
		fmt.Printf("failed to clear JSON file: %v\n", err)
		os.Exit(1)
	}
}
