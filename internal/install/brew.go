package install

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// InstallBrew installs Homebrew on macOS and ensures PATH is correct.
func InstallBrew() error {
	brewPath := getBrewPrefix()

	// 1. Check if brew already exists
	if brewPath != "" {
		fmt.Println("✔ Homebrew already installed")

		// Update PATH for current session if needed
		if !strings.Contains(os.Getenv("PATH"), brewPath) {
			updatePathForCurrentSession(brewPath)
			fmt.Println("✅ Homebrew added to PATH for current session")
		}

		fmt.Println("⚠️  Brew was already installed. No action needed.")
		fmt.Println("⚠️  To use brew in this terminal, run:")
		fmt.Printf("   eval \"$(%s/brew shellenv)\"\n", brewPath)
		return nil
	}

	// 2. Ensure we are in an interactive terminal
	if !isTerminal() {
		fmt.Println("⚠️  Homebrew installation requires an interactive terminal.")
		fmt.Println("Please run:")
		fmt.Println("  ./devbootstrap install brew")
		return nil
	}

	fmt.Println("❌ Homebrew not found. Installing...")

	// 3. Run the official installer
	cmd := exec.Command("/bin/bash", "-c", `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("brew installation failed: %w", err)
	}

	// 4. Detect where brew got installed
	brewPath = getBrewPrefix()
	if brewPath == "" {
		return fmt.Errorf("brew installed but could not detect path")
	}

	fmt.Println("✅ Homebrew installed successfully")

	// 5. Update PATH for current session
	updatePathForCurrentSession(brewPath)
	fmt.Println("✅ Homebrew added to PATH for current session")

	// 6. Add brew to shell config for future sessions
	if err := addBrewToShellConfig(brewPath); err != nil {
		fmt.Println("⚠️ Could not automatically update shell config:", err)
	}
	fmt.Println("✅ To use brew in future shells, it has been added to your shell profile.")

	return nil
}

// isTerminal returns true if stdin is a TTY
func isTerminal() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// getBrewPrefix returns brew install path if it exists
func getBrewPrefix() string {
	paths := []string{"/opt/homebrew/bin", "/usr/local/bin"}
	for _, p := range paths {
		if _, err := os.Stat(p + "/brew"); err == nil {
			return p
		}
	}
	return ""
}

// updatePathForCurrentSession prepends brew to PATH for this process
func updatePathForCurrentSession(brewPath string) {
	currentPath := os.Getenv("PATH")
	if !strings.Contains(currentPath, brewPath) {
		os.Setenv("PATH", brewPath+":"+currentPath)
	}
}

// addBrewToShellConfig appends brew eval line to user's shell file if not present
func addBrewToShellConfig(brewPath string) error {
	shellFile := os.ExpandEnv("$HOME/.zprofile") // zsh default
	evalLine := fmt.Sprintf("eval \"$(%s/brew shellenv)\"\n", brewPath)

	// check if line already exists
	content, err := os.ReadFile(shellFile)
	if err == nil && strings.Contains(string(content), evalLine) {
		return nil // already present
	}

	f, err := os.OpenFile(shellFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(evalLine)
	return err
}
