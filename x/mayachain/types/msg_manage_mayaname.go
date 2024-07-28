package types

import (
	"github.com/blang/semver"
	"gitlab.com/mayachain/mayanode/common"
	cosmos "gitlab.com/mayachain/mayanode/common/cosmos"
)

// NewMsgManageMAYAName create a new instance of MsgManageMAYAName
func NewMsgManageMAYAName(name string, chain common.Chain, addr common.Address, coin common.Coin, exp int64, asset common.Asset, owner, signer cosmos.AccAddress, affiliateSplit int64, subAffiliateSplit int64) *MsgManageMAYAName {
	return &MsgManageMAYAName{
		Name:              name,
		Chain:             chain,
		Address:           addr,
		Coin:              coin,
		ExpireBlockHeight: exp,
		PreferredAsset:    asset,
		Owner:             owner,
		Signer:            signer,
		AffiliateSplit:    affiliateSplit,
		SubAffiliateSplit: subAffiliateSplit,
	}
}

// Route should return the Route of the module
func (m *MsgManageMAYAName) Route() string { return RouterKey }

// Type should return the action
func (m MsgManageMAYAName) Type() string { return "manage_mayaname" }

// ValidateBasic runs stateless checks on the message
func (m *MsgManageMAYAName) ValidateBasic() error {
	// validate n
	if m.Signer.Empty() {
		return cosmos.ErrInvalidAddress(m.Signer.String())
	}
	if m.Chain.IsEmpty() {
		return cosmos.ErrUnknownRequest("chain can't be empty")
	}
	if m.Address.IsEmpty() {
		return cosmos.ErrUnknownRequest("address can't be empty")
	}
	if !m.Address.IsChain(m.Chain, semver.Version{}) {
		return cosmos.ErrUnknownRequest("address and chain must match")
	}
	if !m.Coin.Asset.IsNativeBase() {
		return cosmos.ErrUnknownRequest("coin must be native rune")
	}
	if m.AffiliateSplit < 0 {
		return cosmos.ErrUnknownRequest("affiliate_split cannot be negative")
	}
	if m.SubAffiliateSplit < 0 {
		return cosmos.ErrUnknownRequest("sub_affiliate_split cannot be negative")
	}
	return nil
}

// ValidateBasicV108 runs stateless checks on the message
func (m *MsgManageMAYAName) ValidateBasicV108(version semver.Version) error {
	// validate n
	if m.Signer.Empty() {
		return cosmos.ErrInvalidAddress(m.Signer.String())
	}
	if m.Chain.IsEmpty() {
		return cosmos.ErrUnknownRequest("chain can't be empty")
	}
	if m.Address.IsEmpty() {
		return cosmos.ErrUnknownRequest("address can't be empty")
	}
	if !m.Address.IsChain(m.Chain, version) {
		return cosmos.ErrUnknownRequest("address and chain must match")
	}
	if !m.Coin.Asset.IsNativeBase() {
		return cosmos.ErrUnknownRequest("coin must be native rune")
	}
	if m.AffiliateSplit < 0 {
		return cosmos.ErrUnknownRequest("affiliate_split cannot be negative")
	}
	if m.SubAffiliateSplit < 0 {
		return cosmos.ErrUnknownRequest("sub_affiliate_split cannot be negative")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (m *MsgManageMAYAName) GetSignBytes() []byte {
	return cosmos.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// GetSigners defines whose signature is required
func (m *MsgManageMAYAName) GetSigners() []cosmos.AccAddress {
	return []cosmos.AccAddress{m.Signer}
}
