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

// BlockResponseHeaderVersion struct for BlockResponseHeaderVersion
type BlockResponseHeaderVersion struct {
	Block string `json:"block"`
	App string `json:"app"`
}

// NewBlockResponseHeaderVersion instantiates a new BlockResponseHeaderVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockResponseHeaderVersion(block string, app string) *BlockResponseHeaderVersion {
	this := BlockResponseHeaderVersion{}
	this.Block = block
	this.App = app
	return &this
}

// NewBlockResponseHeaderVersionWithDefaults instantiates a new BlockResponseHeaderVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockResponseHeaderVersionWithDefaults() *BlockResponseHeaderVersion {
	this := BlockResponseHeaderVersion{}
	return &this
}

// GetBlock returns the Block field value
func (o *BlockResponseHeaderVersion) GetBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeaderVersion) GetBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *BlockResponseHeaderVersion) SetBlock(v string) {
	o.Block = v
}

// GetApp returns the App field value
func (o *BlockResponseHeaderVersion) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeaderVersion) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *BlockResponseHeaderVersion) SetApp(v string) {
	o.App = v
}

func (o BlockResponseHeaderVersion) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["block"] = o.Block
	}
	if true {
		toSerialize["app"] = o.App
	}
	return json.Marshal(toSerialize)
}

type NullableBlockResponseHeaderVersion struct {
	value *BlockResponseHeaderVersion
	isSet bool
}

func (v NullableBlockResponseHeaderVersion) Get() *BlockResponseHeaderVersion {
	return v.value
}

func (v *NullableBlockResponseHeaderVersion) Set(val *BlockResponseHeaderVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockResponseHeaderVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockResponseHeaderVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockResponseHeaderVersion(val *BlockResponseHeaderVersion) *NullableBlockResponseHeaderVersion {
	return &NullableBlockResponseHeaderVersion{value: val, isSet: true}
}

func (v NullableBlockResponseHeaderVersion) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockResponseHeaderVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

