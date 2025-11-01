package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devbootstrap",
	Short: "Bootstrap a dev machine with essential tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand. Try 'devbootstrap init'")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// you add subcommands in other files
}
