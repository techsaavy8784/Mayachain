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

// Keygen struct for Keygen
type Keygen struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Members []string `json:"members,omitempty"`
}

// NewKeygen instantiates a new Keygen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeygen() *Keygen {
	this := Keygen{}
	return &this
}

// NewKeygenWithDefaults instantiates a new Keygen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeygenWithDefaults() *Keygen {
	this := Keygen{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Keygen) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keygen) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Keygen) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Keygen) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Keygen) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keygen) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Keygen) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Keygen) SetType(v string) {
	o.Type = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Keygen) GetMembers() []string {
	if o == nil || o.Members == nil {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keygen) GetMembersOk() ([]string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Keygen) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *Keygen) SetMembers(v []string) {
	o.Members = v
}

func (o Keygen) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableKeygen struct {
	value *Keygen
	isSet bool
}

func (v NullableKeygen) Get() *Keygen {
	return v.value
}

func (v *NullableKeygen) Set(val *Keygen) {
	v.value = val
	v.isSet = true
}

func (v NullableKeygen) IsSet() bool {
	return v.isSet
}

func (v *NullableKeygen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeygen(val *Keygen) *NullableKeygen {
	return &NullableKeygen{value: val, isSet: true}
}

func (v NullableKeygen) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeygen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


