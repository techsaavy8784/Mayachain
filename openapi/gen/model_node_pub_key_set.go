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

// NodePubKeySet struct for NodePubKeySet
type NodePubKeySet struct {
	Secp256k1 *string `json:"secp256k1,omitempty"`
	Ed25519 *string `json:"ed25519,omitempty"`
}

// NewNodePubKeySet instantiates a new NodePubKeySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePubKeySet() *NodePubKeySet {
	this := NodePubKeySet{}
	return &this
}

// NewNodePubKeySetWithDefaults instantiates a new NodePubKeySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePubKeySetWithDefaults() *NodePubKeySet {
	this := NodePubKeySet{}
	return &this
}

// GetSecp256k1 returns the Secp256k1 field value if set, zero value otherwise.
func (o *NodePubKeySet) GetSecp256k1() string {
	if o == nil || o.Secp256k1 == nil {
		var ret string
		return ret
	}
	return *o.Secp256k1
}

// GetSecp256k1Ok returns a tuple with the Secp256k1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePubKeySet) GetSecp256k1Ok() (*string, bool) {
	if o == nil || o.Secp256k1 == nil {
		return nil, false
	}
	return o.Secp256k1, true
}

// HasSecp256k1 returns a boolean if a field has been set.
func (o *NodePubKeySet) HasSecp256k1() bool {
	if o != nil && o.Secp256k1 != nil {
		return true
	}

	return false
}

// SetSecp256k1 gets a reference to the given string and assigns it to the Secp256k1 field.
func (o *NodePubKeySet) SetSecp256k1(v string) {
	o.Secp256k1 = &v
}

// GetEd25519 returns the Ed25519 field value if set, zero value otherwise.
func (o *NodePubKeySet) GetEd25519() string {
	if o == nil || o.Ed25519 == nil {
		var ret string
		return ret
	}
	return *o.Ed25519
}

// GetEd25519Ok returns a tuple with the Ed25519 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePubKeySet) GetEd25519Ok() (*string, bool) {
	if o == nil || o.Ed25519 == nil {
		return nil, false
	}
	return o.Ed25519, true
}

// HasEd25519 returns a boolean if a field has been set.
func (o *NodePubKeySet) HasEd25519() bool {
	if o != nil && o.Ed25519 != nil {
		return true
	}

	return false
}

// SetEd25519 gets a reference to the given string and assigns it to the Ed25519 field.
func (o *NodePubKeySet) SetEd25519(v string) {
	o.Ed25519 = &v
}

func (o NodePubKeySet) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secp256k1 != nil {
		toSerialize["secp256k1"] = o.Secp256k1
	}
	if o.Ed25519 != nil {
		toSerialize["ed25519"] = o.Ed25519
	}
	return json.Marshal(toSerialize)
}

type NullableNodePubKeySet struct {
	value *NodePubKeySet
	isSet bool
}

func (v NullableNodePubKeySet) Get() *NodePubKeySet {
	return v.value
}

func (v *NullableNodePubKeySet) Set(val *NodePubKeySet) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePubKeySet) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePubKeySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePubKeySet(val *NodePubKeySet) *NullableNodePubKeySet {
	return &NullableNodePubKeySet{value: val, isSet: true}
}

func (v NullableNodePubKeySet) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePubKeySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

