package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 1. Read the input argument, which is the name of the executable file that the
// utility will be searching for.

// 2. Read the value stored in the PATH environment variable, splitting it, and
// iterating over the directories of the PATH variable.

// 3. Look for the desired binary file in these directories and determine whether
// it can be found or not, whether it is a regular file, and whether it is an
// executable file.

func main() {
	//	Read arguments
	args := os.Args
	//	Check if arguments were provided
	if len(args) == 1 {
		fmt.Fprint(os.Stderr, "please provide an argument!")
		return
	}
	//	Store input argument
	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, p := range pathSplit {
		fullPath := filepath.Join(p, file)
		// Does it exist
		stat, err := os.Stat(fullPath)
		if err != nil {
			continue
		}
		mode := stat.Mode()
		// Check if it is a regular file
		if !mode.IsRegular() {
			continue
		}
		//	It is executable
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}

}
