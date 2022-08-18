package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/shopspring/decimal"
	"github.com/ygcool/filecoin-client/types"
)

// ChainGetMessage reads a message referenced by the specified CID from the chain blockstore.
func (c *Client) ChainGetMessage(ctx context.Context, id cid.Cid) (*types.Message, error) {
	var message *types.Message
	return message, c.Request(ctx, c.FilecoinMethod("ChainGetMessage"), &message, id)
}

// ChainGetBlockMessages returns messages stored in the specified block.
func (c *Client) ChainGetBlockMessages(ctx context.Context, id cid.Cid) (*types.BlockMessages, error) {
	var bm *types.BlockMessages
	return bm, c.Request(ctx, c.FilecoinMethod("ChainGetBlockMessages"), &bm, id)
}

// ChainHead returns the current head of the chain.
func (c *Client) ChainHead(ctx context.Context) (*types.TipSet, error) {
	var ts *types.TipSet
	return ts, c.Request(ctx, c.FilecoinMethod("ChainHead"), &ts)
}

// ChainGetTipSetByHeight looks back for a tipset at the specified epoch. If there are no blocks at the specified epoch, a tipset at an earlier epoch will be returned.
func (c *Client) ChainGetTipSetByHeight(ctx context.Context, height int64, tsk types.TipSetKey) (*types.TipSet, error) {
	var ts *types.TipSet
	return ts, c.Request(ctx, c.FilecoinMethod("ChainGetTipSetByHeight"), &ts, height, tsk)
}

// ChainExport returns a stream of bytes with CAR dump of chain data.
func (c *Client) ChainExport(ctx context.Context, tsk types.TipSetKey) ([]byte, error) {
	var result []byte
	return result, c.Request(ctx, c.FilecoinMethod("ChainExport"), &result, tsk)
}

// ChainGetBlock returns the block specified by the given CID.
func (c *Client) ChainGetBlock(ctx context.Context, id cid.Cid) (*types.BlockHeader, error) {
	var bh *types.BlockHeader
	return bh, c.Request(ctx, c.FilecoinMethod("ChainGetBlock"), &bh, id)
}

// ChainGetGenesis returns the genesis tipset.
func (c *Client) ChainGetGenesis(ctx context.Context) (*types.TipSet, error) {
	var ts *types.TipSet
	return ts, c.Request(ctx, c.FilecoinMethod("ChainGetGenesis"), &ts)
}

// ChainGetNode
func (c *Client) ChainGetNode(ctx context.Context, p string) (*types.IpldObject, error) {
	var ipld *types.IpldObject
	return ipld, c.Request(ctx, c.FilecoinMethod("ChainGetNode"), &ipld, p)
}

// ChainGetParentMessages returns messages stored in parent tipset of the specified block.
func (c *Client) ChainGetParentMessages(ctx context.Context, id cid.Cid) ([]types.ParentMessage, error) {
	var msgs []types.ParentMessage
	return msgs, c.Request(ctx, c.FilecoinMethod("ChainGetParentMessages"), &msgs, id)
}

// ChainGetParentReceipts returns receipts for messages in parent tipset of the specified block.
func (c *Client) ChainGetParentReceipts(ctx context.Context, id cid.Cid) ([]*types.MessageReceipt, error) {
	var mrs []*types.MessageReceipt
	return mrs, c.Request(ctx, c.FilecoinMethod("ChainGetParentReceipts"), &mrs, id)
}

// ChainGetPath returns a set of revert/apply operations needed to get from one tipset to another
func (c *Client) ChainGetPath(ctx context.Context, from types.TipSetKey, to types.TipSetKey) (*types.HeadChange, error) {
	var hc *types.HeadChange
	return hc, c.Request(ctx, c.FilecoinMethod("ChainGetPath"), &hc, from, to)
}

// ChainGetRandomnessFromBeacon is used to sample the beacon for randomness.
func (c *Client) ChainGetRandomnessFromBeacon(ctx context.Context, tsk types.TipSetKey, personalization int64, randEpoch int64, entropy []byte) ([]byte, error) {
	var result []byte
	return result, c.Request(ctx, c.FilecoinMethod("ChainGetRandomnessFromBeacon"), &result, tsk, personalization, randEpoch, entropy)
}

// ChainGetRandomnessFromTickets is used to sample the chain for randomness.
func (c *Client) ChainGetRandomnessFromTickets(ctx context.Context, tsk types.TipSetKey, personalization int64, randEpoch int64, entropy []byte) ([]byte, error) {
	var result []byte
	return result, c.Request(ctx, c.FilecoinMethod("ChainGetRandomnessFromTickets"), &result, tsk, personalization, randEpoch, entropy)
}

// ChainGetTipSet returns the tipset specified by the given TipSetKey.
func (c *Client) ChainGetTipSet(ctx context.Context, tsk types.TipSetKey) (*types.TipSet, error) {
	var ts *types.TipSet
	return ts, c.Request(ctx, c.FilecoinMethod("ChainGetTipSet"), &ts, tsk)
}

// ChainHasObj checks if a given CID exists in the chain blockstore.
func (c *Client) ChainHasObj(ctx context.Context, o cid.Cid) (bool, error) {
	var ok bool
	return ok, c.Request(ctx, c.FilecoinMethod("ChainHasObj"), &ok, o)
}

// ChainNotify returns channel with chain head updates. First message is guaranteed to be of len == 1, and type == 'current'.
func (c *Client) ChainNotify() {
	// todo
}

// ChainReadObj reads ipld nodes referenced by the specified CID from chain blockstore and returns raw bytes.
func (c *Client) ChainReadObj(ctx context.Context, obj cid.Cid) ([]byte, error) {
	var result []byte
	return result, c.Request(ctx, c.FilecoinMethod("ChainReadObj"), &result, obj)
}

// ChainSetHead forcefully sets current chain head. Use with caution.
func (c *Client) ChainSetHead(ctx context.Context, tsk types.TipSetKey) error {
	return c.Request(ctx, c.FilecoinMethod("ChainSetHead"), nil, tsk)
}

// ChainStatObj returns statistics about the graph referenced by 'obj'. If 'base' is also specified, then the returned stat will be a diff between the two objects.
func (c *Client) ChainStatObj(ctx context.Context, obj, base cid.Cid) (types.ObjStat, error) {
	var os types.ObjStat
	return os, c.Request(ctx, c.FilecoinMethod("ChainStatObj"), &os, obj, base)
}

// ChainTipSetWeight computes weight for the specified tipset.
func (c *Client) ChainTipSetWeight(ctx context.Context, tsk types.TipSetKey) (decimal.Decimal, error) {
	var d decimal.Decimal
	return d, c.Request(ctx, c.FilecoinMethod("ChainTipSetWeight"), &d, tsk)
}
