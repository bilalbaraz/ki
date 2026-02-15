package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ki",
	Short: "KI - The Dragon Ball CLI",
	Long:  "KI is a command line interface for interacting with the Dragon Ball API.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to KI: The Dragon Ball CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
