package types

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
