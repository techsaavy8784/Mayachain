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

// NodeBondProviders struct for NodeBondProviders
type NodeBondProviders struct {
	NodeAddress *string `json:"node_address,omitempty"`
	NodeOperatorFee string `json:"node_operator_fee"`
	Providers []NodeBondProvider `json:"providers"`
}

// NewNodeBondProviders instantiates a new NodeBondProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeBondProviders(nodeOperatorFee string, providers []NodeBondProvider) *NodeBondProviders {
	this := NodeBondProviders{}
	this.NodeOperatorFee = nodeOperatorFee
	this.Providers = providers
	return &this
}

// NewNodeBondProvidersWithDefaults instantiates a new NodeBondProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeBondProvidersWithDefaults() *NodeBondProviders {
	this := NodeBondProviders{}
	return &this
}

// GetNodeAddress returns the NodeAddress field value if set, zero value otherwise.
func (o *NodeBondProviders) GetNodeAddress() string {
	if o == nil || o.NodeAddress == nil {
		var ret string
		return ret
	}
	return *o.NodeAddress
}

// GetNodeAddressOk returns a tuple with the NodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBondProviders) GetNodeAddressOk() (*string, bool) {
	if o == nil || o.NodeAddress == nil {
		return nil, false
	}
	return o.NodeAddress, true
}

// HasNodeAddress returns a boolean if a field has been set.
func (o *NodeBondProviders) HasNodeAddress() bool {
	if o != nil && o.NodeAddress != nil {
		return true
	}

	return false
}

// SetNodeAddress gets a reference to the given string and assigns it to the NodeAddress field.
func (o *NodeBondProviders) SetNodeAddress(v string) {
	o.NodeAddress = &v
}

// GetNodeOperatorFee returns the NodeOperatorFee field value
func (o *NodeBondProviders) GetNodeOperatorFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeOperatorFee
}

// GetNodeOperatorFeeOk returns a tuple with the NodeOperatorFee field value
// and a boolean to check if the value has been set.
func (o *NodeBondProviders) GetNodeOperatorFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeOperatorFee, true
}

// SetNodeOperatorFee sets field value
func (o *NodeBondProviders) SetNodeOperatorFee(v string) {
	o.NodeOperatorFee = v
}

// GetProviders returns the Providers field value
func (o *NodeBondProviders) GetProviders() []NodeBondProvider {
	if o == nil {
		var ret []NodeBondProvider
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *NodeBondProviders) GetProvidersOk() ([]NodeBondProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *NodeBondProviders) SetProviders(v []NodeBondProvider) {
	o.Providers = v
}

func (o NodeBondProviders) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeAddress != nil {
		toSerialize["node_address"] = o.NodeAddress
	}
	if true {
		toSerialize["node_operator_fee"] = o.NodeOperatorFee
	}
	if true {
		toSerialize["providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableNodeBondProviders struct {
	value *NodeBondProviders
	isSet bool
}

func (v NullableNodeBondProviders) Get() *NodeBondProviders {
	return v.value
}

func (v *NullableNodeBondProviders) Set(val *NodeBondProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeBondProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeBondProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeBondProviders(val *NodeBondProviders) *NullableNodeBondProviders {
	return &NullableNodeBondProviders{value: val, isSet: true}
}

func (v NullableNodeBondProviders) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeBondProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


