package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

// Helper functions

const BIG_INT = 9223372036854775807

func createTmpFile(content string) *os.File {
	tmpfile, _ := os.CreateTemp("", "example_tmp")
	_, _ = tmpfile.WriteString(content)
	_, _ = tmpfile.Seek(0, 0)
	return tmpfile
}

func removeFilesWithPattern(pattern string) {
	// Need to wait for the file to be created.
	time.Sleep(200 * time.Millisecond)
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

func generateLargeText(size int) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString("abcdefghijklmnopqrstuvwxyz\n")
	}
	return buffer.String()
}

func TestSplitByLinesMultithread(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	_ = SplitByLinesMultithread(tmpfile, 2, baseFileName.String(), 2)

	res, _ := fileNamesWithPattern(baseFileName.String() + "*")
	expected := []string{baseFileName.String() + "aa", baseFileName.String() + "ab", baseFileName.String() + "ac"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestSplitByLinesMultithreadTooLargeFile(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line
		seventh line
		eighth line
		ninth line
		tenth line
		eleventh line
		twelfth line
		thirteenth line
		fourteenth line
		fifteenth line
		sixteenth line
		seventeenth line
		eighteenth line
		nineteenth line
		twentieth line
		twenty-first line
		twenty-second line
		twenty-third line
		twenty-fourth line
		twenty-fifth line
		twenty-sixth line
		twenty-seventh line
		`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))
	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	err := SplitByLinesMultithread(tmpfile, 1, baseFileName.String(), 1)

	expected := fmt.Errorf("error: too many files")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestSplitByFileCounts(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	_ = SplitByFileCounts(tmpfile, 2, baseFileName.String(), 2)

	res, _ := fileNamesWithPattern(baseFileName.String() + "*")
	expected := []string{baseFileName.String() + "aa", baseFileName.String() + "ab"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestSplitByFileCountsTooLargeFile(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line
		seventh line
		eighth line
		ninth line
		tenth line
		eleventh line
		twelfth line
		thirteenth line
		fourteenth line
		fifteenth line
		sixteenth line
		seventeenth line
		eighteenth line
		nineteenth line
		twentieth line
		twenty-first line
		twenty-second line
		twenty-third line
		twenty-fourth line
		twenty-fifth line
		twenty-sixth line
		twenty-seventh line
		`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	err := SplitByFileCounts(tmpfile, 27, baseFileName.String(), 1)

	expected := fmt.Errorf("error: too many files")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestSplitByFileCountsIntoTooManyFiles(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	err := SplitByFileCounts(tmpfile, 1000, baseFileName.String(), 2)

	expected := fmt.Errorf("error: can't split into more than 77 files")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestSplitByFileCountsMultithread(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)
	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	_ = SplitByFileCountsMultithread(tmpfile, 2, baseFileName.String(), 2)

	res, _ := fileNamesWithPattern(baseFileName.String() + "*")
	expected := []string{baseFileName.String() + "aa", baseFileName.String() + "ab"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestSplitByFileCountsMultithreadTooLargeFile(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line
		seventh line
		eighth line
		ninth line
		tenth line
		eleventh line
		twelfth line
		thirteenth line
		fourteenth line
		fifteenth line
		sixteenth line
		seventeenth line
		eighteenth line
		nineteenth line
		twentieth line
		twenty-first line
		twenty-second line
		twenty-third line
		twenty-fourth line
		twenty-fifth line
		twenty-sixth line
		twenty-seventh line
		`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))
	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	err := SplitByFileCountsMultithread(tmpfile, 27, baseFileName.String(), 1)

	expected := fmt.Errorf("error: too many files")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestSplitByFileCountsMultithreadIntoTooManyFiles(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	err := SplitByFileCountsMultithread(tmpfile, 1000, baseFileName.String(), 2)

	expected := fmt.Errorf("error: can't split into more than 77 files")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestSplitByBytes(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line`,
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	_ = SplitByBytes(tmpfile, 2, baseFileName.String(), 2)

	res, _ := fileNamesWithPattern(baseFileName.String() + "*")
	resLen := len(res)
	expected := 39
	if !reflect.DeepEqual(resLen, expected) {
		t.Errorf("expected %v, got %v", expected, resLen)
	}
}

func TestSplitByBytesTooLargeFile(t *testing.T) {
	tmpfile := createTmpFile(
		`first line
		second line
		third line
		fourth line
		fifth line
		sixth line
		seventh line
		eighth line
		ninth line
		tenth line
		eleventh line
		twelfth line
		thirteenth line
		fourteenth line
		fifteenth line
		sixteenth line
		seventeenth line
		eighteenth line
		nineteenth line
		twentieth line
		twenty-first line
		twenty-second line
		twenty-third line
		twenty-fourth line
		twenty-fifth line
		twenty-sixth line
		twenty-seventh line
		`,
	)
	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))
	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	err := SplitByBytes(tmpfile, 1, baseFileName.String(), 1)

	expected := fmt.Errorf("error: too many files")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func BenchmarkSplitByFileCounts(b *testing.B) {
	tmpfile := createTmpFile(
		generateLargeText(1000000),
	)

	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	b.ResetTimer() // タイマーのリセット（セットアップ時間を除外）

	for i := 0; i < b.N; i++ {
		_ = SplitByFileCounts(tmpfile, 100000, baseFileName.String(), 5)
	}
}

func BenchmarkSplitByFileCountsMultithread(b *testing.B) {
	tmpfile := createTmpFile(
		generateLargeText(1000000),
	)
	baseFileName, _ := rand.Int(rand.Reader, big.NewInt(BIG_INT))

	defer func() {
		_ = os.Remove(tmpfile.Name())
		removeFilesWithPattern(baseFileName.String() + "*")
	}()

	b.ResetTimer() // タイマーのリセット（セットアップ時間を除外）

	for i := 0; i < b.N; i++ {
		_ = SplitByFileCountsMultithread(tmpfile, 10000, baseFileName.String(), 5)
	}
}
