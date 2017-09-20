package main

import (
	"fmt"
	"github.com/alrs/gitroots"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("gitroots requires a root directory as its sole argument")
	}
	repos, err := gitroots.Walk(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for r, _ := range repos {
		fmt.Println(r)
	}
}
