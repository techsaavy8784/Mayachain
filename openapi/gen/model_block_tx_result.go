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

// BlockTxResult struct for BlockTxResult
type BlockTxResult struct {
	Code *int64 `json:"code,omitempty"`
	Data *string `json:"data,omitempty"`
	Log *string `json:"log,omitempty"`
	Info *string `json:"info,omitempty"`
	GasWanted *string `json:"gas_wanted,omitempty"`
	GasUsed *string `json:"gas_used,omitempty"`
	Events []map[string]string `json:"events,omitempty"`
	Codespace *string `json:"codespace,omitempty"`
}

// NewBlockTxResult instantiates a new BlockTxResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTxResult() *BlockTxResult {
	this := BlockTxResult{}
	return &this
}

// NewBlockTxResultWithDefaults instantiates a new BlockTxResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTxResultWithDefaults() *BlockTxResult {
	this := BlockTxResult{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BlockTxResult) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BlockTxResult) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *BlockTxResult) SetCode(v int64) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BlockTxResult) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BlockTxResult) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BlockTxResult) SetData(v string) {
	o.Data = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *BlockTxResult) GetLog() string {
	if o == nil || o.Log == nil {
		var ret string
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetLogOk() (*string, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *BlockTxResult) HasLog() bool {
	if o != nil && o.Log != nil {
		return true
	}

	return false
}

// SetLog gets a reference to the given string and assigns it to the Log field.
func (o *BlockTxResult) SetLog(v string) {
	o.Log = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BlockTxResult) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BlockTxResult) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *BlockTxResult) SetInfo(v string) {
	o.Info = &v
}

// GetGasWanted returns the GasWanted field value if set, zero value otherwise.
func (o *BlockTxResult) GetGasWanted() string {
	if o == nil || o.GasWanted == nil {
		var ret string
		return ret
	}
	return *o.GasWanted
}

// GetGasWantedOk returns a tuple with the GasWanted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetGasWantedOk() (*string, bool) {
	if o == nil || o.GasWanted == nil {
		return nil, false
	}
	return o.GasWanted, true
}

// HasGasWanted returns a boolean if a field has been set.
func (o *BlockTxResult) HasGasWanted() bool {
	if o != nil && o.GasWanted != nil {
		return true
	}

	return false
}

// SetGasWanted gets a reference to the given string and assigns it to the GasWanted field.
func (o *BlockTxResult) SetGasWanted(v string) {
	o.GasWanted = &v
}

// GetGasUsed returns the GasUsed field value if set, zero value otherwise.
func (o *BlockTxResult) GetGasUsed() string {
	if o == nil || o.GasUsed == nil {
		var ret string
		return ret
	}
	return *o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetGasUsedOk() (*string, bool) {
	if o == nil || o.GasUsed == nil {
		return nil, false
	}
	return o.GasUsed, true
}

// HasGasUsed returns a boolean if a field has been set.
func (o *BlockTxResult) HasGasUsed() bool {
	if o != nil && o.GasUsed != nil {
		return true
	}

	return false
}

// SetGasUsed gets a reference to the given string and assigns it to the GasUsed field.
func (o *BlockTxResult) SetGasUsed(v string) {
	o.GasUsed = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockTxResult) GetEvents() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockTxResult) GetEventsOk() ([]map[string]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *BlockTxResult) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []map[string]string and assigns it to the Events field.
func (o *BlockTxResult) SetEvents(v []map[string]string) {
	o.Events = v
}

// GetCodespace returns the Codespace field value if set, zero value otherwise.
func (o *BlockTxResult) GetCodespace() string {
	if o == nil || o.Codespace == nil {
		var ret string
		return ret
	}
	return *o.Codespace
}

// GetCodespaceOk returns a tuple with the Codespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTxResult) GetCodespaceOk() (*string, bool) {
	if o == nil || o.Codespace == nil {
		return nil, false
	}
	return o.Codespace, true
}

// HasCodespace returns a boolean if a field has been set.
func (o *BlockTxResult) HasCodespace() bool {
	if o != nil && o.Codespace != nil {
		return true
	}

	return false
}

// SetCodespace gets a reference to the given string and assigns it to the Codespace field.
func (o *BlockTxResult) SetCodespace(v string) {
	o.Codespace = &v
}

func (o BlockTxResult) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.GasWanted != nil {
		toSerialize["gas_wanted"] = o.GasWanted
	}
	if o.GasUsed != nil {
		toSerialize["gas_used"] = o.GasUsed
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Codespace != nil {
		toSerialize["codespace"] = o.Codespace
	}
	return json.Marshal(toSerialize)
}

type NullableBlockTxResult struct {
	value *BlockTxResult
	isSet bool
}

func (v NullableBlockTxResult) Get() *BlockTxResult {
	return v.value
}

func (v *NullableBlockTxResult) Set(val *BlockTxResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTxResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTxResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTxResult(val *BlockTxResult) *NullableBlockTxResult {
	return &NullableBlockTxResult{value: val, isSet: true}
}

func (v NullableBlockTxResult) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTxResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


