package taskmanipulation

import (
	"Task-Tracker/basicoperations"
	"Task-Tracker/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Update(args []string, content []byte) {

	if len(args) != 4 {
		fmt.Printf("Please use ./programm update ID 'New Description'\n")
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

	if err != nil {
		log.Fatal(err)
	}

	// Changing the Description
	all_tasks[index].Description = os.Args[3]
	all_tasks[index].UpdatedAt = time.Now().String()

	// Writing to a Json File
	basicoperations.WriteToJson(all_tasks)

	fmt.Printf("Sucessfully changed to %q", os.Args[3])
}
