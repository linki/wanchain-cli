package util

import (
	"math"
	"math/big"

	"github.com/jedib0t/go-pretty/table"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/linki/wanchain-cli/types"
)

func WeiToEth(eth *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(eth), big.NewFloat(math.Pow10(18)))
}

func PowerWeight(validators []types.Validator) *big.Int {
	power := big.NewInt(0)
	for _, validator := range validators {
		power = power.Add(power, hexutil.MustDecodeBig(validator.VotingPower))
		for _, client := range validator.Clients {
			power = power.Add(power, hexutil.MustDecodeBig(client.VotingPower))
		}
		for _, partner := range validator.Partners {
			power = power.Add(power, hexutil.MustDecodeBig(partner.VotingPower))
		}
	}
	return power
}

func RenderTable(t table.Writer, format string) {
	switch format {
	case "csv":
		t.RenderCSV()
	case "html":
		t.RenderHTML()
	case "markdown":
		t.RenderMarkdown()
	default:
		t.Render()
	}
}
