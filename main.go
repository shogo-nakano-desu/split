// The main file of the split command.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	res, err := ParseArgs(fs)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	lineCount, fileCount, byteSize, suffixLen, args := res.LineCount, res.FileCount, res.ByteSize, res.SuffixLen, res.Args

	err = IllegalArgsChecker(Args{lineCount, fileCount, byteSize, args})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nonFlagArgs := fs.Args()
	reader := bufio.NewReader(os.Stdin)
	splitFileName, err := GetFileName(nonFlagArgs, reader)
	if err != nil {
		fmt.Printf("Error getting the file name: %v\n", err)
		os.Exit(1)
	}

	prefixFileName := ""
	if len(nonFlagArgs) >= 2 {
		prefixFileName = nonFlagArgs[1]
	}

	file, err := os.Open(splitFileName)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing the file: %v\n", err)
			os.Exit(1)
		}
	}()

	if lineCount > 0 {
		err := SplitByLinesMultithread(file, lineCount, prefixFileName, suffixLen)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	} else if fileCount > 0 {
		err := SplitByFileCounts(file, fileCount, prefixFileName, suffixLen)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	} else if byteSize > 0 {
		err := SplitByBytesMultithread(file, byteSize, prefixFileName, suffixLen)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Please specify a splitting option (-l, -n, -b).")
		os.Exit(1)
	}
}
