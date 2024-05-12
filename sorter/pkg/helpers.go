package pkg

import "os"

func moveFileToNewDir(oldPath string, newPath string) error {
	e := os.Rename(oldPath, newPath)
	if e != nil {
		return e
	}

	return nil
}
