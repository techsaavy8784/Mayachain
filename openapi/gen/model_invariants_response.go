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

// InvariantsResponse struct for InvariantsResponse
type InvariantsResponse struct {
	Invariants []string `json:"invariants,omitempty"`
}

// NewInvariantsResponse instantiates a new InvariantsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvariantsResponse() *InvariantsResponse {
	this := InvariantsResponse{}
	return &this
}

// NewInvariantsResponseWithDefaults instantiates a new InvariantsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvariantsResponseWithDefaults() *InvariantsResponse {
	this := InvariantsResponse{}
	return &this
}

// GetInvariants returns the Invariants field value if set, zero value otherwise.
func (o *InvariantsResponse) GetInvariants() []string {
	if o == nil || o.Invariants == nil {
		var ret []string
		return ret
	}
	return o.Invariants
}

// GetInvariantsOk returns a tuple with the Invariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvariantsResponse) GetInvariantsOk() ([]string, bool) {
	if o == nil || o.Invariants == nil {
		return nil, false
	}
	return o.Invariants, true
}

// HasInvariants returns a boolean if a field has been set.
func (o *InvariantsResponse) HasInvariants() bool {
	if o != nil && o.Invariants != nil {
		return true
	}

	return false
}

// SetInvariants gets a reference to the given []string and assigns it to the Invariants field.
func (o *InvariantsResponse) SetInvariants(v []string) {
	o.Invariants = v
}

func (o InvariantsResponse) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invariants != nil {
		toSerialize["invariants"] = o.Invariants
	}
	return json.Marshal(toSerialize)
}

type NullableInvariantsResponse struct {
	value *InvariantsResponse
	isSet bool
}

func (v NullableInvariantsResponse) Get() *InvariantsResponse {
	return v.value
}

func (v *NullableInvariantsResponse) Set(val *InvariantsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvariantsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvariantsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvariantsResponse(val *InvariantsResponse) *NullableInvariantsResponse {
	return &NullableInvariantsResponse{value: val, isSet: true}
}

func (v NullableInvariantsResponse) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvariantsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


