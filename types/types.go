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
