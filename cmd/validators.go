package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/linki/wanchain-cli/types"
	"github.com/linki/wanchain-cli/util"
)

var (
	validatorsCmd = &cobra.Command{
		Use: "validators",
		Run: listValidators,
	}
	validatorsParams = struct {
		validatorAddress string
		blockHeight      uint64
	}{}
)

func init() {
	validatorsCmd.PersistentFlags().StringVar(&validatorsParams.validatorAddress, "validator-address", "", "Address of the validator by which to filter")
	validatorsCmd.PersistentFlags().Uint64Var(&validatorsParams.blockHeight, "block-height", 0, "Block height to query for")

	rootCmd.AddCommand(validatorsCmd)
}

func listValidators(cmd *cobra.Command, _ []string) {
	client, err := rpc.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if validatorsParams.blockHeight == 0 {
		validatorsParams.blockHeight, err = util.GetCurrentBlockHeight(context.Background(), client)
		if err != nil {
			log.Fatal(err)
		}
	}

	var validators []types.Validator
	if err := client.CallContext(context.Background(), &validators, "pos_getStakerInfo", validatorsParams.blockHeight); err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(validators)
	}

	for _, validator := range validators {
		t := table.NewWriter()
		t.SetOutputMirror(cmd.OutOrStdout())

		if validatorsParams.validatorAddress == "" || validator.Address == common.HexToAddress(validatorsParams.validatorAddress) {
			if debug {
				spew.Dump(validator)
			}

			t.AppendRows([]table.Row{
				{"Address", validator.Address.Hex()},
				{"PubSec256", validator.PubSec256},
				{"PubBn256", validator.PubBn256},
				{"Amount", util.WeiToEth(hexutil.MustDecodeBig(validator.Amount))},
				{"VotingPower", util.WeiToEth(hexutil.MustDecodeBig(validator.VotingPower))},
				{"LockEpochs", validator.LockEpochs},
				{"NextLockEpochs", validator.NextLockEpochs},
				{"From", validator.From.Hex()},
				{"StakingEpoch", validator.StakingEpoch},
				{"FeeRate", fmt.Sprintf("%.2f%%", float64(validator.FeeRate)/100)},
				{"# Delegators", len(validator.Clients)},
				{"# Partners", len(validator.Partners)},
				{"MaxFeeRate", fmt.Sprintf("%.2f%%", float64(validator.MaxFeeRate)/100)},
				{"FeeRateChangedEpoch", validator.FeeRateChangedEpoch},
			})
		}

		t.Render()
	}
}
