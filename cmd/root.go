// Package cmd is cobra cli
package cmd

import (
	"os"

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
	rootCmd.Flags().Bool("sha256", true, "Use the sha256 hashing algorithm")
	rootCmd.Flags().StringP("output", "o", "hex", "Output in the given format (hex, dec)")
}
