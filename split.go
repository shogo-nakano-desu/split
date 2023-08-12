package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

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

// SplitByBytesMultithread is a function that splits a file by the number of bytes using goroutines.
func SplitByBytesMultithread(file *os.File, byteSize int, baseFileName string, suffixLen int) error {
	buffer := make([]byte, byteSize)
	strings, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}

	const maxGoroutines = 10
	var wg sync.WaitGroup
	var errorCh = make(chan error, 1)
	goroutineCh := make(chan struct{}, maxGoroutines)

	fileIdx := 0
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error: reading file: %v", err)
		}

		content := string(buffer[:n])
		suffix := strings[fileIdx]

		goroutineCh <- struct{}{}
		wg.Add(1)
		go func(content, suffix string) {
			defer wg.Done()
			defer func() { <-goroutineCh }()

			err := writeToFile(content, baseFileName, suffix)
			if err != nil {
				select {
				case errorCh <- err:
				default:
				}
			}
		}(content, suffix)

		fileIdx++
	}

	wg.Wait()

	select {
	case err := <-errorCh:
		return err
	default:
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
