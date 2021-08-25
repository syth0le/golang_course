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

type Node interface {
	fmt.Stringer
}

type Directory struct{
	name string
	children []Node
}

type File struct{
	name string
	size int64
}

func (dir Directory) String() string {
	return dir.name
}

func (file File) String() string {
	if file.size != 0 {
		return fmt.Sprintf("%s (%db)", file.name, file.size)
	}
	return fmt.Sprintf("%s (empty)", file.name)
}

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
	nodes := readDir(out , path, printFiles, []Node{})
	printDir(out, nodes)
	return nil
}


func readDir(out io.Writer, path string, printFiles bool, nodes []Node) []Node {
	dir, _ := ioutil.ReadDir(path)
	var newNode Node

	for _, elem := range dir {

		//_ = toStringRecreator(elem)

		switch elem.IsDir(){
		case true:
			levelPath := filepath.Join(path, elem.Name())
			//println("\n"+levelPath)
			children := readDir(out, levelPath, printFiles, []Node{})
			newNode = Directory{elem.Name(), children}
		case false:
			newNode = File{elem.Name(), elem.Size()}

		}

		nodes = append(nodes, newNode)

	}
	return nodes
}

func printDir(out io.Writer, nodes []Node){
	if len(nodes) == 1 {
		fmt.Fprintf(out, "%s%s\n", "└───", nodes[0])
		if directory, ok := nodes[0].(Directory); ok {
			printDir(out, directory.children)
		}
		return
	}
	fmt.Fprintf(out, "%s%s\n", "├───", nodes[0])
	if directory, ok := nodes[0].(Directory); ok {
		printDir(out, directory.children)
	}
	printDir(out, nodes[1:])
}


func toStringRecreator (elem fs.FileInfo) error {
	if elem == nil{
		panic("can't get elements from slice")
	}
	toReturn := fmt.Sprintf("ELEMENT: name=%s size=%d isDir=%v", elem.Name(), elem.Size(), elem.IsDir())
	println(toReturn)

	return nil
}