package main

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/command"
	"os"
)

func main() {
	if err := command.RootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
