package main

import (
	"fmt"
	"os"

	"github.com/cucumber/cucumber-pretty-formatter"
	_ "github.com/cucumber/cucumber-pretty-formatter/progress"
)

func main() {
	if err := formatter.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
