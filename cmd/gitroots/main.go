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
	repos, err := gitroots.Walk(path)
	if err != nil {
		log.Fatal(err)
	}
	for r, _ := range repos {
		fmt.Println(r)
	}
}
