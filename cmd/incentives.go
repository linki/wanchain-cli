package cmd

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/linki/wanchain-cli/types"
	"github.com/linki/wanchain-cli/util"
)

var (
	incentivesCmd = &cobra.Command{
		Use: "incentives",
		Run: listIncentives,
	}
	incentivesParams = struct {
		validatorAddress string
		fromEpochID      uint64
		toEpochID        uint64
	}{}
)

func init() {
	incentivesCmd.PersistentFlags().StringVar(&incentivesParams.validatorAddress, "validator-address", "", "Address of the validator by which to filter")
	incentivesCmd.PersistentFlags().Uint64Var(&incentivesParams.fromEpochID, "from-epoch-id", firstEpochID, "Starting Epoch ID to query for")
	incentivesCmd.PersistentFlags().Uint64Var(&incentivesParams.toEpochID, "to-epoch-id", 0, "Last Epoch ID to query for")

	rootCmd.AddCommand(incentivesCmd)
}

func listIncentives(cmd *cobra.Command, _ []string) {
	client, err := rpc.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if incentivesParams.toEpochID == 0 {
		incentivesParams.toEpochID, err = util.GetCurrentEpochID(context.Background(), client)
		if err != nil {
			log.Fatal(err)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"Epoch ID", "Type", "Address", "Incentive"})
	t.SetAlign([]text.Align{text.AlignRight, text.AlignLeft, text.AlignLeft, text.AlignRight})

	for e := incentivesParams.fromEpochID; e <= incentivesParams.toEpochID; e++ {
		var incentives []types.Incentive
		if err := client.CallContext(context.Background(), &incentives, "pos_getEpochIncentivePayDetail", e); err != nil {
			log.Fatal(err)
		}
		if debug {
			spew.Dump(incentives)
		}

		for _, incentive := range incentives {
			if incentivesParams.validatorAddress == "" || common.HexToAddress(incentive.Address) == common.HexToAddress(incentivesParams.validatorAddress) {
				t.AppendRow(table.Row{e, incentive.Type, incentive.Address, util.WeiToEth(hexutil.MustDecodeBig(incentive.Incentive))})
			}
		}
	}

	t.Render()
}
