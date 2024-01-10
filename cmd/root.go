package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/timmattison/golang-cli-tool-template/internal"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   internal.ProgramName,
	Short: internal.ProgramDescription,
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
	// TODO - put global flags here
	internal.Verbose = rootCmd.PersistentFlags().BoolP("verbose", "v", false, "toggle verbose output")
}
