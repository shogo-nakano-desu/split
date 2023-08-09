package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func SplitByLines(file *os.File, lineCount int, baseFileName string) {
	scanner := bufio.NewScanner(file)
	var buffer strings.Builder

	fileIdx := 1
	lineIdx := 0

	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
		lineIdx++

		if lineIdx == lineCount {
			writeToFile(buffer.String(), baseFileName, fileIdx)
			buffer.Reset()
			lineIdx = 0
			fileIdx++
		}
	}

	if buffer.Len() > 0 {
		writeToFile(buffer.String(), baseFileName, fileIdx)
	}
}

func SplitByFileCounts(file *os.File, fileCount int, baseFileName string) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalSize := fileInfo.Size()
	averageSize := totalSize / int64(fileCount)

	buffer := make([]byte, averageSize)
	fileIdx := 1

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		writeToFile(string(buffer[:n]), baseFileName, fileIdx)
		fileIdx++
	}
}

func SplitByBytes(file *os.File, byteSize int, baseFileName string) {
	buffer := make([]byte, byteSize)
	fileIdx := 1

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		writeToFile(string(buffer[:n]), baseFileName, fileIdx)
		fileIdx++
	}
}
