package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	inputHtml := "test.html"
	outputPdF := "output.pdf"

	pythonScript := "generate.py"

	cmd := exec.Command("python", pythonScript, inputHtml, outputPdF)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing Python script: %s\n", err)
	} else {
		fmt.Println("Python script executed successfully")
	}

}
