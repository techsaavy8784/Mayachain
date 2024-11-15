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

// InboundConfirmationCountedStage struct for InboundConfirmationCountedStage
type InboundConfirmationCountedStage struct {
	// the MAYAChain block height when confirmation counting began
	CountingStartHeight *int64 `json:"counting_start_height,omitempty"`
	// the external source chain for which confirmation counting takes place
	Chain *string `json:"chain,omitempty"`
	// the block height on the external source chain when the transaction was observed
	ExternalObservedHeight *int64 `json:"external_observed_height,omitempty"`
	// the block height on the external source chain when confirmation counting will be complete
	ExternalConfirmationDelayHeight *int64 `json:"external_confirmation_delay_height,omitempty"`
	// the estimated remaining seconds before confirmation counting completes
	RemainingConfirmationSeconds *int64 `json:"remaining_confirmation_seconds,omitempty"`
	// returns true if no transaction confirmation counting remains to be done
	Completed bool `json:"completed"`
}

// NewInboundConfirmationCountedStage instantiates a new InboundConfirmationCountedStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundConfirmationCountedStage(completed bool) *InboundConfirmationCountedStage {
	this := InboundConfirmationCountedStage{}
	this.Completed = completed
	return &this
}

// NewInboundConfirmationCountedStageWithDefaults instantiates a new InboundConfirmationCountedStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundConfirmationCountedStageWithDefaults() *InboundConfirmationCountedStage {
	this := InboundConfirmationCountedStage{}
	return &this
}

// GetCountingStartHeight returns the CountingStartHeight field value if set, zero value otherwise.
func (o *InboundConfirmationCountedStage) GetCountingStartHeight() int64 {
	if o == nil || o.CountingStartHeight == nil {
		var ret int64
		return ret
	}
	return *o.CountingStartHeight
}

// GetCountingStartHeightOk returns a tuple with the CountingStartHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundConfirmationCountedStage) GetCountingStartHeightOk() (*int64, bool) {
	if o == nil || o.CountingStartHeight == nil {
		return nil, false
	}
	return o.CountingStartHeight, true
}

// HasCountingStartHeight returns a boolean if a field has been set.
func (o *InboundConfirmationCountedStage) HasCountingStartHeight() bool {
	if o != nil && o.CountingStartHeight != nil {
		return true
	}

	return false
}

// SetCountingStartHeight gets a reference to the given int64 and assigns it to the CountingStartHeight field.
func (o *InboundConfirmationCountedStage) SetCountingStartHeight(v int64) {
	o.CountingStartHeight = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *InboundConfirmationCountedStage) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundConfirmationCountedStage) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *InboundConfirmationCountedStage) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *InboundConfirmationCountedStage) SetChain(v string) {
	o.Chain = &v
}

// GetExternalObservedHeight returns the ExternalObservedHeight field value if set, zero value otherwise.
func (o *InboundConfirmationCountedStage) GetExternalObservedHeight() int64 {
	if o == nil || o.ExternalObservedHeight == nil {
		var ret int64
		return ret
	}
	return *o.ExternalObservedHeight
}

// GetExternalObservedHeightOk returns a tuple with the ExternalObservedHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundConfirmationCountedStage) GetExternalObservedHeightOk() (*int64, bool) {
	if o == nil || o.ExternalObservedHeight == nil {
		return nil, false
	}
	return o.ExternalObservedHeight, true
}

// HasExternalObservedHeight returns a boolean if a field has been set.
func (o *InboundConfirmationCountedStage) HasExternalObservedHeight() bool {
	if o != nil && o.ExternalObservedHeight != nil {
		return true
	}

	return false
}

// SetExternalObservedHeight gets a reference to the given int64 and assigns it to the ExternalObservedHeight field.
func (o *InboundConfirmationCountedStage) SetExternalObservedHeight(v int64) {
	o.ExternalObservedHeight = &v
}

// GetExternalConfirmationDelayHeight returns the ExternalConfirmationDelayHeight field value if set, zero value otherwise.
func (o *InboundConfirmationCountedStage) GetExternalConfirmationDelayHeight() int64 {
	if o == nil || o.ExternalConfirmationDelayHeight == nil {
		var ret int64
		return ret
	}
	return *o.ExternalConfirmationDelayHeight
}

