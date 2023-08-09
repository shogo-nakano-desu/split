package main

import (
	"fmt"
	"os"
)

func writeToFile(content string, baseFileName string, index int) {
	newFileName := fmt.Sprintf("%s_%d", baseFileName, index)
	outFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	outFile.WriteString(content)
}
