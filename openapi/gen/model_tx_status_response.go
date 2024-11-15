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

// TxStatusResponse struct for TxStatusResponse
type TxStatusResponse struct {
	Tx *Tx `json:"tx,omitempty"`
	PlannedOutTxs []PlannedOutTx `json:"planned_out_txs,omitempty"`
	OutTxs []Tx `json:"out_txs,omitempty"`
	Stages TxStagesResponse `json:"stages"`
}

// NewTxStatusResponse instantiates a new TxStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxStatusResponse(stages TxStagesResponse) *TxStatusResponse {
	this := TxStatusResponse{}
	this.Stages = stages
	return &this
}

// NewTxStatusResponseWithDefaults instantiates a new TxStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxStatusResponseWithDefaults() *TxStatusResponse {
	this := TxStatusResponse{}
	return &this
}

// GetTx returns the Tx field value if set, zero value otherwise.
func (o *TxStatusResponse) GetTx() Tx {
	if o == nil || o.Tx == nil {
		var ret Tx
		return ret
	}
	return *o.Tx
}

// GetTxOk returns a tuple with the Tx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxStatusResponse) GetTxOk() (*Tx, bool) {
	if o == nil || o.Tx == nil {
		return nil, false
	}
	return o.Tx, true
}

// HasTx returns a boolean if a field has been set.
func (o *TxStatusResponse) HasTx() bool {
	if o != nil && o.Tx != nil {
		return true
	}

	return false
}

// SetTx gets a reference to the given Tx and assigns it to the Tx field.
func (o *TxStatusResponse) SetTx(v Tx) {
	o.Tx = &v
}

// GetPlannedOutTxs returns the PlannedOutTxs field value if set, zero value otherwise.
func (o *TxStatusResponse) GetPlannedOutTxs() []PlannedOutTx {
	if o == nil || o.PlannedOutTxs == nil {
		var ret []PlannedOutTx
		return ret
	}
	return o.PlannedOutTxs
}

// GetPlannedOutTxsOk returns a tuple with the PlannedOutTxs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxStatusResponse) GetPlannedOutTxsOk() ([]PlannedOutTx, bool) {
	if o == nil || o.PlannedOutTxs == nil {
		return nil, false
	}
	return o.PlannedOutTxs, true
}

// HasPlannedOutTxs returns a boolean if a field has been set.
func (o *TxStatusResponse) HasPlannedOutTxs() bool {
	if o != nil && o.PlannedOutTxs != nil {
		return true
	}

	return false
}

// SetPlannedOutTxs gets a reference to the given []PlannedOutTx and assigns it to the PlannedOutTxs field.
func (o *TxStatusResponse) SetPlannedOutTxs(v []PlannedOutTx) {
	o.PlannedOutTxs = v
}

// GetOutTxs returns the OutTxs field value if set, zero value otherwise.
func (o *TxStatusResponse) GetOutTxs() []Tx {
	if o == nil || o.OutTxs == nil {
		var ret []Tx
		return ret
	}
	return o.OutTxs
}

// GetOutTxsOk returns a tuple with the OutTxs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxStatusResponse) GetOutTxsOk() ([]Tx, bool) {
	if o == nil || o.OutTxs == nil {
		return nil, false
	}
	return o.OutTxs, true
}

// HasOutTxs returns a boolean if a field has been set.
func (o *TxStatusResponse) HasOutTxs() bool {
	if o != nil && o.OutTxs != nil {
		return true
	}

	return false
}

// SetOutTxs gets a reference to the given []Tx and assigns it to the OutTxs field.
func (o *TxStatusResponse) SetOutTxs(v []Tx) {
	o.OutTxs = v
}

// GetStages returns the Stages field value
func (o *TxStatusResponse) GetStages() TxStagesResponse {
	if o == nil {
		var ret TxStagesResponse
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *TxStatusResponse) GetStagesOk() (*TxStagesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stages, true
}

// SetStages sets field value
func (o *TxStatusResponse) SetStages(v TxStagesResponse) {
	o.Stages = v
}

func (o TxStatusResponse) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tx != nil {
		toSerialize["tx"] = o.Tx
	}
	if o.PlannedOutTxs != nil {
		toSerialize["planned_out_txs"] = o.PlannedOutTxs
	}
	if o.OutTxs != nil {
		toSerialize["out_txs"] = o.OutTxs
	}
	if true {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableTxStatusResponse struct {
	value *TxStatusResponse
	isSet bool
}

func (v NullableTxStatusResponse) Get() *TxStatusResponse {
	return v.value
}

func (v *NullableTxStatusResponse) Set(val *TxStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTxStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTxStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxStatusResponse(val *TxStatusResponse) *NullableTxStatusResponse {
	return &NullableTxStatusResponse{value: val, isSet: true}
}

func (v NullableTxStatusResponse) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


