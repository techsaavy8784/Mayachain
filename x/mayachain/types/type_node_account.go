package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/blang/semver"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
)

// Valid check whether the node status is valid or not
func (x NodeStatus) Valid() error {
	caser := cases.Title(language.Und)
	if _, ok := NodeStatus_value[caser.String(x.String())]; !ok {
		return fmt.Errorf("invalid node status")
	}
	return nil
}

// MarshalJSON marshal NodeStatus to JSON in string form
func (x NodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON convert string form back to NodeStatus
func (x *NodeStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*x = getNodeStatus(s)
	return nil
}

// getNodeStatus from string
func getNodeStatus(ps string) NodeStatus {
	caser := cases.Title(language.Und)
	if val, ok := NodeStatus_value[caser.String(ps)]; ok {
		return NodeStatus(val)
	}
	return NodeStatus_Unknown
}

func NewBondProviders(acc cosmos.AccAddress) BondProviders {
	return BondProviders{
		NodeAddress:     acc,
		NodeOperatorFee: cosmos.ZeroUint(),
		Providers:       make([]BondProvider, 0),
	}
}

func NewBondProvider(acc cosmos.AccAddress) BondProvider {
	return BondProvider{
		BondAddress: acc,
		Bonded:      false,
		Reward:      nil,
	}
}

// NewNodeAccount create new instance of NodeAccount
func NewNodeAccount(nodeAddress cosmos.AccAddress, status NodeStatus, nodePubKeySet common.PubKeySet, validatorConsPubKey string, aztecAddress common.Address, bond cosmos.Uint, bondAddress common.Address, height int64) NodeAccount {
	na := NodeAccount{
		NodeAddress:         nodeAddress,
		PubKeySet:           nodePubKeySet,
		ValidatorConsPubKey: validatorConsPubKey,
		AztecAddress:        aztecAddress,
		Bond:                bond,
		BondAddress:         bondAddress,
		Reward:              cosmos.ZeroUint(),
	}
	na.UpdateStatus(status, height)
	return na
}

// IsEmpty decide whether NodeAccount is empty
func (m *NodeAccount) IsEmpty() bool {
	return m.NodeAddress.Empty() || m.Status == NodeStatus_Unknown
}

// Valid check whether NodeAccount has all necessary values
func (m *NodeAccount) Valid() error {
	if m.NodeAddress.Empty() {
		return errors.New("node thor address is empty")
	}
	if m.BondAddress.IsEmpty() {
		return errors.New("bond address is empty")
	}
	if m.Status == NodeStatus_Unknown {
		return errors.New("node status cannot be unknown")
	}

	return nil
}

// GetSignerMembership return a list of pubkey that the node are part of
func (m *NodeAccount) GetSignerMembership() common.PubKeys {
	pubkeys := make(common.PubKeys, 0)
	for _, pk := range m.SignerMembership {
		pk, err := common.NewPubKey(pk)
		if err != nil {
			continue
		}
		pubkeys = append(pubkeys, pk)
	}
	return pubkeys
}

// GetVersion return the node account's version
func (m *NodeAccount) GetVersion() semver.Version {
	version, _ := semver.Make(m.Version)
	return version
}

// UpdateStatus change the status of node account, in the mean time update StatusSince field
func (m *NodeAccount) UpdateStatus(status NodeStatus, height int64) {
	if m.Status == status {
		return
	}
	m.Status = status
	m.StatusSince = height
	if status != NodeStatus_Active {
		m.ActiveBlockHeight = 0
	}
}

// Equals compare two node account, to see whether they are equal
func (m *NodeAccount) Equals(n1 NodeAccount) bool {
	if m.NodeAddress.Equals(n1.NodeAddress) &&
		m.PubKeySet.Equals(n1.PubKeySet) &&
		m.ValidatorConsPubKey == n1.ValidatorConsPubKey &&
		m.BondAddress.Equals(n1.BondAddress) &&
		m.Reward.Equal(n1.Reward) &&
		m.GetVersion().Equals(n1.GetVersion()) {
		return true
	}
	return false
}

// String implement fmt.Stringer interface
func (m *NodeAccount) String() string {
	sb := strings.Builder{}
	sb.WriteString("node:" + m.NodeAddress.String() + "\n")
	sb.WriteString("status:" + m.Status.String() + "\n")
	sb.WriteString("node pubkeys:" + m.PubKeySet.String() + "\n")
	sb.WriteString("validator consensus pub key:" + m.ValidatorConsPubKey + "\n")
	sb.WriteString("reward:" + m.Reward.String() + "\n")
	sb.WriteString("version:" + m.Version + "\n")
	sb.WriteString("bond address:" + m.BondAddress.String() + "\n")
	sb.WriteString("requested to leave:" + strconv.FormatBool(m.RequestedToLeave) + "\n")
	sb.WriteString("aztec address" + m.AztecAddress.String() + "\n")
	return sb.String()
}

// CalcBondUnits calculate bond
func (m *NodeAccount) CalcBondUnits(height, slashpoints int64) cosmos.Uint {
	// ensure slashpoints is not negative
	slashpoints = int64(math.Max(float64(0), float64(slashpoints)))
	if height < 0 || m.ActiveBlockHeight < 0 || slashpoints < 0 {
		return cosmos.ZeroUint()
	}

	blockCount := height - (m.ActiveBlockHeight + slashpoints)
	if blockCount < 0 { // ensure we're never negative
		blockCount = 0
	}

	return cosmos.NewUint(uint64(blockCount))
}

// TryAddSignerPubKey add a key to node account
func (m *NodeAccount) TryAddSignerPubKey(key common.PubKey) {
	if key.IsEmpty() {
		return
	}
	for _, item := range m.GetSignerMembership() {
		if item.Equals(key) {
			return
		}
	}
	m.SignerMembership = append(m.SignerMembership, key.String())
}

// TryRemoveSignerPubKey remove the given pubkey from signer membership
func (m *NodeAccount) TryRemoveSignerPubKey(key common.PubKey) {
	if key.IsEmpty() {
		return
	}
	idxToDelete := -1
	for idx, item := range m.GetSignerMembership() {
		if item.Equals(key) {
			idxToDelete = idx
		}
	}
	if idxToDelete != -1 {
		m.SignerMembership = append(m.SignerMembership[:idxToDelete], m.SignerMembership[idxToDelete+1:]...)
	}
}

// NodeAccounts just a list of NodeAccount
type NodeAccounts []NodeAccount

// IsEmpty to check whether the NodeAccounts is empty
func (nas NodeAccounts) IsEmpty() bool {
	return len(nas) == 0
}

// IsNodeKeys validate whether the given account address belongs to an currently active validator
func (nas NodeAccounts) IsNodeKeys(addr cosmos.AccAddress) bool {
	for _, na := range nas {
		if na.Status == NodeStatus_Active && addr.Equals(na.NodeAddress) {
			return true
		}
	}
	return false
}

// Less sort interface , it will sort by StatusSince field, and then by SignerBNBAddress
func (nas NodeAccounts) Less(i, j int) bool {
	if nas[i].StatusSince < nas[j].StatusSince {
		return true
	}
	if nas[i].StatusSince > nas[j].StatusSince {
		return false
	}
	return nas[i].NodeAddress.String() < nas[j].NodeAddress.String()
}

// Len return the number of accounts in it
func (nas NodeAccounts) Len() int { return len(nas) }

// Swap node account
func (nas NodeAccounts) Swap(i, j int) {
	nas[i], nas[j] = nas[j], nas[i]
}

// Contains will check whether the given node account is in the list
func (nas NodeAccounts) Contains(na NodeAccount) bool {
	for _, item := range nas {
		if item.Equals(na) {
			return true
		}
	}
	return false
}

func (nas NodeAccounts) GetNodeAddresses() []cosmos.AccAddress {
	addrs := make([]cosmos.AccAddress, len(nas))
	for i, na := range nas {
		addrs[i] = na.NodeAddress
	}
	return addrs
}

func (m *BondProvider) IsEmpty() bool {
	return m.BondAddress.Empty()
}

func (m *BondProvider) HasBonded() bool {
	return m.Bonded
}

func (m *BondProvider) HasRewards() bool {
	if m.Reward != nil && m.Reward.GT(cosmos.ZeroUint()) {
		return true
	}
	return false
}

func (bp *BondProviders) Has(acc cosmos.AccAddress) bool {
	provider := bp.Get(acc)
	return !provider.IsEmpty()
}

func (bp *BondProviders) Get(acc cosmos.AccAddress) BondProvider {
	for _, provider := range bp.Providers {
		if provider.BondAddress.Equals(acc) {
			return provider
		}
	}
	return BondProvider{}
}

func (bp *BondProviders) BondLiquidity(acc cosmos.AccAddress) {
	for i, provider := range bp.Providers {
		if provider.BondAddress.Equals(acc) {
			bp.Providers[i].Bonded = true
			return
		}
	}
}

func (bp *BondProviders) Unbond(acc cosmos.AccAddress) {
	for i, provider := range bp.Providers {
		if provider.BondAddress.Equals(acc) {
			bp.Providers[i].Bonded = false
			return
		}
	}
}

func (bp *BondProviders) HasRewards(acc cosmos.AccAddress) bool {
	for i, provider := range bp.Providers {
		if provider.BondAddress.Equals(acc) && bp.Providers[i].Reward != nil && bp.Providers[i].Reward.GT(cosmos.ZeroUint()) {
			return true
		}
	}
	return false
}

func (bp *BondProviders) RemoveRewards(acc cosmos.AccAddress) {
	for i, provider := range bp.Providers {
		if provider.BondAddress.Equals(acc) {
			zero := cosmos.ZeroUint()
			bp.Providers[i].Reward = &zero
			return
		}
	}
}

// remove provider (only if bond is zero)
func (bp *BondProviders) Remove(acc cosmos.AccAddress) bool {
	for i, provider := range bp.Providers {
		if i == 0 {
			// cannot remove the first bond provider
			continue
		}
		if provider.BondAddress.Equals(acc) && !provider.Bonded {
			bp.Providers = append(bp.Providers[:i], bp.Providers[i+1:]...)
			return true
		}
	}
	return false
}

// HasProviderBonded: Checks if a bond provider (not the operator) has bonded to the node
func (bp *BondProviders) HasProviderBonded(opBondAddress cosmos.AccAddress) bool {
	for i := range bp.Providers {
		if !bp.Providers[i].BondAddress.Equals(opBondAddress) && bp.Providers[i].Bonded {
			return true
		}
	}

	return false
}