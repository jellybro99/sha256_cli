package main

import (
	"fmt"
	"os"
)

func main() {
	message := os.Args[1]
	hash := sha256(message)
	fmt.Println(hash)
}
