package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "byopm",
	Short: "byopm - Build-Your-Own Password Manager",
	Long: `BYOPM is a CLI-based password manager.
This application is a tool to store your password in CLI way.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
