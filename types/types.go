package types

import (
	"github.com/ethereum/go-ethereum/common"
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

// Client is the type under Clients returned by pos_getStakerInfo
type Client struct {
	Address     common.Address
	Amount      string
	VotingPower string
	QuitEpoch   uint64
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
