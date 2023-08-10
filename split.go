package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// SplitByLines is a function that splits a file by the number of lines.
func SplitByLines(file *os.File, lineCount int, baseFileName string, suffixLen int) {
	scanner := bufio.NewScanner(file)
	var buffer strings.Builder

	strings := GenerateStrings(suffixLen, "", 0)
	fileIdx := 0
	lineIdx := 0

	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
		lineIdx++

		if lineIdx == lineCount {
			writeToFile(buffer.String(), baseFileName, strings[fileIdx])
			buffer.Reset()
			lineIdx = 0
			fileIdx++
		}
	}

	if buffer.Len() > 0 {
		writeToFile(buffer.String(), baseFileName, strings[fileIdx])
	}
}

// SplitByFileCounts is a function that splits a file to the number of files.
func SplitByFileCounts(file *os.File, fileCount int, baseFileName string, suffixLen int) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalSize := fileInfo.Size()
	averageSize := totalSize / int64(fileCount)

	buffer := make([]byte, averageSize)
	strings := GenerateStrings(suffixLen, "", 0)
	fileIdx := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		writeToFile(string(buffer[:n]), baseFileName, strings[fileIdx])
		fileIdx++
	}
}

// SplitByBytes is a function that splits a file by the number of bytes.
func SplitByBytes(file *os.File, byteSize int, baseFileName string, suffixLen int) {
	buffer := make([]byte, byteSize)
	strings := GenerateStrings(suffixLen, "", 0)
	fileIdx := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		writeToFile(string(buffer[:n]), baseFileName, strings[fileIdx])
		fileIdx++
	}
}

// writeToFile is a function that writes the given content to the file.
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