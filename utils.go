package main

import (
	"fmt"
	"os"
	"strings"
)

// NormalizeArgs is a function that normalizes the arguments passed to the program.
// For example, if the user passes "-l10" instead of "-l 10", this function will
// normalize the arguments to "-l 10".
func NormalizeArgs(args []string) []string {
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-l") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-l", args[i][2:]}, args[i+1:]...)...)
			// break
		}
		if strings.HasPrefix(args[i], "-n") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-n", args[i][2:]}, args[i+1:]...)...)
			// break
		}
		if strings.HasPrefix(args[i], "-b") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-b", args[i][2:]}, args[i+1:]...)...)
			// break
		}
		if strings.HasPrefix(args[i], "-a") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-a", args[i][2:]}, args[i+1:]...)...)
			// break
		}
	}
	return args
}

// GenerateStrings is a function that generates strings from the given length.
func GenerateStrings(length int, prefix string, counter int) []string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	if counter == 0 && length == 0 {
		fmt.Println("Err: suffix length must be greater than 0")
		os.Exit(1)
	}
	if length == 0 {
		return []string{prefix}
	}
	if length > 5 {
		fmt.Println("Error: suffix length must be less than or equal to 5")
		os.Exit(1)
	}

	var result []string
	for _, char := range alphabet {
		counter++
		result = append(result, GenerateStrings(length-1, prefix+string(char), counter)...)
	}

	return result
}

// Args is a struct that represents the arguments passed to the program.
type Args struct {
	LineCount int
	FileCount int
	ByteSize  int
	Args      []string
}

// IllegalArgsChecker is a function that checks if the arguments passed to the program are valid.
func IllegalArgsChecker(params Args) {
	lineSetCount := 0
	fileSetCount := 0
	byteSetCount := 0

	lineCount, fileCount, byteSize, args := params.LineCount, params.FileCount, params.ByteSize, params.Args

	fmt.Println(args)
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
				fmt.Println("Unknown option:", arg)
				os.Exit(1)
			}
		}
	}

	if lineSetCount+fileSetCount+byteSetCount > 1 {
		fmt.Println(
			`usage: split [-l line_count] [-a suffix_length] [file [prefix]]
			split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]
			split -n chunk_count [-a suffix_length] [file [prefix]]
			split -p pattern [-a suffix_length] [file [prefix]]`,
		)
		os.Exit(1)
	}

	if lineCount <= 0 && lineSetCount == 1 {
		fmt.Printf("split: %d: illegal line count\n", lineCount)
		os.Exit(1)
	}

	if fileCount <= 0 && fileSetCount == 1 {
		fmt.Printf("split: %d: illegal file count\n", fileCount)
		os.Exit(1)
	}

	if byteSize <= 0 && byteSetCount == 1 {
		fmt.Printf("split: %d: illegal byte size\n", byteSize)
		os.Exit(1)
	}
}
