package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var wg sync.WaitGroup
var targetFolder string

func listDir(searchStart string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(searchStart)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		scanFile(searchStart, file)
	}

	wg.Done()
	return files, nil
}

func scanFile(path string, file os.FileInfo) {
	if file.IsDir() {
		newDirFound(path, file)
	}
}

func newDirFound(path string, dir os.FileInfo) {
	dirName := dir.Name()
	if isNodeFolder(dirName) {
		fmt.Println(path + "/" + dirName)
	} else {
		wg.Add(1)
		go listDir(path + "/" + dirName)
	}
}

func isNodeFolder(name string) bool {
	return name == targetFolder
}

func main() {
	wg.Add(1)
	mainPath := os.Args[1]
	targetFolder = os.Args[2]
	go listDir(mainPath)
	wg.Wait()
}
