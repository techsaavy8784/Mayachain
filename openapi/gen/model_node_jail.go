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

// NodeJail struct for NodeJail
type NodeJail struct {
	NodeAddress *string `json:"node_address,omitempty"`
	ReleaseHeight *int64 `json:"release_height,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewNodeJail instantiates a new NodeJail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeJail() *NodeJail {
	this := NodeJail{}
	return &this
}

// NewNodeJailWithDefaults instantiates a new NodeJail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeJailWithDefaults() *NodeJail {
	this := NodeJail{}
	return &this
}

// GetNodeAddress returns the NodeAddress field value if set, zero value otherwise.
func (o *NodeJail) GetNodeAddress() string {
	if o == nil || o.NodeAddress == nil {
		var ret string
		return ret
	}
	return *o.NodeAddress
}

// GetNodeAddressOk returns a tuple with the NodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeJail) GetNodeAddressOk() (*string, bool) {
	if o == nil || o.NodeAddress == nil {
		return nil, false
	}
	return o.NodeAddress, true
}

// HasNodeAddress returns a boolean if a field has been set.
func (o *NodeJail) HasNodeAddress() bool {
	if o != nil && o.NodeAddress != nil {
		return true
	}

	return false
}

// SetNodeAddress gets a reference to the given string and assigns it to the NodeAddress field.
func (o *NodeJail) SetNodeAddress(v string) {
	o.NodeAddress = &v
}

// GetReleaseHeight returns the ReleaseHeight field value if set, zero value otherwise.
func (o *NodeJail) GetReleaseHeight() int64 {
	if o == nil || o.ReleaseHeight == nil {
		var ret int64
		return ret
	}
	return *o.ReleaseHeight
}

// GetReleaseHeightOk returns a tuple with the ReleaseHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeJail) GetReleaseHeightOk() (*int64, bool) {
	if o == nil || o.ReleaseHeight == nil {
		return nil, false
	}
	return o.ReleaseHeight, true
}

// HasReleaseHeight returns a boolean if a field has been set.
func (o *NodeJail) HasReleaseHeight() bool {
	if o != nil && o.ReleaseHeight != nil {
		return true
	}

	return false
}

// SetReleaseHeight gets a reference to the given int64 and assigns it to the ReleaseHeight field.
func (o *NodeJail) SetReleaseHeight(v int64) {
	o.ReleaseHeight = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *NodeJail) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeJail) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *NodeJail) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *NodeJail) SetReason(v string) {
	o.Reason = &v
}

func (o NodeJail) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeAddress != nil {
		toSerialize["node_address"] = o.NodeAddress
	}
	if o.ReleaseHeight != nil {
		toSerialize["release_height"] = o.ReleaseHeight
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableNodeJail struct {
	value *NodeJail
	isSet bool
}

func (v NullableNodeJail) Get() *NodeJail {
	return v.value
}

func (v *NullableNodeJail) Set(val *NodeJail) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeJail) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeJail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeJail(val *NodeJail) *NullableNodeJail {
	return &NullableNodeJail{value: val, isSet: true}
}

func (v NullableNodeJail) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeJail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


