package mayachain

import (
	"fmt"

	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
)

func (h TssHandler) validateV71(ctx cosmos.Context, msg MsgTssPool) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	newMsg, err := NewMsgTssPool(msg.PubKeys, msg.PoolPubKey, msg.KeygenType, msg.Height, msg.Blame, msg.Chains, msg.Signer, msg.KeygenTime)
	if err != nil {
		return fmt.Errorf("fail to recreate MsgTssPool,err: %w", err)
	}
	if msg.ID != newMsg.ID {
		return cosmos.ErrUnknownRequest("invalid tss message")
	}

	churnRetryBlocks := h.mgr.GetConstants().GetInt64Value(constants.ChurnRetryInterval)
	if msg.Height <= ctx.BlockHeight()-churnRetryBlocks {
		return cosmos.ErrUnknownRequest("invalid keygen block")
	}

	keygenBlock, err := h.mgr.Keeper().GetKeygenBlock(ctx, msg.Height)
	if err != nil {
		return fmt.Errorf("fail to get keygen block from data store: %w", err)
	}

	for _, keygen := range keygenBlock.Keygens {
		keyGenMembers := keygen.GetMembers()
		if !msg.GetPubKeys().Equals(keyGenMembers) {
			continue
		}
		// Make sure the keygen type are consistent
		if msg.KeygenType != keygen.Type {
			continue
		}
		var addr cosmos.AccAddress
		for _, member := range keygen.GetMembers() {
			addr, err = member.GetThorAddress()
			if err == nil && addr.Equals(msg.Signer) {
				return h.validateSigner(ctx, msg.Signer)
			}
		}
	}

	return cosmos.ErrUnauthorized("not authorized")
}
