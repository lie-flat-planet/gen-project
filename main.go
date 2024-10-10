package main

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/command"
	"os"
)

func main() {
	if err := command.Dir.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
