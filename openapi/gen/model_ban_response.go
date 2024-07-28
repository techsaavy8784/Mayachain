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

// BanResponse struct for BanResponse
type BanResponse struct {
	NodeAddress *string `json:"node_address,omitempty"`
	BlockHeight *int64 `json:"block_height,omitempty"`
	Signers []string `json:"signers,omitempty"`
}

// NewBanResponse instantiates a new BanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanResponse() *BanResponse {
	this := BanResponse{}
	return &this
}

// NewBanResponseWithDefaults instantiates a new BanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanResponseWithDefaults() *BanResponse {
	this := BanResponse{}
	return &this
}

// GetNodeAddress returns the NodeAddress field value if set, zero value otherwise.
func (o *BanResponse) GetNodeAddress() string {
	if o == nil || o.NodeAddress == nil {
		var ret string
		return ret
	}
	return *o.NodeAddress
}

// GetNodeAddressOk returns a tuple with the NodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanResponse) GetNodeAddressOk() (*string, bool) {
	if o == nil || o.NodeAddress == nil {
		return nil, false
	}
	return o.NodeAddress, true
}

// HasNodeAddress returns a boolean if a field has been set.
func (o *BanResponse) HasNodeAddress() bool {
	if o != nil && o.NodeAddress != nil {
		return true
	}

	return false
}

// SetNodeAddress gets a reference to the given string and assigns it to the NodeAddress field.
func (o *BanResponse) SetNodeAddress(v string) {
	o.NodeAddress = &v
}

// GetBlockHeight returns the BlockHeight field value if set, zero value otherwise.
func (o *BanResponse) GetBlockHeight() int64 {
	if o == nil || o.BlockHeight == nil {
		var ret int64
		return ret
	}
	return *o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanResponse) GetBlockHeightOk() (*int64, bool) {
	if o == nil || o.BlockHeight == nil {
		return nil, false
	}
	return o.BlockHeight, true
}

// HasBlockHeight returns a boolean if a field has been set.
func (o *BanResponse) HasBlockHeight() bool {
	if o != nil && o.BlockHeight != nil {
		return true
	}

	return false
}

// SetBlockHeight gets a reference to the given int64 and assigns it to the BlockHeight field.
func (o *BanResponse) SetBlockHeight(v int64) {
	o.BlockHeight = &v
}

// GetSigners returns the Signers field value if set, zero value otherwise.
func (o *BanResponse) GetSigners() []string {
	if o == nil || o.Signers == nil {
		var ret []string
		return ret
	}
	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BanResponse) GetSignersOk() ([]string, bool) {
	if o == nil || o.Signers == nil {
		return nil, false
	}
	return o.Signers, true
}

// HasSigners returns a boolean if a field has been set.
func (o *BanResponse) HasSigners() bool {
	if o != nil && o.Signers != nil {
		return true
	}

	return false
}

// SetSigners gets a reference to the given []string and assigns it to the Signers field.
func (o *BanResponse) SetSigners(v []string) {
	o.Signers = v
}

func (o BanResponse) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeAddress != nil {
		toSerialize["node_address"] = o.NodeAddress
	}
	if o.BlockHeight != nil {
		toSerialize["block_height"] = o.BlockHeight
	}
	if o.Signers != nil {
		toSerialize["signers"] = o.Signers
	}
	return json.Marshal(toSerialize)
}

type NullableBanResponse struct {
	value *BanResponse
	isSet bool
}

func (v NullableBanResponse) Get() *BanResponse {
	return v.value
}

func (v *NullableBanResponse) Set(val *BanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanResponse(val *BanResponse) *NullableBanResponse {
	return &NullableBanResponse{value: val, isSet: true}
}

func (v NullableBanResponse) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


