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

// TssKeysignMetric struct for TssKeysignMetric
type TssKeysignMetric struct {
	TxId *string `json:"tx_id,omitempty"`
	NodeTssTimes []TssMetric `json:"node_tss_times"`
}

// NewTssKeysignMetric instantiates a new TssKeysignMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTssKeysignMetric(nodeTssTimes []TssMetric) *TssKeysignMetric {
	this := TssKeysignMetric{}
	this.NodeTssTimes = nodeTssTimes
	return &this
}

// NewTssKeysignMetricWithDefaults instantiates a new TssKeysignMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTssKeysignMetricWithDefaults() *TssKeysignMetric {
	this := TssKeysignMetric{}
	return &this
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *TssKeysignMetric) GetTxId() string {
	if o == nil || o.TxId == nil {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TssKeysignMetric) GetTxIdOk() (*string, bool) {
	if o == nil || o.TxId == nil {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *TssKeysignMetric) HasTxId() bool {
	if o != nil && o.TxId != nil {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *TssKeysignMetric) SetTxId(v string) {
	o.TxId = &v
}

// GetNodeTssTimes returns the NodeTssTimes field value
func (o *TssKeysignMetric) GetNodeTssTimes() []TssMetric {
	if o == nil {
		var ret []TssMetric
		return ret
	}

	return o.NodeTssTimes
}

// GetNodeTssTimesOk returns a tuple with the NodeTssTimes field value
// and a boolean to check if the value has been set.
func (o *TssKeysignMetric) GetNodeTssTimesOk() ([]TssMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeTssTimes, true
}

// SetNodeTssTimes sets field value
func (o *TssKeysignMetric) SetNodeTssTimes(v []TssMetric) {
	o.NodeTssTimes = v
}

func (o TssKeysignMetric) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TxId != nil {
		toSerialize["tx_id"] = o.TxId
	}
	if true {
		toSerialize["node_tss_times"] = o.NodeTssTimes
	}
	return json.Marshal(toSerialize)
}

type NullableTssKeysignMetric struct {
	value *TssKeysignMetric
	isSet bool
}

func (v NullableTssKeysignMetric) Get() *TssKeysignMetric {
	return v.value
}

func (v *NullableTssKeysignMetric) Set(val *TssKeysignMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableTssKeysignMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableTssKeysignMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTssKeysignMetric(val *TssKeysignMetric) *NullableTssKeysignMetric {
	return &NullableTssKeysignMetric{value: val, isSet: true}
}

func (v NullableTssKeysignMetric) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTssKeysignMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


