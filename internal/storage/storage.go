package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kuipid01/task-cli/internal/models"
)

const filePath = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []models.Task{}, nil
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}
