package handlers

import (
	"io/ioutil"
	"log"
)

// this function open a .json file and return the content into Object
func ReadJsonFile(file_name string) []byte {
	file, err := ioutil.ReadFile("./seeds/" + file_name + ".json")
	if err != nil {
		log.Fatal("Cant read file", err)
	}
	return file
}
