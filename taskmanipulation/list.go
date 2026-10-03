package taskmanipulation

import (
	"Task-Tracker/basicoperations"
	"Task-Tracker/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func List(args []string, content []byte) {

	if len(args) < 2 || len(args) > 3 {
		fmt.Printf("Please use ./programm list (done/in-progress)\n")
		os.Exit(2)
	}

	var all_tasks = []models.Task{}

	err := json.Unmarshal(content, &all_tasks)

	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 2 {
		for _, i := range all_tasks {
			basicoperations.TaskPP(i)
		}
		return
	}

	switch args[2] {
	case "done":
		for _, i := range all_tasks {
			if i.Status.Done == true {
				basicoperations.TaskPP(i)
			}
		}
	case "todo":
		for _, i := range all_tasks {
			if i.Status.ToDo == true {
				basicoperations.TaskPP(i)
			}
		}
	case "in-progress":
		for _, i := range all_tasks {
			if i.Status.InProgress == true {
				basicoperations.TaskPP(i)
			}
		}
	}

}
