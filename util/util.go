package util

import (
	"math"
	"math/big"

	"github.com/jedib0t/go-pretty/table"

	"github.com/linki/wanchain-cli/types"
)

func WeiToEth(eth *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(eth), big.NewFloat(math.Pow10(18)))
}

func Weight(value, total *big.Int) *big.Float {
	return big.NewFloat(0).Mul(big.NewFloat(100), big.NewFloat(0).Quo(big.NewFloat(0).SetInt(value), big.NewFloat(0).SetInt(total)))
}

func PowerWeight(validators []types.Validator) *big.Int {
	power := big.NewInt(0)
	for _, validator := range validators {
		power = power.Add(power, validator.Stake())
	}
	return power
}

func TotalAmountValidators(validators []types.Validator) *big.Int {
	amount := big.NewInt(0)
	for _, validator := range validators {
		amount = amount.Add(amount, validator.TotalAmount())
	}
	return amount
}

func TotalAmountClients(validators []types.Client) *big.Int {
	amount := big.NewInt(0)
	for _, validator := range validators {
		amount = amount.Add(amount, validator.TotalAmount())
	}
	return amount
}

func TotalAmountPartners(validators []types.Partner) *big.Int {
	amount := big.NewInt(0)
	for _, validator := range validators {
		amount = amount.Add(amount, validator.TotalAmount())
	}
	return amount
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
