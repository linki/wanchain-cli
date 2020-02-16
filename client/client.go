package client

import (
	"github.com/wanchain/go-wanchain/common/hexutil"
	"github.com/wanchain/go-wanchain/rpc"

	"github.com/linki/wanchain-cli/types"
)

type Client struct {
	client *rpc.Client
}

func NewClient(url string) (*Client, error) {
	client, err := rpc.Dial(url)
	return &Client{client: client}, err
}

func (c *Client) Close() {
	c.client.Close()
}

func (c *Client) GetCurrentBlockHeight() (uint64, error) {
	var blockHeight hexutil.Uint64
	return uint64(blockHeight), c.client.Call(&blockHeight, "eth_blockNumber")
}

func (c *Client) GetCurrentEpochID() (epochID uint64, _ error) {
	return epochID, c.client.Call(&epochID, "pos_getEpochID")
}

func (c *Client) GetCurrentSlotID() (slotID uint64, _ error) {
	return slotID, c.client.Call(&slotID, "pos_getSlotID")
}

func (c *Client) GetActivity(epochID uint64) (activity types.Activity, _ error) {
	return activity, c.client.Call(&activity, "pos_getActivity", epochID)
}

func (c *Client) GetIncentives(epochID uint64) (incentives []types.ValidatorIncentive, _ error) {
	return incentives, c.client.Call(&incentives, "pos_getEpochIncentivePayDetail", epochID)
}

func (c *Client) GetEpochLeaders(epochID uint64) (epochLeaders []string, _ error) {
	return epochLeaders, c.client.Call(&epochLeaders, "pos_getEpochLeadersAddrByEpochID", epochID)
}

func (c *Client) GetRandomProposers(epochID uint64) (randomProposers []string, _ error) {
	return randomProposers, c.client.Call(&randomProposers, "pos_getRandomProposersAddrByEpochID", epochID)
}

func (c *Client) GetValidators(blockHeight uint64) (validators []types.Validator, _ error) {
	return validators, c.client.Call(&validators, "pos_getStakerInfo", blockHeight)
}
