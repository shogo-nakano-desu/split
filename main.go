package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var lineCount int
	var fileCount int
	var byteSize int
	var isLineFlagSet, isFileFlagSet, isByteFlagSet bool

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.IntVar(&lineCount, "l", 0, "Number of lines per split file.")
	fs.IntVar(&fileCount, "n", 0, "Number of files to split into.")
	fs.IntVar(&byteSize, "b", 0, "Number of bytes per split file.")

	fs.BoolVar(&isLineFlagSet, "lflag", false, "Indicate if -l flag is set")
	fs.BoolVar(&isFileFlagSet, "nflag", false, "Indicate if -n flag is set")
	fs.BoolVar(&isByteFlagSet, "bflag", false, "Indicate if -b flag is set")

	// TODO 引数を見て、フラグが２つ以上指定された場合にはsplitと同じエラーを出すようにする。
	args := NormalizeArgs(os.Args[1:])

	fs.Parse(args)

	if lineCount <= 0 && isLineFlagSet {
		fmt.Printf("split: %d: illegal line count\n", lineCount)
		os.Exit(1)
	}

	if fileCount <= 0 && isFileFlagSet {
		fmt.Printf("split: %d: illegal file count\n", fileCount)
		os.Exit(1)
	}

	if byteSize <= 0 && isByteFlagSet {
		fmt.Printf("split: %d: illegal byte size\n", byteSize)
		os.Exit(1)
	}

	nonFlagArgs := fs.Args()
	if len(nonFlagArgs) <= 0 {
		// TODO: ファイル名がなかった場合には、追加でファイル名が与えられるのを待つようにする。
		fmt.Printf("Please provide a file to split.")
		os.Exit(1)
	}
	fileName := nonFlagArgs[0]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	if lineCount > 0 {
		SplitByLines(file, lineCount, fileName)
	} else if fileCount > 0 {
		SplitByFileCounts(file, fileCount, fileName)
	} else if byteSize > 0 {
		SplitByBytes(file, byteSize, fileName)
	} else {
		fmt.Println("Please specify a splitting option (-l, -n, -b).")
		return
	}
}
