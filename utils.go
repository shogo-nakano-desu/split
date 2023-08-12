package main

import (
	"fmt"
	"strings"
)

// NormalizeArgs is a function that normalizes the arguments passed to the program.
// For example, if the user passes "-l10" instead of "-l 10", this function will
// normalize the arguments to "-l 10".
func NormalizeArgs(args []string) []string {
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-l") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-l", args[i][2:]}, args[i+1:]...)...)
		}
		if strings.HasPrefix(args[i], "-n") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-n", args[i][2:]}, args[i+1:]...)...)
		}
		if strings.HasPrefix(args[i], "-b") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-b", args[i][2:]}, args[i+1:]...)...)
		}
		if strings.HasPrefix(args[i], "-a") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-a", args[i][2:]}, args[i+1:]...)...)
		}
	}
	return args
}

// GenerateStrings is a function that generates strings from the given length.
func GenerateStrings(length int, prefix string, counter int) ([]string, error) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	if counter == 0 && length == 0 {
		return []string{}, fmt.Errorf("Error: suffix length must be greater than 0")
	}
	if length == 0 {
		return []string{prefix}, nil
	}
	if length > 5 {
		return []string{}, fmt.Errorf("Error: suffix length must be less than or equal to 5")
	}

	var result []string
	for _, char := range alphabet {
		counter++
		res, err := GenerateStrings(length-1, prefix+string(char), counter)
		if err != nil {
			return []string{}, err
		}
		result = append(result, res...)
	}

	return result, nil
}

// Args is a struct that represents the arguments passed to the program.
type Args struct {
	LineCount int
	FileCount int
	ByteSize  int
	Args      []string
}

// IllegalArgsChecker is a function that checks if the arguments passed to the program are valid.
func IllegalArgsChecker(params Args) error {
	lineSetCount := 0
	fileSetCount := 0
	byteSetCount := 0

	lineCount, fileCount, byteSize, args := params.LineCount, params.FileCount, params.ByteSize, params.Args

	for _, arg := range args {
		switch arg {
		case "-l":
			lineSetCount++
		case "-n":
			fileSetCount++
		case "-b":
			byteSetCount++
		case "-a":
			continue
		default:
			if strings.HasPrefix(arg, "-") {
				return fmt.Errorf("Error: unknown option %s", arg)
			}
		}
	}

	if lineSetCount+fileSetCount+byteSetCount > 1 {
		return fmt.Errorf(
			`usage: split [-l line_count] [-a suffix_length] [file [prefix]]
			split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]
			split -n chunk_count [-a suffix_length] [file [prefix]]
			split -p pattern [-a suffix_length] [file [prefix]]`,
		)
	}

	if lineCount <= 0 && lineSetCount == 1 {
		return fmt.Errorf("Error: %d: illegal line count\n", lineCount)
	}

	if fileCount <= 0 && fileSetCount == 1 {
		return fmt.Errorf("Error: %d: illegal file count\n", fileCount)
	}

	if byteSize <= 0 && byteSetCount == 1 {
		return fmt.Errorf("Error: %d: illegal byte size\n", byteSize)
	}
	return nil
}
