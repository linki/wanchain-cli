package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "undefined"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
