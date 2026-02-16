package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ki",
	Short: "KI - The Dragon Ball CLI",
	Long: `KI is a command line interface for interacting with the Dragon Ball API.

Explore characters, planets, and more from the Dragon Ball universe.

Use 'ki tui' to launch the interactive Terminal User Interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🐉 Welcome to KI: The Dragon Ball CLI")
		fmt.Println()
		fmt.Println("Use 'ki tui' to launch the interactive interface")
		fmt.Println("Use 'ki --help' to see all available commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
