package main

import (
	"os"

	"github.com/jellybro99/sha/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
