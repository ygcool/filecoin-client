package filecoin

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/ygcool/filecoin-client/types"
)

// GasEstimateGasLimit estimates gas used by the message and returns it. It fails if message fails to execute.
func (c *Client) GasEstimateGasLimit(ctx context.Context, message *types.Message, cids []*cid.Cid) (int64, error) {
	var gasLimit int64
	return gasLimit, c.Request(ctx, c.FilecoinMethod("GasEstimateGasLimit"), &gasLimit, message, cids)
}

// GasEstimateFeeCap estimates gas fee cap
func (c *Client) GasEstimateFeeCap(ctx context.Context, message *types.Message, nblocksincl int64, cids []*cid.Cid) (string, error) {
	var feeCap string
	return feeCap, c.Request(ctx, c.FilecoinMethod("GasEstimateFeeCap"), &feeCap, message, nblocksincl, cids)
}

// GasEstimateGasPremium estimates what gas price should be used for a message to have high likelihood of inclusion in nblocksincl epochs.
func (c *Client) GasEstimateGasPremium(ctx context.Context, nblocksincl int64, addr address.Address, gasLimit int64, cids []*cid.Cid) (string, error) {
	var gasPremium string
	return gasPremium, c.Request(ctx, c.FilecoinMethod("GasEstimateGasPremium"), &gasPremium, nblocksincl, addr, gasLimit, cids)
}

// GasEstimateMessageGas estimates gas values for unset message gas fields
func (c *Client) GasEstimateMessageGas(ctx context.Context, message *types.Message, spec *types.MessageSendSpec, cids []*cid.Cid) (*types.Message, error) {
	var msg *types.Message
	return msg, c.Request(ctx, c.FilecoinMethod("GasEstimateMessageGas"), &msg, message, spec, cids)
}
