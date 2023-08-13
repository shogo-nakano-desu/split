package main

import (
	"bufio"
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, chunk := range chunks {
		wg.Add(1)

		go func(ctx context.Context, idx int, lines []string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			select {
			case <-ctx.Done():
				return
			default:
				content := strings.Join(lines, "\n")
				if len(strs) <= idx {
					errChan <- fmt.Errorf("error: too many files")
					cancel()
					return
				}
				err := writeToFile(content, baseFileName, strs[idx])
				if err != nil {
					errChan <- err
					cancel()
					return
				}
			}
		}(ctx, i, chunk)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// SplitByFileCountsMultithread is a function that splits a file to the number of files using goroutines.
func SplitByFileCountsMultithread(file *os.File, fileCount int, baseFileName string, suffixLen int) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	totalSize := fileInfo.Size()
	bytesPerChunk := totalSize / int64(fileCount)
	if bytesPerChunk < 1 {
		return fmt.Errorf("error: can't split into more than %v files", totalSize)
	}
	remainingBytes := totalSize % int64(fileCount)

	strs, err := GenerateStrings(suffixLen, "", 0)
	if err != nil {
		return err
	}

	errChan := make(chan error, fileCount)
	sem := make(chan struct{}, 10)

	for i := 0; i < fileCount; i++ {
		if len(strs) <= i {
			return fmt.Errorf("error: too many files")
		}
		var currentChunkSize int64
		if i == fileCount-1 {
			currentChunkSize = bytesPerChunk + remainingBytes
		} else {
			currentChunkSize = bytesPerChunk
		}

		buffer := make([]byte, currentChunkSize)
		_, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		go func(data []byte, filenameSuffix string) {
			sem <- struct{}{}
			err := writeToFile(string(data), baseFileName, filenameSuffix)
			errChan <- err
			<-sem
		}(buffer, strs[i])
	}

	for i := 0; i < fileCount; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	close(errChan)

	err = <-errChan
	if err != nil {
		return err
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
		if len(strings) <= fileIdx {
			return fmt.Errorf("error: too many files")
		}
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
