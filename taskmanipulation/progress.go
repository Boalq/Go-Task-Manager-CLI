package taskmanipulation

import (
	"Task-Tracker/basicoperations"
	"Task-Tracker/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ChangeProgress(args []string, content []byte, status int){
	if len(args) != 3 {
		fmt.Printf("Please use ./programm \n")
		os.Exit(2)
	}

	targetID, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatalf("Use a number(ID) to update a given task")
	}

	var all_tasks = []models.Task{}


	err = json.Unmarshal(content, &all_tasks)

	if err != nil {
		log.Fatal(err)
	}

	//Checking for ID
	index, err := basicoperations.SearchByID(all_tasks, targetID)

	if err != nil{
		log.Fatalf("Found an error when searching for the ID %d \n %e", targetID, err)
	}

	switch(status){
	case 1:
		all_tasks[index].Status.InProgress = true
		all_tasks[index].Status.ToDo = false
		all_tasks[index].Status.Done = false
	case 2:
		all_tasks[index].Status.InProgress = false
		all_tasks[index].Status.ToDo = false
		all_tasks[index].Status.Done = true
	}

	basicoperations.WriteToJson(all_tasks)
}