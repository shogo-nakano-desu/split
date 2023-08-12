package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// SplitByLines is a function that splits a file by the number of lines.
func SplitByLines(file *os.File, lineCount int, baseFileName string, suffixLen int) error {
	scanner := bufio.NewScanner(file)
	var buffer strings.Builder

	strings, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}
	fileIdx := 0
	lineIdx := 0

	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
		lineIdx++

		if lineIdx == lineCount {
			err := writeToFile(buffer.String(), baseFileName, strings[fileIdx])
			if err != nil {
				return err
			}
			buffer.Reset()
			lineIdx = 0
			fileIdx++
		}
	}

	if buffer.Len() > 0 {
		err := writeToFile(buffer.String(), baseFileName, strings[fileIdx])
		if err != nil {
			return err
		}
	}
	return nil
}

// SplitByFileCounts is a function that splits a file to the number of files.
func SplitByFileCounts(file *os.File, fileCount int, baseFileName string, suffixLen int) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	totalSize := fileInfo.Size()
	averageSize := totalSize / int64(fileCount)

	buffer := make([]byte, averageSize)
	strings, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}
	fileIdx := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error: reading file: %v", err)
		}

		err = writeToFile(string(buffer[:n]), baseFileName, strings[fileIdx])
		if err != nil {
			return err
		}
		fileIdx++
	}
	return nil
}

// SplitByBytes is a function that splits a file by the number of bytes.
func SplitByBytes(file *os.File, byteSize int, baseFileName string, suffixLen int) error {
	buffer := make([]byte, byteSize)
	strings, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}
	fileIdx := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error: reading file: %v", err)
		}

		err = writeToFile(string(buffer[:n]), baseFileName, strings[fileIdx])
		if err != nil {
			return err
		}
		fileIdx++
	}
	return nil
}

// writeToFile is a function that writes the given content to the file.
func writeToFile(content string, baseFileName string, suffix string) error {
	if baseFileName == "" {
		baseFileName = "x"
	}
	newFileName := fmt.Sprintf("%s%s", baseFileName, suffix)
	outFile, err := os.Create(newFileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	defer func() {
		closeErr := outFile.Close()
		if closeErr != nil && err == nil {
			err = fmt.Errorf("error closing the file: %v", closeErr)
		}
	}()

	_, err = outFile.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to the file: %v", err)
	}
	return nil
}
