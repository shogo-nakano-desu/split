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
			break
		}
		if strings.HasPrefix(args[i], "-n") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-n", args[i][2:]}, args[i+1:]...)...)
			break
		}
		if strings.HasPrefix(args[i], "-b") && len(args[i]) > 2 {
			args = append(args[:i], append([]string{"-b", args[i][2:]}, args[i+1:]...)...)
			break
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
		fmt.Println("Error: suffix length must be less than 5")
		os.Exit(1)
	}

	var result []string
	for _, char := range alphabet {
		counter++
		result = append(result, GenerateStrings(length-1, prefix+string(char), counter)...)
	}

	return result
}
