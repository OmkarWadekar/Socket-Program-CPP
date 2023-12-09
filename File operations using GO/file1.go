package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// File path for reading and writing
	filePath := "example.txt"

	// Writing to a file
	contentToWrite := []byte("Hello, this is a sample content.")
	err := ioutil.WriteFile(filePath, contentToWrite, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully.")

	// Reading from a file
	contentRead, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("File content:", string(contentRead))
}
