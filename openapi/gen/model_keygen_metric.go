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

// KeygenMetric struct for KeygenMetric
type KeygenMetric struct {
	PubKey *string `json:"pub_key,omitempty"`
	NodeTssTimes []NodeKeygenMetric `json:"node_tss_times"`
}

// NewKeygenMetric instantiates a new KeygenMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeygenMetric(nodeTssTimes []NodeKeygenMetric) *KeygenMetric {
	this := KeygenMetric{}
	this.NodeTssTimes = nodeTssTimes
	return &this
}

// NewKeygenMetricWithDefaults instantiates a new KeygenMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeygenMetricWithDefaults() *KeygenMetric {
	this := KeygenMetric{}
	return &this
}

// GetPubKey returns the PubKey field value if set, zero value otherwise.
func (o *KeygenMetric) GetPubKey() string {
	if o == nil || o.PubKey == nil {
		var ret string
		return ret
	}
	return *o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeygenMetric) GetPubKeyOk() (*string, bool) {
	if o == nil || o.PubKey == nil {
		return nil, false
	}
	return o.PubKey, true
}

// HasPubKey returns a boolean if a field has been set.
func (o *KeygenMetric) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

// SetPubKey gets a reference to the given string and assigns it to the PubKey field.
func (o *KeygenMetric) SetPubKey(v string) {
	o.PubKey = &v
}

// GetNodeTssTimes returns the NodeTssTimes field value
func (o *KeygenMetric) GetNodeTssTimes() []NodeKeygenMetric {
	if o == nil {
		var ret []NodeKeygenMetric
		return ret
	}

	return o.NodeTssTimes
}

// GetNodeTssTimesOk returns a tuple with the NodeTssTimes field value
// and a boolean to check if the value has been set.
func (o *KeygenMetric) GetNodeTssTimesOk() ([]NodeKeygenMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeTssTimes, true
}

// SetNodeTssTimes sets field value
func (o *KeygenMetric) SetNodeTssTimes(v []NodeKeygenMetric) {
	o.NodeTssTimes = v
}

func (o KeygenMetric) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PubKey != nil {
		toSerialize["pub_key"] = o.PubKey
	}
	if true {
		toSerialize["node_tss_times"] = o.NodeTssTimes
	}
	return json.Marshal(toSerialize)
}

type NullableKeygenMetric struct {
	value *KeygenMetric
	isSet bool
}

func (v NullableKeygenMetric) Get() *KeygenMetric {
	return v.value
}

func (v *NullableKeygenMetric) Set(val *KeygenMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableKeygenMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableKeygenMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeygenMetric(val *KeygenMetric) *NullableKeygenMetric {
	return &NullableKeygenMetric{value: val, isSet: true}
}

func (v NullableKeygenMetric) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeygenMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


