package main

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/cmd"
	"os"
)

func main() {
	if err := cmd.DirCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
