// package main
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
		os.Exit(1)
	}

	defer func() {
		err := outFile.Close()
		if err != nil {
			fmt.Printf("Error closing the file: %v\n", err)
			os.Exit(1)
		}
	}()

	_, err = outFile.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to the file: %v\n", err)
		os.Exit(1)
	}

}
