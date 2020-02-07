package main

import (
	"fmt"
	"github.com/alrs/gitroots"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("gitroots requires a root directory as its sole argument")
	}
	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	isDir, err := dirTest(path)
	if err != nil {
		log.Fatal(err)
	}
	if !isDir {
		log.Fatalf("%s does not represent a directory", path)
	}
	repos, err := gitroots.Walk(path)
	if err != nil {
		log.Fatal(err)
	}
	for r, _ := range repos {
		fmt.Println(r)
	}
}

func dirTest(p string) (bool, error) {
	fi, err := os.Stat(p)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}
