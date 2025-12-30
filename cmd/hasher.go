package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jellybro99/sha256_cli/sha256"
	"github.com/spf13/cobra"
)

func runHasher(cmd *cobra.Command, args []string) {
	hashFunction, err := cmd.Flags().GetString("hash")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	messages, err := getInputs(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(messages) == 0 {
		fmt.Println("no input given")
	}

	switch hashFunction {
	case "sha256":
		if len(messages) == 1 {
			fmt.Println(formatHash(sha256.Hash(messages[0])))
			break
		}
		for _, message := range messages {
			fmt.Printf("%s: %s\n", message, formatHash(sha256.Hash(message)))
		}
	default:
		fmt.Println("Given hash function is not supported")
	}
}

func formatHash(hash [8]uint32) string {
	var sb strings.Builder

	for _, word := range hash {
		fmt.Fprintf(&sb, "%X ", word)
	}

	return sb.String()
}

func getInputs(args []string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}

	file, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if file.Size() == 0 {
		return nil, errors.New("no input")
	}

	message, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	return []string{string(message)}, nil
}
