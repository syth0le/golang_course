package main

import (
	//"fmt"
	"io"
	"os"
	//"path/filepath"
	//"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
	println(err.Error())
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	println(out, path, printFiles)
	return nil
}
