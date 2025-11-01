package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system requirements",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("== DevBootstrap System Check ==")

		// OS check
		fmt.Println("OS:", runtime.GOOS)
		if runtime.GOOS != "darwin" {
			fmt.Println("⚠️  Non-macOS detected. Full support coming later.")
		}

		checkBinary("zsh")
		checkBinary("git")
		checkBinary("brew")
		checkBinary("curl")

		fmt.Println("\n✅ System check complete.")
		return nil
	},
}

func checkBinary(name string) {
	_, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("❌ %s not found\n", name)
	} else {
		fmt.Printf("✔ %s installed\n", name)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
