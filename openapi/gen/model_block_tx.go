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

// BlockTx struct for BlockTx
type BlockTx struct {
	Hash string `json:"hash"`
	Tx map[string]interface{} `json:"tx"`
	Result BlockTxResult `json:"result"`
}

// NewBlockTx instantiates a new BlockTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTx(hash string, tx map[string]interface{}, result BlockTxResult) *BlockTx {
	this := BlockTx{}
	this.Hash = hash
	this.Tx = tx
	this.Result = result
	return &this
}

// NewBlockTxWithDefaults instantiates a new BlockTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTxWithDefaults() *BlockTx {
	this := BlockTx{}
	return &this
}

// GetHash returns the Hash field value
func (o *BlockTx) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlockTx) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlockTx) SetHash(v string) {
	o.Hash = v
}

// GetTx returns the Tx field value
func (o *BlockTx) GetTx() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *BlockTx) GetTxOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tx, true
}

// SetTx sets field value
func (o *BlockTx) SetTx(v map[string]interface{}) {
	o.Tx = v
}

// GetResult returns the Result field value
func (o *BlockTx) GetResult() BlockTxResult {
	if o == nil {
		var ret BlockTxResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *BlockTx) GetResultOk() (*BlockTxResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *BlockTx) SetResult(v BlockTxResult) {
	o.Result = v
}

func (o BlockTx) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["tx"] = o.Tx
	}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableBlockTx struct {
	value *BlockTx
	isSet bool
}

func (v NullableBlockTx) Get() *BlockTx {
	return v.value
}

func (v *NullableBlockTx) Set(val *BlockTx) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTx) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTx(val *BlockTx) *NullableBlockTx {
	return &NullableBlockTx{value: val, isSet: true}
}

func (v NullableBlockTx) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


