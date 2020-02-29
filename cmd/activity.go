package cmd

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/wanchain-cli/client"
	"github.com/linki/wanchain-cli/util"
)

var (
	activityCmd = &cobra.Command{
		Use: "activity",
		Run: listActivity,
	}
	activityParams = struct {
		validatorAddress string
		fromEpochID      uint64
		toEpochID        uint64
	}{}
)

func init() {
	activityCmd.PersistentFlags().StringVar(&activityParams.validatorAddress, "validator-address", "", "Address of the validator by which to filter")
	activityCmd.PersistentFlags().Uint64Var(&activityParams.fromEpochID, "from-epoch-id", 0, "Starting Epoch ID to query for, defaults to the last three epochs")
	activityCmd.PersistentFlags().Uint64Var(&activityParams.toEpochID, "to-epoch-id", 0, "Last Epoch ID to query for")

	rootCmd.AddCommand(activityCmd)
}

func listActivity(cmd *cobra.Command, _ []string) {
	client, err := client.NewClient(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if activityParams.toEpochID == 0 {
		activityParams.toEpochID, err = client.GetCurrentEpochID()
		if err != nil {
			log.Fatal(err)
		}
	}

	if activityParams.fromEpochID == 0 {
		activityParams.fromEpochID = activityParams.toEpochID - defaultEpochRange
	}

	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"Epoch ID", "Role", "Address", "Active", "Blocks"})
	t.SetAlign([]text.Align{text.AlignRight, text.AlignLeft, text.AlignLeft, text.AlignLeft, text.AlignRight})

	for e := activityParams.fromEpochID; e <= activityParams.toEpochID; e++ {
		activity, err := client.GetActivity(e)
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			spew.Dump(activity)
		}

		for i, addr := range activity.EPLeader {
			if activityParams.validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(activityParams.validatorAddress) {
				t.AppendRow(table.Row{e, "EP", addr, activity.EPActivity[i] == 1, ""})
			}
		}

		for i, addr := range activity.RPLeader {
			if activityParams.validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(activityParams.validatorAddress) {
				t.AppendRow(table.Row{e, "RP", addr, activity.RPActivity[i] == 1, ""})
			}
		}

		for i, addr := range activity.SLTLeader {
			if activityParams.validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(activityParams.validatorAddress) {
				t.AppendRow(table.Row{e, "SL", addr, "", activity.SLBlocks[i]})
			}
		}
	}

	util.RenderTable(t, format)
}
