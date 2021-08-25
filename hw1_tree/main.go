package main

import (
	"fmt"
	"path/filepath"

	//"fmt"
	"io"
	"io/fs"
	"io/ioutil"
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
	//println(out, path, printFiles)
	dir, _ := ioutil.ReadDir(path)

	for _, elem := range dir {

		_ = toStringRecognizer(elem)

		if elem.IsDir() {
			levelPath := filepath.Join(path, elem.Name())
			println("\n"+levelPath)

			dirTree(out, levelPath, printFiles)
		}
	}

	return nil
}

func toStringRecognizer(elem fs.FileInfo) error {
	if elem == nil{
		panic("can't get elements from slice")
	}
	toReturn := fmt.Sprintf("ELEMENT: name=%s size=%d isDir=%v", elem.Name(), elem.Size(), elem.IsDir())
	println(toReturn)

	return nil
}


func readDir(){

}

func printDir(){

}