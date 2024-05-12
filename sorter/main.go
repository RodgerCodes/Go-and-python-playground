package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// newDirectory := os.Chdir("./Downloads")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	downloadDir := homeDir + "/Downloads"

	files, err := os.ReadDir(downloadDir)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {

		fileExt := filepath.Ext(file.Name())
		fileDir := downloadDir + "/" + file.Name()

		if fileExt == ".mkv" || fileExt == ".mp4" {
			videoDir := homeDir + "/Videos/" + file.Name()
			moveFileToNewDir(fileDir, videoDir)
		} else if fileExt == ".png" || fileExt == ".jpeg" || fileExt == ".svg" || fileExt == ".jpg" {
			imageDir := homeDir + "/Pictures/" + file.Name()
			moveFileToNewDir(fileDir, imageDir)
		} else {
			documentsDir := homeDir + "/Documents/" + file.Name()
			moveFileToNewDir(fileDir, documentsDir)
		}
	}

	fmt.Println("Successfully organized downloads directory")

}

func moveFileToNewDir(oldPath string, newPath string) error {
	e := os.Rename(oldPath, newPath)
	if e != nil {
		return e
	}

	return nil
}
