package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	files := listFiles("testdata")
	fmt.Println(strings.Join(files, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
