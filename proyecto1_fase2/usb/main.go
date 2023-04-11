package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	dir := "/media"
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Println(filepath.Join(dir, file.Name()))
	}
}
