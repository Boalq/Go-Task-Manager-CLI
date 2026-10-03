package taskmanipulation

import (
	"Task-Tracker/basicoperations"
	"Task-Tracker/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func Delete(args []string, content []byte) {
	if len(args) != 3 {
		fmt.Printf("Please use ./programm delete ID\n")
		os.Exit(2)
	}

	targetID, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatalf("Use a number(ID) to deleta a given task")
	}

	var all_tasks = []models.Task{}

	err = json.Unmarshal(content, &all_tasks)

	if err != nil {
		log.Fatal(err)
	}

	index, err := basicoperations.SearchByID(all_tasks, targetID)

	if err != nil {
		log.Fatal(err)
	}

	all_tasks = slices.Delete(all_tasks, index, index + 1)

	basicoperations.WriteToJson(all_tasks)

	fmt.Printf("Deleted.\n")
}
