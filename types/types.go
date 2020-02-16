package types

import (
	"math/big"

	"github.com/wanchain/go-wanchain/common"
	"github.com/wanchain/go-wanchain/common/hexutil"
)

// Activity is the partial type returned by pos_getActivity
type Activity struct {
	EPActivity []uint64
	EPLeader   []string
	RPActivity []uint64
	RPLeader   []string
	SLBlocks   []uint64
	SLTLeader  []string
}

// DelegatorIncentive is the partial type returned by pos_getEpochIncentivePayDetail
type DelegatorIncentive struct {
	Address   string
	Incentive string
	Type      string
}

// ValidatorIncentive is the type returned by pos_getEpochIncentivePayDetail
type ValidatorIncentive struct {
	DelegatorIncentive
	Delegators      []DelegatorIncentive
	StakeInFromAddr string
}

// Validator is the type returned by pos_getStakerInfo
type Validator struct {
	Address             common.Address
	PubSec256           string
	PubBn256            string
	Amount              string
	VotingPower         string
	LockEpochs          uint64
	NextLockEpochs      uint64
	From                common.Address
	StakingEpoch        uint64
	FeeRate             uint64
	Clients             []Client
	Partners            []Partner
	MaxFeeRate          uint64
	FeeRateChangedEpoch uint64
}

func (v *Validator) Stake() *big.Int {
	power := hexutil.MustDecodeBig(v.VotingPower)

	for _, client := range v.Clients {
		power = power.Add(power, client.Stake())
	}

	for _, partner := range v.Partners {
		power = power.Add(power, partner.Stake())
	}

	return power
}

func (v *Validator) TotalAmount() *big.Int {
	amount := hexutil.MustDecodeBig(v.Amount)

	for _, client := range v.Clients {
		amount = amount.Add(amount, client.TotalAmount())
	}

	for _, partner := range v.Partners {
		amount = amount.Add(amount, partner.TotalAmount())
	}

	return amount
}

// Client is the type under Clients returned by pos_getStakerInfo
type Client struct {
	Address     common.Address
	Amount      string
	VotingPower string
	QuitEpoch   uint64
}

func (c *Client) Stake() *big.Int {
	return hexutil.MustDecodeBig(c.VotingPower)
}

func (c *Client) TotalAmount() *big.Int {
	return hexutil.MustDecodeBig(c.Amount)
}

// Partner is the type under Partners returned by pos_getStakerInfo
type Partner struct {
	Address      common.Address
	Amount       string
	VotingPower  string
	Renewal      bool
	LockEpochs   uint64
	StakingEpoch uint64
}

func (p *Partner) Stake() *big.Int {
	return hexutil.MustDecodeBig(p.VotingPower)
}

func (p *Partner) TotalAmount() *big.Int {
	return hexutil.MustDecodeBig(p.Amount)
}
