package taskmanipulation

import (
	"Task-Tracker/basicoperations"
	"Task-Tracker/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Add(args []string, content []byte) {

	if len(args) != 3 {
		fmt.Printf("Please Use: ./programm add 'Task Description'\n")
		os.Exit(2)
	}

	var all_tasks []models.Task

	var IDs int

	//fmt.Println(content)

	//Clearing a empty Json
	if strings.Compare(string(content), "[]") == 0{content = []byte{}}


	//Doing stuff
	if strings.Compare(string(content), "") != 0{
		err := json.Unmarshal(content, &all_tasks)

		if err != nil {
			log.Fatal("This is not a json file to read from!")
		}

		IDs = all_tasks[len(all_tasks)-1].ID
	} else {
		IDs = 0
	}

	// fmt.Println(IDs)

	all_tasks = append(all_tasks, models.Task{
		ID:          IDs + 1,
		Description: args[2],
		Status:      models.StatusComponent{ToDo: true, InProgress: false, Done: false},
		CreatedAt:   time.Now().String(),
		UpdatedAt:   "",
	})

	basicoperations.WriteToJson(all_tasks)

	fmt.Printf("Task Added Successfully ID: %d", IDs+1)
}
