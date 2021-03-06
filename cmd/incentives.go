package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/linki/wanchain-cli/client"
	"github.com/linki/wanchain-cli/util"
)

var (
	incentivesCmd = &cobra.Command{
		Use: "incentives",
		Run: listIncentives,
	}
	incentivesParams = struct {
		validatorAddress string
		delegatorAddress string
		fromEpochID      int64
		toEpochID        int64
		filterType       string
	}{}
)

func init() {
	incentivesCmd.PersistentFlags().StringVar(&incentivesParams.validatorAddress, "validator-address", "", "Address of the validator by which to filter")
	incentivesCmd.PersistentFlags().StringVar(&incentivesParams.delegatorAddress, "delegator-address", "", "Address of the delegator by which to filter")
	incentivesCmd.PersistentFlags().Int64Var(&incentivesParams.fromEpochID, "from-epoch-id", defaultEpochRange, "Starting Epoch ID to query for, defaults to the last three epochs before --to-epoch-id")
	incentivesCmd.PersistentFlags().Int64Var(&incentivesParams.toEpochID, "to-epoch-id", 0, "Last Epoch ID to query for, defaults to the current Epoch ID")
	incentivesCmd.PersistentFlags().StringVar(&incentivesParams.filterType, "filter-type", "", "Filter by incentive type, either validator or delegator")

	rootCmd.AddCommand(incentivesCmd)
}

func listIncentives(cmd *cobra.Command, _ []string) {
	client, err := client.NewClient(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var (
		fromEpochID = uint64(incentivesParams.fromEpochID)
		toEpochID   = uint64(incentivesParams.toEpochID)
	)

	if incentivesParams.toEpochID <= 0 {
		currentEpochID, err := client.GetCurrentEpochID()
		if err != nil {
			log.Fatal(err)
		}
		toEpochID += currentEpochID
	}

	if incentivesParams.fromEpochID <= 0 {
		fromEpochID += toEpochID
	}

	sumValidator := big.NewInt(0)
	sumDelegator := big.NewInt(0)

	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"Epoch ID", "Type", "Address", "Incentive"})
	t.SetAlign([]text.Align{text.AlignRight, text.AlignLeft, text.AlignLeft, text.AlignRight})

	for e := fromEpochID; e <= toEpochID; e++ {
		incentives, err := client.GetIncentives(e)
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			spew.Dump(incentives)
		}

		if incentivesParams.filterType == "" || incentivesParams.filterType == "validator" {
			for _, validatorIncentive := range incentives {
				if (incentivesParams.validatorAddress == "" && incentivesParams.delegatorAddress == "") || common.HexToAddress(validatorIncentive.Address) == common.HexToAddress(incentivesParams.validatorAddress) {
					t.AppendRow(table.Row{e, validatorIncentive.Type, validatorIncentive.Address, fmt.Sprintf("%.8f", util.WeiToEth(hexutil.MustDecodeBig(validatorIncentive.Incentive)))})
					sumValidator.Add(sumValidator, hexutil.MustDecodeBig(validatorIncentive.Incentive))
				}
			}
		}

		if incentivesParams.filterType == "" || incentivesParams.filterType == "delegator" {
			for _, validatorIncentive := range incentives {
				for _, delegatorIncentive := range validatorIncentive.Delegators {
					if (incentivesParams.validatorAddress == "" && incentivesParams.delegatorAddress == "") || common.HexToAddress(validatorIncentive.Address) == common.HexToAddress(incentivesParams.validatorAddress) || common.HexToAddress(delegatorIncentive.Address) == common.HexToAddress(incentivesParams.delegatorAddress) {
						t.AppendRow(table.Row{e, delegatorIncentive.Type, delegatorIncentive.Address, fmt.Sprintf("%.8f", util.WeiToEth(hexutil.MustDecodeBig(delegatorIncentive.Incentive)))})
						sumDelegator.Add(sumDelegator, hexutil.MustDecodeBig(delegatorIncentive.Incentive))
					}
				}
			}
		}
	}

	t.SetAlignFooter([]text.Align{text.AlignRight, text.AlignLeft, text.AlignLeft, text.AlignRight})
	t.AppendFooter(table.Row{"", "validator", "", fmt.Sprintf("%.8f", util.WeiToEth(sumValidator))})
	t.AppendFooter(table.Row{"", "delegator", "", fmt.Sprintf("%.8f", util.WeiToEth(sumDelegator))})

	util.RenderTable(t, format)
}
