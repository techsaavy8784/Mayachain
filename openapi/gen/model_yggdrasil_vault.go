/*
Mayanode API

Mayanode REST API.

Contact: devs@mayachain.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// YggdrasilVault struct for YggdrasilVault
type YggdrasilVault struct {
	BlockHeight *int64 `json:"block_height,omitempty"`
	PubKey *string `json:"pub_key,omitempty"`
	Coins []Coin `json:"coins"`
	Type *string `json:"type,omitempty"`
	StatusSince *int64 `json:"status_since,omitempty"`
	// the list of node public keys which are members of the vault
	Membership []string `json:"membership,omitempty"`
	Chains []string `json:"chains,omitempty"`
	InboundTxCount *int64 `json:"inbound_tx_count,omitempty"`
	OutboundTxCount *int64 `json:"outbound_tx_count,omitempty"`
	PendingTxBlockHeights []int64 `json:"pending_tx_block_heights,omitempty"`
	Routers []VaultRouter `json:"routers"`
	Status string `json:"status"`
	// current node bond
	Bond string `json:"bond"`
	// value in cacao of the vault's assets
	TotalValue string `json:"total_value"`
	Addresses []VaultAddress `json:"addresses"`
}

// NewYggdrasilVault instantiates a new YggdrasilVault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYggdrasilVault(coins []Coin, routers []VaultRouter, status string, bond string, totalValue string, addresses []VaultAddress) *YggdrasilVault {
	this := YggdrasilVault{}
	this.Coins = coins
	this.Routers = routers
	this.Status = status
	this.Bond = bond
	this.TotalValue = totalValue
	this.Addresses = addresses
	return &this
}

// NewYggdrasilVaultWithDefaults instantiates a new YggdrasilVault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYggdrasilVaultWithDefaults() *YggdrasilVault {
	this := YggdrasilVault{}
	return &this
}

// GetBlockHeight returns the BlockHeight field value if set, zero value otherwise.
func (o *YggdrasilVault) GetBlockHeight() int64 {
	if o == nil || o.BlockHeight == nil {
		var ret int64
		return ret
	}
	return *o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetBlockHeightOk() (*int64, bool) {
	if o == nil || o.BlockHeight == nil {
		return nil, false
	}
	return o.BlockHeight, true
}

// HasBlockHeight returns a boolean if a field has been set.
func (o *YggdrasilVault) HasBlockHeight() bool {
	if o != nil && o.BlockHeight != nil {
		return true
	}

	return false
}

// SetBlockHeight gets a reference to the given int64 and assigns it to the BlockHeight field.
func (o *YggdrasilVault) SetBlockHeight(v int64) {
	o.BlockHeight = &v
}

// GetPubKey returns the PubKey field value if set, zero value otherwise.
func (o *YggdrasilVault) GetPubKey() string {
	if o == nil || o.PubKey == nil {
		var ret string
		return ret
	}
	return *o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetPubKeyOk() (*string, bool) {
	if o == nil || o.PubKey == nil {
		return nil, false
	}
	return o.PubKey, true
}

// HasPubKey returns a boolean if a field has been set.
func (o *YggdrasilVault) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

// SetPubKey gets a reference to the given string and assigns it to the PubKey field.
func (o *YggdrasilVault) SetPubKey(v string) {
	o.PubKey = &v
}

// GetCoins returns the Coins field value
func (o *YggdrasilVault) GetCoins() []Coin {
	if o == nil {
		var ret []Coin
		return ret
	}

	return o.Coins
}

// GetCoinsOk returns a tuple with the Coins field value
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetCoinsOk() ([]Coin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coins, true
}

// SetCoins sets field value
func (o *YggdrasilVault) SetCoins(v []Coin) {
	o.Coins = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *YggdrasilVault) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *YggdrasilVault) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *YggdrasilVault) SetType(v string) {
	o.Type = &v
}

// GetStatusSince returns the StatusSince field value if set, zero value otherwise.
func (o *YggdrasilVault) GetStatusSince() int64 {
	if o == nil || o.StatusSince == nil {
		var ret int64
		return ret
	}
	return *o.StatusSince
}

// GetStatusSinceOk returns a tuple with the StatusSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetStatusSinceOk() (*int64, bool) {
	if o == nil || o.StatusSince == nil {
		return nil, false
	}
	return o.StatusSince, true
}

// HasStatusSince returns a boolean if a field has been set.
func (o *YggdrasilVault) HasStatusSince() bool {
	if o != nil && o.StatusSince != nil {
		return true
	}

	return false
}

// SetStatusSince gets a reference to the given int64 and assigns it to the StatusSince field.
func (o *YggdrasilVault) SetStatusSince(v int64) {
	o.StatusSince = &v
}

// GetMembership returns the Membership field value if set, zero value otherwise.
func (o *YggdrasilVault) GetMembership() []string {
	if o == nil || o.Membership == nil {
		var ret []string
		return ret
	}
	return o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetMembershipOk() ([]string, bool) {
	if o == nil || o.Membership == nil {
		return nil, false
	}
	return o.Membership, true
}

// HasMembership returns a boolean if a field has been set.
func (o *YggdrasilVault) HasMembership() bool {
	if o != nil && o.Membership != nil {
		return true
	}

	return false
}

// SetMembership gets a reference to the given []string and assigns it to the Membership field.
func (o *YggdrasilVault) SetMembership(v []string) {
	o.Membership = v
}

// GetChains returns the Chains field value if set, zero value otherwise.
func (o *YggdrasilVault) GetChains() []string {
	if o == nil || o.Chains == nil {
		var ret []string
		return ret
	}
	return o.Chains
}

// GetChainsOk returns a tuple with the Chains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetChainsOk() ([]string, bool) {
	if o == nil || o.Chains == nil {
		return nil, false
	}
	return o.Chains, true
}

// HasChains returns a boolean if a field has been set.
func (o *YggdrasilVault) HasChains() bool {
	if o != nil && o.Chains != nil {
		return true
	}

	return false
}

// SetChains gets a reference to the given []string and assigns it to the Chains field.
func (o *YggdrasilVault) SetChains(v []string) {
	o.Chains = v
}

// GetInboundTxCount returns the InboundTxCount field value if set, zero value otherwise.
func (o *YggdrasilVault) GetInboundTxCount() int64 {
	if o == nil || o.InboundTxCount == nil {
		var ret int64
		return ret
	}
	return *o.InboundTxCount
}

// GetInboundTxCountOk returns a tuple with the InboundTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetInboundTxCountOk() (*int64, bool) {
	if o == nil || o.InboundTxCount == nil {
		return nil, false
	}
	return o.InboundTxCount, true
}

// HasInboundTxCount returns a boolean if a field has been set.
func (o *YggdrasilVault) HasInboundTxCount() bool {
	if o != nil && o.InboundTxCount != nil {
		return true
	}

	return false
}

// SetInboundTxCount gets a reference to the given int64 and assigns it to the InboundTxCount field.
func (o *YggdrasilVault) SetInboundTxCount(v int64) {
	o.InboundTxCount = &v
}

// GetOutboundTxCount returns the OutboundTxCount field value if set, zero value otherwise.
func (o *YggdrasilVault) GetOutboundTxCount() int64 {
	if o == nil || o.OutboundTxCount == nil {
		var ret int64
		return ret
	}
	return *o.OutboundTxCount
}

// GetOutboundTxCountOk returns a tuple with the OutboundTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetOutboundTxCountOk() (*int64, bool) {
	if o == nil || o.OutboundTxCount == nil {
		return nil, false
	}
	return o.OutboundTxCount, true
}

// HasOutboundTxCount returns a boolean if a field has been set.
func (o *YggdrasilVault) HasOutboundTxCount() bool {
	if o != nil && o.OutboundTxCount != nil {
		return true
	}

	return false
}

// SetOutboundTxCount gets a reference to the given int64 and assigns it to the OutboundTxCount field.
func (o *YggdrasilVault) SetOutboundTxCount(v int64) {
	o.OutboundTxCount = &v
}

// GetPendingTxBlockHeights returns the PendingTxBlockHeights field value if set, zero value otherwise.
func (o *YggdrasilVault) GetPendingTxBlockHeights() []int64 {
	if o == nil || o.PendingTxBlockHeights == nil {
		var ret []int64
		return ret
	}
	return o.PendingTxBlockHeights
}

// GetPendingTxBlockHeightsOk returns a tuple with the PendingTxBlockHeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetPendingTxBlockHeightsOk() ([]int64, bool) {
	if o == nil || o.PendingTxBlockHeights == nil {
		return nil, false
	}
	return o.PendingTxBlockHeights, true
}

// HasPendingTxBlockHeights returns a boolean if a field has been set.
func (o *YggdrasilVault) HasPendingTxBlockHeights() bool {
	if o != nil && o.PendingTxBlockHeights != nil {
		return true
	}

	return false
}

// SetPendingTxBlockHeights gets a reference to the given []int64 and assigns it to the PendingTxBlockHeights field.
func (o *YggdrasilVault) SetPendingTxBlockHeights(v []int64) {
	o.PendingTxBlockHeights = v
}

// GetRouters returns the Routers field value
func (o *YggdrasilVault) GetRouters() []VaultRouter {
	if o == nil {
		var ret []VaultRouter
		return ret
	}

	return o.Routers
}

// GetRoutersOk returns a tuple with the Routers field value
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetRoutersOk() ([]VaultRouter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Routers, true
}

// SetRouters sets field value
func (o *YggdrasilVault) SetRouters(v []VaultRouter) {
	o.Routers = v
}

// GetStatus returns the Status field value
func (o *YggdrasilVault) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *YggdrasilVault) SetStatus(v string) {
	o.Status = v
}

// GetBond returns the Bond field value
func (o *YggdrasilVault) GetBond() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bond
}

// GetBondOk returns a tuple with the Bond field value
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetBondOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bond, true
}

// SetBond sets field value
func (o *YggdrasilVault) SetBond(v string) {
	o.Bond = v
}

// GetTotalValue returns the TotalValue field value
func (o *YggdrasilVault) GetTotalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalValue
}

// GetTotalValueOk returns a tuple with the TotalValue field value
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetTotalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalValue, true
}

// SetTotalValue sets field value
func (o *YggdrasilVault) SetTotalValue(v string) {
	o.TotalValue = v
}

// GetAddresses returns the Addresses field value
func (o *YggdrasilVault) GetAddresses() []VaultAddress {
	if o == nil {
		var ret []VaultAddress
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *YggdrasilVault) GetAddressesOk() ([]VaultAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *YggdrasilVault) SetAddresses(v []VaultAddress) {
	o.Addresses = v
}

func (o YggdrasilVault) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockHeight != nil {
		toSerialize["block_height"] = o.BlockHeight
	}
	if o.PubKey != nil {
		toSerialize["pub_key"] = o.PubKey
	}
	if true {
		toSerialize["coins"] = o.Coins
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.StatusSince != nil {
		toSerialize["status_since"] = o.StatusSince
	}
	if o.Membership != nil {
		toSerialize["membership"] = o.Membership
	}
	if o.Chains != nil {
		toSerialize["chains"] = o.Chains
	}
	if o.InboundTxCount != nil {
		toSerialize["inbound_tx_count"] = o.InboundTxCount
	}
	if o.OutboundTxCount != nil {
		toSerialize["outbound_tx_count"] = o.OutboundTxCount
	}
	if o.PendingTxBlockHeights != nil {
		toSerialize["pending_tx_block_heights"] = o.PendingTxBlockHeights
	}
	if true {
		toSerialize["routers"] = o.Routers
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["bond"] = o.Bond
	}
	if true {
		toSerialize["total_value"] = o.TotalValue
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableYggdrasilVault struct {
	value *YggdrasilVault
	isSet bool
}

func (v NullableYggdrasilVault) Get() *YggdrasilVault {
	return v.value
}

func (v *NullableYggdrasilVault) Set(val *YggdrasilVault) {
	v.value = val
	v.isSet = true
}

func (v NullableYggdrasilVault) IsSet() bool {
	return v.isSet
}

func (v *NullableYggdrasilVault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYggdrasilVault(val *YggdrasilVault) *NullableYggdrasilVault {
	return &NullableYggdrasilVault{value: val, isSet: true}
}

func (v NullableYggdrasilVault) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYggdrasilVault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

