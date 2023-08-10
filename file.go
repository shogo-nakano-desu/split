package main

import (
	"fmt"
	"os"
)

func writeToFile(content string, baseFileName string, suffix string) {
	if baseFileName == "" {
		baseFileName = "x"
	}
	newFileName := fmt.Sprintf("%s%s", baseFileName, suffix)
	outFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	outFile.WriteString(content)
}
