package cmd

import (
	"fmt"

	"devbootstrap/internal/install"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a development tool (brew, asdf, node, etc.)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify a tool. Example: devbootstrap install brew")
	},
}

var installBrewCmd = &cobra.Command{
	Use:   "brew",
	Short: "Install Homebrew",
	RunE: func(cmd *cobra.Command, args []string) error {
		return install.InstallBrew()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.AddCommand(installBrewCmd)
}
