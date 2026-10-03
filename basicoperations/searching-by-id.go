package basicoperations

import (
	"Task-Tracker/models"
	"errors"
)

func SearchByID(all_tasks []models.Task, tID int) (int, error){
	for i, task := range all_tasks {
		if task.ID == tID {
			return i, nil
		}
	}
	return -1, errors.New("Task not found by ID")
}