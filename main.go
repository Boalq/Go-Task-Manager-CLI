package main

import (
	"fmt"
	"os"
)

const json_file = "tasks.json"

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

	// Learn basic JSON handling
	// Learn How to read and write to files
	// Look more on how to implement the list function!

	switch os.Args[1] {
	case "add":
		add(os.Args)
	case "update":
		update(os.Args)
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

func update(args []string) {
	fmt.Printf("You called the %s function", args[1])
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
