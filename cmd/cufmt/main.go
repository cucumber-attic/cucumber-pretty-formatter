package main

import (
	"log"

	"github.com/cucumber/cucumber-pretty-formatter"
	_ "github.com/cucumber/cucumber-pretty-formatter/pretty"
)

func main() {
	if err := formatter.Run(); err != nil {
		log.Fatal(err)
	}
}
