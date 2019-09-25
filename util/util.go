package util

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetCurrentEpochID(ctx context.Context, client *rpc.Client) (uint64, error) {
	var epochID uint64
	if err := client.CallContext(ctx, &epochID, "pos_getEpochID"); err != nil {
		return 0, err
	}
	return epochID, nil
}

func GetCurrentBlockHeight(ctx context.Context, client *rpc.Client) (uint64, error) {
	var blockHeight hexutil.Uint64
	if err := client.CallContext(ctx, &blockHeight, "eth_blockNumber"); err != nil {
		return 0, err
	}
	return hexutil.MustDecodeUint64(blockHeight.String()), nil
}

func WeiToEth(eth *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(eth), big.NewFloat(math.Pow10(18)))
}
