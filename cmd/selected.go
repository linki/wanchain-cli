package cmd

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"github.com/wanchain/go-wanchain/common"

	"github.com/linki/wanchain-cli/client"
	"github.com/linki/wanchain-cli/util"
)

var (
	selectedCmd = &cobra.Command{
		Use: "selected",
		Run: listSelected,
	}
	selectedParams = struct {
		validatorAddress string
		epochID          uint64
	}{}
)

func init() {
	selectedCmd.PersistentFlags().StringVar(&selectedParams.validatorAddress, "validator-address", "", "Address of the validator by which to filter")
	selectedCmd.PersistentFlags().Uint64Var(&selectedParams.epochID, "epoch-id", 0, "Epoch ID to query for")

	rootCmd.AddCommand(selectedCmd)
}

func listSelected(cmd *cobra.Command, _ []string) {
	client, err := client.NewClient(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if selectedParams.epochID == 0 {
		selectedParams.epochID, err = client.GetCurrentEpochID()
		if err != nil {
			log.Fatal(err)
		}
		selectedParams.epochID++
	}

	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"Epoch ID", "Role", "Address"})

	EPAddrs, err := client.GetEpochLeaders(selectedParams.epochID)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(EPAddrs)
	}

	RPAddrs, err := client.GetRandomProposers(selectedParams.epochID)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(RPAddrs)
	}

	for _, addr := range EPAddrs {
		if selectedParams.validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(selectedParams.validatorAddress) {
			t.AppendRow(table.Row{selectedParams.epochID, "EP", addr})
		}
	}

	for _, addr := range RPAddrs {
		if selectedParams.validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(selectedParams.validatorAddress) {
			t.AppendRow(table.Row{selectedParams.epochID, "RP", addr})
		}
	}

	util.RenderTable(t, format)
}
