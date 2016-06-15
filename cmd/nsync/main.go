package main

import (
	"fmt"
	"os"

	"github.com/cf-furnace/k8s-nsync/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
