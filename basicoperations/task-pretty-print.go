package basicoperations

import (
	"Task-Tracker/models"
	"fmt"
)

func TaskPP(task models.Task) {
	fmt.Printf("ID: %d \t Description: %s \t\n",
		task.ID,
		task.Description,
	)
}
