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

	if(len(args) != 4){
		fmt.Printf("Please use ./programm ID 'New Description'\n")
		os.Exit(2)
	}

	targetID, _ := strconv.Atoi(args[2])

	var all_tasks = []models.Task{}

	err := json.Unmarshal(content, &all_tasks)

	if err != nil {
		log.Fatal(err)
	}

	//Checking for ID
	found := false
	for _, task := range all_tasks {
		if task.ID == targetID {
			found = true
			break
		}
	}
	if !found {log.Fatalf("ID not Found")}
	
	// Changing the Description
	all_tasks[targetID - 1].Description = os.Args[3]
	all_tasks[targetID - 1].UpdatedAt = time.Now().String()	

	// Writing to a Json File
	basicoperations.WriteToJson(all_tasks)

	fmt.Println("Sucessfully changed")
}
