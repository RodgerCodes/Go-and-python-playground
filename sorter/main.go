package main

import (
	"fmt"
	"os"
)

func main() {

	// newDirectory := os.Chdir("./Downloads")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	newHomeDir := homeDir + "/Downloads"

	files, err := os.ReadDir(newHomeDir)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
