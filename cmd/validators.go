package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/linki/wanchain-cli/client"
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
	client, err := client.NewClient(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if validatorsParams.blockHeight == 0 {
		validatorsParams.blockHeight, err = client.GetCurrentBlockHeight()
		if err != nil {
			log.Fatal(err)
		}
	}

	validators, err := client.GetValidators(validatorsParams.blockHeight)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		spew.Dump(validators)
	}

	totalPower := util.PowerWeight(validators)

	for _, validator := range validators {
		t := table.NewWriter()
		t.SetOutputMirror(cmd.OutOrStdout())

		if validatorsParams.validatorAddress == "" || validator.Address == common.HexToAddress(validatorsParams.validatorAddress) {
			if debug {
				spew.Dump(validator)
			}

			validatorPower := util.PowerWeight([]types.Validator{validator})

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
				{"PowerWeight", fmt.Sprintf("%.2f%%", big.NewFloat(0).Mul(big.NewFloat(100), big.NewFloat(0).Quo(big.NewFloat(0).SetInt(validatorPower), big.NewFloat(0).SetInt(totalPower))))},
			})
		}

		util.RenderTable(t, format)
	}
}
