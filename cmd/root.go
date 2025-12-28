// Package cmd is cobra cli
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jellybro99/sha256_cli/sha256"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sha",
	Short: "A go implementation of the sha256 hashing algorithm",
	Long: `sha is a CLI tool for the sha256 hashing algorithm.
	You can provide text as arguments or pipe it in via stdin.`,
	Args: cobra.ArbitraryArgs,
	Run:  runHasher,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("hash", "sha256", "Use the given hash function. Currently supported: sha256")
}

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
			fmt.Printf("%X\n", sha256.Hash(messages[0]))
			break
		}
		for _, message := range messages {
			fmt.Printf("%s: %X\n", message, sha256.Hash(message))
		}
	default:
		fmt.Println("Given hash is not supported")
	}
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
