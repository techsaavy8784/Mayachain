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

// QueueResponse struct for QueueResponse
type QueueResponse struct {
	Swap int64 `json:"swap"`
	// number of signed outbound tx in the queue
	Outbound int64 `json:"outbound"`
	Internal int64 `json:"internal"`
	// scheduled outbound value in CACAO
	ScheduledOutboundValue string `json:"scheduled_outbound_value"`
}

// NewQueueResponse instantiates a new QueueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueResponse(swap int64, outbound int64, internal int64, scheduledOutboundValue string) *QueueResponse {
	this := QueueResponse{}
	this.Swap = swap
	this.Outbound = outbound
	this.Internal = internal
	this.ScheduledOutboundValue = scheduledOutboundValue
	return &this
}

// NewQueueResponseWithDefaults instantiates a new QueueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueResponseWithDefaults() *QueueResponse {
	this := QueueResponse{}
	return &this
}

// GetSwap returns the Swap field value
func (o *QueueResponse) GetSwap() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Swap
}

// GetSwapOk returns a tuple with the Swap field value
// and a boolean to check if the value has been set.
func (o *QueueResponse) GetSwapOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Swap, true
}

// SetSwap sets field value
func (o *QueueResponse) SetSwap(v int64) {
	o.Swap = v
}

// GetOutbound returns the Outbound field value
func (o *QueueResponse) GetOutbound() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value
// and a boolean to check if the value has been set.
func (o *QueueResponse) GetOutboundOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outbound, true
}

// SetOutbound sets field value
func (o *QueueResponse) SetOutbound(v int64) {
	o.Outbound = v
}

// GetInternal returns the Internal field value
func (o *QueueResponse) GetInternal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *QueueResponse) GetInternalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *QueueResponse) SetInternal(v int64) {
	o.Internal = v
}

// GetScheduledOutboundValue returns the ScheduledOutboundValue field value
func (o *QueueResponse) GetScheduledOutboundValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledOutboundValue
}

// GetScheduledOutboundValueOk returns a tuple with the ScheduledOutboundValue field value
// and a boolean to check if the value has been set.
func (o *QueueResponse) GetScheduledOutboundValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledOutboundValue, true
}

// SetScheduledOutboundValue sets field value
func (o *QueueResponse) SetScheduledOutboundValue(v string) {
	o.ScheduledOutboundValue = v
}

func (o QueueResponse) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["swap"] = o.Swap
	}
	if true {
		toSerialize["outbound"] = o.Outbound
	}
	if true {
		toSerialize["internal"] = o.Internal
	}
	if true {
		toSerialize["scheduled_outbound_value"] = o.ScheduledOutboundValue
	}
	return json.Marshal(toSerialize)
}

type NullableQueueResponse struct {
	value *QueueResponse
	isSet bool
}

func (v NullableQueueResponse) Get() *QueueResponse {
	return v.value
}

func (v *NullableQueueResponse) Set(val *QueueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueResponse(val *QueueResponse) *NullableQueueResponse {
	return &NullableQueueResponse{value: val, isSet: true}
}

func (v NullableQueueResponse) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

