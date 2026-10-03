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
		log.Fatalf("Make Sure that you have a tasks.json in the root directory \nCouldn't open the file %e", err)
	}

	defer file.Close()

	//Getting the File Content
	content, err := os.ReadFile(models.Json_file)
	if err != nil {
		log.Fatal("Couldn't Read the file!")
	}

	// Learn basic JSON handling
	// Learn How to read and write to files
	// Look more on how to implement the list function!

	switch os.Args[1] {
	case "add":
		taskmanipulation.Add(os.Args, content)
	case "update":
		taskmanipulation.Update(os.Args, content)
	case "delete":
		delete(os.Args)
	case "mark-in-progress":
		markInProgress(os.Args)
	case "mark-done":
		markDone(os.Args)
	case "list":
		list(os.Args)
	}
}

func delete(args []string) {
	fmt.Printf("You called the %s function", args[1])
}

func markInProgress(args []string) {
	fmt.Printf("You called the %s function", args[1])
}

func markDone(args []string) {
	fmt.Printf("You called the %s function", args[1])
}

func list(args []string) {
	fmt.Printf("You called the %s function", args[1])
}
