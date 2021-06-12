package lib

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func LogToFile(text, filename string, errorInstance error) {
	CreateFolder(LOGS_FOLDER_PATH)

	dirPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting dir path: %v", err)
	}
	f, err := os.OpenFile(path.Join(dirPath, LOGS_FOLDER_PATH, filename), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	currentTime := time.Now().String()
	log.Println(fmt.Sprintf("%s    %s. Error: %v", currentTime, text, errorInstance))
	f.Close()
	log.SetOutput(os.Stdout)
}
