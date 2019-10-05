package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultEpochRange = 3
)

var (
	rpcURL string
	debug  bool
)

var (
	rootCmd = &cobra.Command{}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&rpcURL, "rpc", "https://mywanwallet.io/api", "Wanchain RPC endpoint")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Print parsed RPC response for debugging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
