package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func add(args []string) {

	file, err := os.Open(json_file)

	if err != nil {
		log.Fatalf("Couldn't open the file %e", err)
	}

	defer file.Close()

	content, err := os.ReadFile(json_file)
	if err != nil {
		log.Fatal("Couldn't Read the file!")
	}

	var all_tasks []Task

	var IDs int

	//fmt.Println(content)

	if strings.Compare(string(content), "") != 0 {
		err = json.Unmarshal(content, &all_tasks)

		if err != nil {
			log.Fatal("This is not a json file to read from!")
		}

		IDs = all_tasks[len(all_tasks) - 1].ID
	} else {
		IDs = 0
	}

	// fmt.Println(IDs)

	all_tasks = append(all_tasks, Task{ID: IDs+1, Todo: args[2], InProgress: false, Done: false})


	var bufferiono bytes.Buffer

	ss, err := json.Marshal(all_tasks)
	
	if err != nil{
		log.Fatal(err)
	}

	err = json.Indent(&bufferiono, ss, "", "\t")

	if err != nil{
		log.Fatal(err)
	}


	err = os.WriteFile(json_file, bufferiono.Bytes(), 0x2)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Task Added Successfully ID: %d", IDs+1)
}