// GetExternalConfirmationDelayHeightOk returns a tuple with the ExternalConfirmationDelayHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundConfirmationCountedStage) GetExternalConfirmationDelayHeightOk() (*int64, bool) {
	if o == nil || o.ExternalConfirmationDelayHeight == nil {
		return nil, false
	}
	return o.ExternalConfirmationDelayHeight, true
}

// HasExternalConfirmationDelayHeight returns a boolean if a field has been set.
func (o *InboundConfirmationCountedStage) HasExternalConfirmationDelayHeight() bool {
	if o != nil && o.ExternalConfirmationDelayHeight != nil {
		return true
	}

	return false
}

// SetExternalConfirmationDelayHeight gets a reference to the given int64 and assigns it to the ExternalConfirmationDelayHeight field.
func (o *InboundConfirmationCountedStage) SetExternalConfirmationDelayHeight(v int64) {
	o.ExternalConfirmationDelayHeight = &v
}

// GetRemainingConfirmationSeconds returns the RemainingConfirmationSeconds field value if set, zero value otherwise.
func (o *InboundConfirmationCountedStage) GetRemainingConfirmationSeconds() int64 {
	if o == nil || o.RemainingConfirmationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.RemainingConfirmationSeconds
}

// GetRemainingConfirmationSecondsOk returns a tuple with the RemainingConfirmationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundConfirmationCountedStage) GetRemainingConfirmationSecondsOk() (*int64, bool) {
	if o == nil || o.RemainingConfirmationSeconds == nil {
		return nil, false
	}
	return o.RemainingConfirmationSeconds, true
}

// HasRemainingConfirmationSeconds returns a boolean if a field has been set.
func (o *InboundConfirmationCountedStage) HasRemainingConfirmationSeconds() bool {
	if o != nil && o.RemainingConfirmationSeconds != nil {
		return true
	}

	return false
}

// SetRemainingConfirmationSeconds gets a reference to the given int64 and assigns it to the RemainingConfirmationSeconds field.
func (o *InboundConfirmationCountedStage) SetRemainingConfirmationSeconds(v int64) {
	o.RemainingConfirmationSeconds = &v
}

// GetCompleted returns the Completed field value
func (o *InboundConfirmationCountedStage) GetCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value
// and a boolean to check if the value has been set.
func (o *InboundConfirmationCountedStage) GetCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completed, true
}

// SetCompleted sets field value
func (o *InboundConfirmationCountedStage) SetCompleted(v bool) {
	o.Completed = v
}

func (o InboundConfirmationCountedStage) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountingStartHeight != nil {
		toSerialize["counting_start_height"] = o.CountingStartHeight
	}
	if o.Chain != nil {
		toSerialize["chain"] = o.Chain
	}
	if o.ExternalObservedHeight != nil {
		toSerialize["external_observed_height"] = o.ExternalObservedHeight
	}
	if o.ExternalConfirmationDelayHeight != nil {
		toSerialize["external_confirmation_delay_height"] = o.ExternalConfirmationDelayHeight
	}
	if o.RemainingConfirmationSeconds != nil {
		toSerialize["remaining_confirmation_seconds"] = o.RemainingConfirmationSeconds
	}
	if true {
		toSerialize["completed"] = o.Completed
	}
	return json.Marshal(toSerialize)
}

type NullableInboundConfirmationCountedStage struct {
	value *InboundConfirmationCountedStage
	isSet bool
}

func (v NullableInboundConfirmationCountedStage) Get() *InboundConfirmationCountedStage {
	return v.value
}

func (v *NullableInboundConfirmationCountedStage) Set(val *InboundConfirmationCountedStage) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundConfirmationCountedStage) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundConfirmationCountedStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundConfirmationCountedStage(val *InboundConfirmationCountedStage) *NullableInboundConfirmationCountedStage {
	return &NullableInboundConfirmationCountedStage{value: val, isSet: true}
}

func (v NullableInboundConfirmationCountedStage) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundConfirmationCountedStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


