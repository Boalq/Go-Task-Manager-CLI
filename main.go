package main

import (
	"Task-Tracker/models"
	"Task-Tracker/taskmanipulation"
	"fmt"
	"log"
	"os"
)

type Task struct {
	ID         int    `json:"ID"`
	Todo       string `json:"ToDo"`
	InProgress bool   `json:"InProgress"`
	Done       bool   `json:"Done"`
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Invalid Amount of Arguments!")
		return
	}

	// Handling the tasks.json
	file, err := os.Open(models.Json_file)

	if err != nil {
		file, err = os.Create("tasks.json")

		if err != nil{
			log.Fatalln("Couldn't either find or create the File 'tasks.json'")
		}
	}

	defer file.Close()

	//Getting the File Content
	content, err := os.ReadFile(models.Json_file)
	if err != nil {
		log.Fatal("Couldn't Read the file!")
	}

	switch os.Args[1] {
	case "add":
		taskmanipulation.Add(os.Args, content)
	case "update":
		taskmanipulation.Update(os.Args, content)
	case "delete":
		taskmanipulation.Delete(os.Args, content)
	case "mark-in-progress":
		taskmanipulation.ChangeProgress(os.Args, content, 1)
	case "mark-done":
		taskmanipulation.ChangeProgress(os.Args, content, 2)
	case "list":
		taskmanipulation.List(os.Args, content)
	}
}
