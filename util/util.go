package util

import (
	"math"
	"math/big"

	"github.com/jedib0t/go-pretty/table"
)

func WeiToEth(eth *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(eth), big.NewFloat(math.Pow10(18)))
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
