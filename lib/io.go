package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func WriteFile(reader io.Reader, title *string) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		LogToFile("Can't read the file", LOGGER_FILE_PATH, err)
		log.Fatal(fmt.Sprintf("Error during reading the file: %v", err))
	}
	err = ioutil.WriteFile(*title, body, 0644)
	if err != nil {
		LogToFile("Can't write to file", LOGGER_FILE_PATH, err)
		log.Fatal(fmt.Sprintf("Error during writing to file: %v", err))
	}
}

func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			LogToFile("Can't create a folder", LOGGER_FILE_PATH, err)
			log.Fatal(fmt.Sprintf("Error during creating a folder: %v", err))
		}
	}
}
