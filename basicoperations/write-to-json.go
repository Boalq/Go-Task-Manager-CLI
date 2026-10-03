package basicoperations

import (
	"Task-Tracker/models"
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func WriteToJson(all_tasks []models.Task) {
	var bufferiono bytes.Buffer

	ss, err := json.Marshal(all_tasks)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Indent(&bufferiono, ss, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(models.Json_file, bufferiono.Bytes(), 0x2)

	if err != nil {
		log.Fatal(err)
	}
}
