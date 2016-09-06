package main

import (
	"fmt"
	"os"

	"github.com/cucumber/cucumber-pretty-formatter"
	_ "github.com/cucumber/cucumber-pretty-formatter/progress"
)

func main() {
	if err := formatter.Run(os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
