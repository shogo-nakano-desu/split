package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
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

// SplitByLinesMultithread is a function that splits a file by the number of lines using goroutines.
func SplitByLinesMultithread(file *os.File, lineCount int, baseFileName string, suffixLen int) error {
	scanner := bufio.NewScanner(file)
	chunks := make([][]string, 0)

	var currentChunk []string
	for scanner.Scan() {
		currentChunk = append(currentChunk, scanner.Text())
		if len(currentChunk) == lineCount {
			chunks = append(chunks, currentChunk)
			currentChunk = []string{}
		}
	}
	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}

	strs, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}

	const maxGoroutines = 10
	sem := make(chan struct{}, maxGoroutines)
	errChan := make(chan error, len(chunks))
	var wg sync.WaitGroup

	for i, chunk := range chunks {
		sem <- struct{}{}
		wg.Add(1)

		go func(idx int, lines []string) {
			defer wg.Done()
			content := strings.Join(lines, "\n")
			err := writeToFile(content, baseFileName, strs[idx])
			if err != nil {
				errChan <- err
			}
			<-sem
		}(i, chunk)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("encountered %d errors, first error: %v", len(errors), errors[0])
	}

	return nil
}

// SplitByFileCounts is a function that splits a file to the number of files.
func SplitByFileCounts(file *os.File, fileCount int, baseFileName string, suffixLen int) error {
	totalLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalLines++
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}

	linesPerFile := (totalLines + fileCount - 1) / fileCount
	strs, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}

	lineIdx := 0
	fileIdx := 0
	buffer := strings.Builder{}

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
		lineIdx++

		if lineIdx == linesPerFile {
			err := writeToFile(buffer.String(), baseFileName, strs[fileIdx])
			if err != nil {
				return err
			}
			buffer.Reset()
			lineIdx = 0
			fileIdx++
		}
	}

	if buffer.Len() > 0 {
		err := writeToFile(buffer.String(), baseFileName, strs[fileIdx])
		if err != nil {
			return err
		}
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
