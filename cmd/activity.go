package cmd

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/linki/wanchain-cli/types"
)

var (
	validatorAddress string
	fromEpochID      uint64
	toEpochID        uint64
)

var (
	activityCmd = &cobra.Command{
		Use: "activity",
		Run: listActivity,
	}
)

func init() {
	activityCmd.PersistentFlags().StringVar(&validatorAddress, "validator-address", "", "Address of the validator by which to filter")
	activityCmd.PersistentFlags().Uint64Var(&fromEpochID, "from-epoch-id", firstEpochID, "Starting Epoch ID to query for")
	activityCmd.PersistentFlags().Uint64Var(&toEpochID, "to-epoch-id", 0, "Last Epoch ID to query for")

	rootCmd.AddCommand(activityCmd)
}

func listActivity(cmd *cobra.Command, _ []string) {
	client, err := rpc.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if toEpochID == 0 {
		toEpochID, err = getCurrentEpochID(context.Background(), client)
		if err != nil {
			log.Fatal(err)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"Epoch ID", "Role", "Address", "Active", "Blocks"})

	for e := fromEpochID; e <= toEpochID; e++ {
		var activity types.Activity
		if err := client.CallContext(context.Background(), &activity, "pos_getActivity", e); err != nil {
			log.Fatal(err)
		}
		if debug {
			spew.Dump(activity)
		}

		for i, addr := range activity.EPLeader {
			if validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(validatorAddress) {
				t.AppendRow(table.Row{e, "EP", addr, activity.EPActivity[i] == 1, ""})
			}
		}

		for i, addr := range activity.RPLeader {
			if validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(validatorAddress) {
				t.AppendRow(table.Row{e, "RP", addr, activity.RPActivity[i] == 1, ""})
			}
		}

		for i, addr := range activity.SLTLeader {
			if validatorAddress == "" || common.HexToAddress(addr) == common.HexToAddress(validatorAddress) {
				t.AppendRow(table.Row{e, "SL", addr, "", activity.SLBlocks[i]})
			}
		}
	}

	t.Render()
}

func getCurrentEpochID(ctx context.Context, client *rpc.Client) (uint64, error) {
	var epochID uint64
	if err := client.CallContext(ctx, &epochID, "pos_getEpochID"); err != nil {
		return 0, err
	}
	if debug {
		spew.Dump(epochID)
	}

	return epochID, nil
}
