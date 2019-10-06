package cmd

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"github.com/linki/wanchain-cli/client"
	"github.com/linki/wanchain-cli/util"
)

var (
	currentCmd = &cobra.Command{
		Use: "current",
		Run: displayCurrent,
	}
)

func init() {
	rootCmd.AddCommand(currentCmd)
}

func displayCurrent(cmd *cobra.Command, _ []string) {
	client, err := client.NewClient(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	epochID, err := client.GetCurrentEpochID()
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(epochID)
	}

	slotID, err := client.GetCurrentSlotID()
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(slotID)
	}

	blockHeight, err := client.GetCurrentBlockHeight()
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(blockHeight)
	}

	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())

	t.AppendRows([]table.Row{
		{"Epoch ID", epochID},
		{"Slot ID", slotID},
		{"Block Height", blockHeight},
	})

	util.RenderTable(t, format)
}
