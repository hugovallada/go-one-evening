package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	files := listFiles("testdata")
	fmt.Println(strings.Join(files, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := ioutil.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
