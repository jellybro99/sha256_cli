package main

import (
	"fmt"
	"os"
)

func main() {
	message := os.Args[1]
	// check if os.Args[1] exists
	hash := sha256Hash(message)
	fmt.Println(hash)
}
