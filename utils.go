package main

import (
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
