package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func createTempFile(content string) *os.File {
	tmpfile, _ := os.CreateTemp("", "example_tmp")
	_, _ = tmpfile.WriteString(content)
	_, _ = tmpfile.Seek(0, 0) // reset the offset for reading
	return tmpfile
}

func removeFilesWithPattern(pattern string) {
	matches, _ := filepath.Glob(pattern)
	for _, match := range matches {
		_ = os.Remove(match)
	}
}

func fileNamesWithPattern(pattern string) ([]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func TestSplitByLines(t *testing.T) {
	tmpfile := createTempFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern("testing_file*")
	}()

	_ = SplitByLines(tmpfile, 2, "testing_file", 2)

	res, _ := fileNamesWithPattern("testing_file*")
	expected := []string{"testing_fileaa", "testing_fileab", "testing_fileac"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestSplitByFileCounts(t *testing.T) {
	tmpfile := createTempFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern("testing_file*")
	}()

	_ = SplitByFileCounts(tmpfile, 2, "testing_file", 2)

	res, _ := fileNamesWithPattern("testing_file*")
	expected := []string{"testing_fileaa", "testing_fileab"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestSplitByBytes(t *testing.T) {
	tmpfile := createTempFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern("testing_file*")
	}()

	_ = SplitByBytes(tmpfile, 2, "testing_file", 2)

	res, _ := fileNamesWithPattern("testing_file*")
	resLen := len(res)
	expected := 39
	if !reflect.DeepEqual(resLen, expected) {
		t.Errorf("expected %v, got %v", expected, resLen)
	}
}
