package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timmattison/golang-cli-tool-template/internal"
)

const (
	Flag1Long  = "flag-1"
	Flag1Short = "1"
	Flag2Long  = "flag-2"
	Flag2Short = "2"
)

var (
	subCommand1Flag1 *string
	subCommand1Flag2 *bool
)

var subCommand1 = &cobra.Command{
	Use:   "subCommand1",
	Short: "Do subCommand1 stuff",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement the sub-command
		fmt.Println("Got into sub-command 1")

		fmt.Printf("Flag 1: %s\n", *subCommand1Flag1)
		fmt.Printf("Flag 2: %t\n", *subCommand1Flag2)
		fmt.Printf("Verbose: %t\n", *internal.Verbose)
	},
}

func init() {
	// Make sure rootCmd picks up this sub-command
	rootCmd.AddCommand(subCommand1)

	// TODO put your flags here
	// Two flags for this sub-command
	subCommand1Flag1 = subCommand1.PersistentFlags().StringP(Flag1Long, Flag1Short, "", "sub-command 1 flag 1")
	subCommand1Flag2 = subCommand1.PersistentFlags().BoolP(Flag2Long, Flag2Short, false, "sub-command 1 flag 2")

	// The first flag is required or we panic
	if err := subCommand1.MarkPersistentFlagRequired(Flag1Long); err != nil {
		panic(err)
	}
}
