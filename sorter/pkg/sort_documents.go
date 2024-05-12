package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func SortDocuments() error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	documentsDir := homeDir + "/Documents"

	files, err := os.ReadDir(documentsDir)

	if err != nil {
		return err
	}

	for _, file := range files {

		fileExt := filepath.Ext(file.Name())
		fileDir := documentsDir + "/" + file.Name()

		switch fileExt {
		case ".mkv", ".mp4":
			videoDir := homeDir + "/Videos/" + file.Name()
			moveFileToNewDir(fileDir, videoDir)
		case ".png", ".jpeg", ".svg", ".jpg":
			imageDir := homeDir + "/Pictures/" + file.Name()
			moveFileToNewDir(fileDir, imageDir)
		default:
			fmt.Println("Already a document")
			// Do nothing
		}
	}

	return nil
}
