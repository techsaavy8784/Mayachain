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

// TxOutItem struct for TxOutItem
type TxOutItem struct {
	Chain string `json:"chain"`
	ToAddress string `json:"to_address"`
	VaultPubKey *string `json:"vault_pub_key,omitempty"`
	Coin Coin `json:"coin"`
	Memo *string `json:"memo,omitempty"`
	MaxGas []Coin `json:"max_gas"`
	GasRate *int64 `json:"gas_rate,omitempty"`
	InHash *string `json:"in_hash,omitempty"`
	OutHash *string `json:"out_hash,omitempty"`
	Height *int64 `json:"height,omitempty"`
}

// NewTxOutItem instantiates a new TxOutItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxOutItem(chain string, toAddress string, coin Coin, maxGas []Coin) *TxOutItem {
	this := TxOutItem{}
	this.Chain = chain
	this.ToAddress = toAddress
	this.Coin = coin
	this.MaxGas = maxGas
	return &this
}

// NewTxOutItemWithDefaults instantiates a new TxOutItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxOutItemWithDefaults() *TxOutItem {
	this := TxOutItem{}
	return &this
}

// GetChain returns the Chain field value
func (o *TxOutItem) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *TxOutItem) SetChain(v string) {
	o.Chain = v
}

// GetToAddress returns the ToAddress field value
func (o *TxOutItem) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *TxOutItem) SetToAddress(v string) {
	o.ToAddress = v
}

// GetVaultPubKey returns the VaultPubKey field value if set, zero value otherwise.
func (o *TxOutItem) GetVaultPubKey() string {
	if o == nil || o.VaultPubKey == nil {
		var ret string
		return ret
	}
	return *o.VaultPubKey
}

// GetVaultPubKeyOk returns a tuple with the VaultPubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetVaultPubKeyOk() (*string, bool) {
	if o == nil || o.VaultPubKey == nil {
		return nil, false
	}
	return o.VaultPubKey, true
}

// HasVaultPubKey returns a boolean if a field has been set.
func (o *TxOutItem) HasVaultPubKey() bool {
	if o != nil && o.VaultPubKey != nil {
		return true
	}

	return false
}

// SetVaultPubKey gets a reference to the given string and assigns it to the VaultPubKey field.
func (o *TxOutItem) SetVaultPubKey(v string) {
	o.VaultPubKey = &v
}

// GetCoin returns the Coin field value
func (o *TxOutItem) GetCoin() Coin {
	if o == nil {
		var ret Coin
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetCoinOk() (*Coin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *TxOutItem) SetCoin(v Coin) {
	o.Coin = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *TxOutItem) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *TxOutItem) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *TxOutItem) SetMemo(v string) {
	o.Memo = &v
}

// GetMaxGas returns the MaxGas field value
func (o *TxOutItem) GetMaxGas() []Coin {
	if o == nil {
		var ret []Coin
		return ret
	}

	return o.MaxGas
}

// GetMaxGasOk returns a tuple with the MaxGas field value
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetMaxGasOk() ([]Coin, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxGas, true
}

// SetMaxGas sets field value
func (o *TxOutItem) SetMaxGas(v []Coin) {
	o.MaxGas = v
}

// GetGasRate returns the GasRate field value if set, zero value otherwise.
func (o *TxOutItem) GetGasRate() int64 {
	if o == nil || o.GasRate == nil {
		var ret int64
		return ret
	}
	return *o.GasRate
}

// GetGasRateOk returns a tuple with the GasRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetGasRateOk() (*int64, bool) {
	if o == nil || o.GasRate == nil {
		return nil, false
	}
	return o.GasRate, true
}

// HasGasRate returns a boolean if a field has been set.
func (o *TxOutItem) HasGasRate() bool {
	if o != nil && o.GasRate != nil {
		return true
	}

	return false
}

// SetGasRate gets a reference to the given int64 and assigns it to the GasRate field.
func (o *TxOutItem) SetGasRate(v int64) {
	o.GasRate = &v
}

// GetInHash returns the InHash field value if set, zero value otherwise.
func (o *TxOutItem) GetInHash() string {
	if o == nil || o.InHash == nil {
		var ret string
		return ret
	}
	return *o.InHash
}

// GetInHashOk returns a tuple with the InHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetInHashOk() (*string, bool) {
	if o == nil || o.InHash == nil {
		return nil, false
	}
	return o.InHash, true
}

// HasInHash returns a boolean if a field has been set.
func (o *TxOutItem) HasInHash() bool {
	if o != nil && o.InHash != nil {
		return true
	}

	return false
}

// SetInHash gets a reference to the given string and assigns it to the InHash field.
func (o *TxOutItem) SetInHash(v string) {
	o.InHash = &v
}

// GetOutHash returns the OutHash field value if set, zero value otherwise.
func (o *TxOutItem) GetOutHash() string {
	if o == nil || o.OutHash == nil {
		var ret string
		return ret
	}
	return *o.OutHash
}

// GetOutHashOk returns a tuple with the OutHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetOutHashOk() (*string, bool) {
	if o == nil || o.OutHash == nil {
		return nil, false
	}
	return o.OutHash, true
}

// HasOutHash returns a boolean if a field has been set.
func (o *TxOutItem) HasOutHash() bool {
	if o != nil && o.OutHash != nil {
		return true
	}

	return false
}

// SetOutHash gets a reference to the given string and assigns it to the OutHash field.
func (o *TxOutItem) SetOutHash(v string) {
	o.OutHash = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TxOutItem) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutItem) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TxOutItem) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *TxOutItem) SetHeight(v int64) {
	o.Height = &v
}

func (o TxOutItem) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain"] = o.Chain
	}
	if true {
		toSerialize["to_address"] = o.ToAddress
	}
	if o.VaultPubKey != nil {
		toSerialize["vault_pub_key"] = o.VaultPubKey
	}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["max_gas"] = o.MaxGas
	}
	if o.GasRate != nil {
		toSerialize["gas_rate"] = o.GasRate
	}
	if o.InHash != nil {
		toSerialize["in_hash"] = o.InHash
	}
	if o.OutHash != nil {
		toSerialize["out_hash"] = o.OutHash
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	return json.Marshal(toSerialize)
}

type NullableTxOutItem struct {
	value *TxOutItem
	isSet bool
}

func (v NullableTxOutItem) Get() *TxOutItem {
	return v.value
}

func (v *NullableTxOutItem) Set(val *TxOutItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTxOutItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTxOutItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxOutItem(val *TxOutItem) *NullableTxOutItem {
	return &NullableTxOutItem{value: val, isSet: true}
}

func (v NullableTxOutItem) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxOutItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

