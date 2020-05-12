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
	totalAmount := util.TotalAmountValidators(validators)

	for _, validator := range validators {
		if validatorsParams.validatorAddress == "" || validator.Address == common.HexToAddress(validatorsParams.validatorAddress) {
			if debug {
				spew.Dump(validator)
			}

			if debug {
				for _, client := range validator.Clients {
					t := table.NewWriter()
					t.SetOutputMirror(cmd.OutOrStdout())

					t.AppendRows([]table.Row{
						{"Address", client.Address.Hex()},
						{"Amount", fmt.Sprintf("%.2f", util.WeiToEth(hexutil.MustDecodeBig(client.Amount)))},
						{"VotingPower", fmt.Sprintf("%.2f", util.WeiToEth(hexutil.MustDecodeBig(client.VotingPower)))},
						{"QuitEpoch", client.QuitEpoch},
					})

					util.RenderTable(t, format)
				}

				for _, partner := range validator.Partners {
					t := table.NewWriter()
					t.SetOutputMirror(cmd.OutOrStdout())

					t.AppendRows([]table.Row{
						{"Address", partner.Address.Hex()},
						{"Amount", fmt.Sprintf("%.2f", util.WeiToEth(hexutil.MustDecodeBig(partner.Amount)))},
						{"VotingPower", fmt.Sprintf("%.2f", util.WeiToEth(hexutil.MustDecodeBig(partner.VotingPower)))},
						{"Renewal", partner.Renewal},
						{"LockEpochs", partner.LockEpochs},
						{"StakingEpoch", partner.StakingEpoch},
					})

					util.RenderTable(t, format)
				}
			}

			t := table.NewWriter()
			t.SetOutputMirror(cmd.OutOrStdout())

			validatorPower := util.PowerWeight([]types.Validator{validator})
			validatorAmount := util.TotalAmountValidators([]types.Validator{validator})

			clientAmount := util.TotalAmountClients(validator.Clients)
			partnerAmount := util.TotalAmountPartners(validator.Partners)

			stakeOutAmount := util.TotalAmountStakeOuts(validator.Clients)

			t.AppendRows([]table.Row{
				{"Address", validator.Address.Hex()},
				{"PubSec256", validator.PubSec256},
				{"PubBn256", validator.PubBn256},
				{"Amount", fmt.Sprintf("%.2f", util.WeiToEth(hexutil.MustDecodeBig(validator.Amount)))},
				{"VotingPower", fmt.Sprintf("%.2f (%.2fx)", util.WeiToEth(hexutil.MustDecodeBig(validator.VotingPower)), big.NewFloat(0).Quo(big.NewFloat(0).Quo(big.NewFloat(0).SetInt(hexutil.MustDecodeBig(validator.VotingPower)), big.NewFloat(0).SetInt(hexutil.MustDecodeBig(validator.Amount))), big.NewFloat(1000)))},
				{"LockEpochs", validator.LockEpochs},
				{"NextLockEpochs", validator.NextLockEpochs},
				{"From", validator.From.Hex()},
				{"StakingEpoch", validator.StakingEpoch},
				{"FeeRate", fmt.Sprintf("%.2f%%", float64(validator.FeeRate)/100)},
				{"# Delegators", fmt.Sprintf("%d (%.2f)", len(validator.Clients), util.WeiToEth(clientAmount))},
				{"# Partners", fmt.Sprintf("%d (%.2f)", len(validator.Partners), util.WeiToEth(partnerAmount))},
				{"ValidatorAmount", fmt.Sprintf("%.2f (of %.2f)", util.WeiToEth(validatorAmount), util.WeiToEth(totalAmount))},
				{"PendingStakeOut", util.WeiToEth(stakeOutAmount)},
				{"StakeWeight", fmt.Sprintf("%.2f%%", util.Weight(validatorAmount, totalAmount))},
				{"PowerWeight", fmt.Sprintf("%.2f%%", util.Weight(validatorPower, totalPower))},
				{"MaxFeeRate", fmt.Sprintf("%.2f%%", float64(validator.MaxFeeRate)/100)},
				{"FeeRateChangedEpoch", validator.FeeRateChangedEpoch},
			})

			util.RenderTable(t, format)
		}
	}
}
