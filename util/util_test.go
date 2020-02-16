package util

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/suite"
)

type UtilSuite struct {
	suite.Suite
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(UtilSuite))
}

func (suite *UtilSuite) TestWeiToEth() {
	for _, tc := range []struct {
		wei *big.Int
		eth *big.Float
	}{
		{common.Big0, big.NewFloat(0)},                             // 0 wei => 0 eth
		{common.Big1, big.NewFloat(0.000000000000000001)},          // 1 wei => 0.000000000000000001 eth
		{big.NewInt(1_500_000_000_000_000_000), big.NewFloat(1.5)}, // 1.5*10**18 wei => 1.5 eth
	} {
		actual, _ := WeiToEth(tc.wei).Float64()
		expected, _ := tc.eth.Float64()

		suite.Equal(expected, actual)
	}
}

func (suite *UtilSuite) TestWeight() {
	for _, tc := range []struct {
		value, total *big.Int
		weight       *big.Float
	}{
		{common.Big0, common.Big0, big.NewFloat(0)},           // 0 of 0 => 0%
		{common.Big0, big.NewInt(100), big.NewFloat(0)},       // 0 of 100 => 0%
		{big.NewInt(100), big.NewInt(100), big.NewFloat(100)}, // 100 of 100 => 100%
		{big.NewInt(20), big.NewInt(1000), big.NewFloat(2)},   // 20 of 1000 => 2%
	} {
		actual, _ := Weight(tc.value, tc.total).Float64()
		expected, _ := tc.weight.Float64()

		suite.Equal(expected, actual)
	}
}
