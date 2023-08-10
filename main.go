package main

import (
	"flag"
	"fmt"
	"os"
)

// Usage
// split [-l line_count] [-a suffix_length] [file [prefix]]

// TODOs
// ・ファイル名がなかった場合には、追加でファイル名が与えられるのを待つようにする。
// ・でかいファイルでも高速に読み込むことができるようにする。

func main() {
	var lineCount int
	// lineSet := false
	var fileCount int
	// fileSet := false
	var byteSize int
	// byteSet := false
	var suffixLen int

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.IntVar(&lineCount, "l", 0, "Number of lines per split file.")
	fs.IntVar(&fileCount, "n", 0, "Number of files to split into.")
	fs.IntVar(&byteSize, "b", 0, "Number of bytes per split file.")
	fs.IntVar(&suffixLen, "a", 2, "Suffix length.")

	args := NormalizeArgs(os.Args[1:])

	fs.Parse(args)

	IllegalArgsChecker(A{lineCount, fileCount, byteSize, args})

	nonFlagArgs := fs.Args()
	if len(nonFlagArgs) <= 0 {
		// TODO: ファイル名がなかった場合には、追加でファイル名が与えられるのを待つようにする。
		fmt.Printf("Please provide a file to split.")
		os.Exit(1)
	}
	splitFileName := nonFlagArgs[0]

	prefixFileName := ""
	if len(nonFlagArgs) >= 2 {
		prefixFileName = nonFlagArgs[1]
	}

	file, err := os.Open(splitFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	if lineCount > 0 {
		SplitByLines(file, lineCount, prefixFileName, suffixLen)
	} else if fileCount > 0 {
		SplitByFileCounts(file, fileCount, prefixFileName, suffixLen)
	} else if byteSize > 0 {
		SplitByBytes(file, byteSize, prefixFileName, suffixLen)
	} else {
		fmt.Println("Please specify a splitting option (-l, -n, -b).")
		return
	}
}
