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

// Saver struct for Saver
type Saver struct {
	Asset string `json:"asset"`
	AssetAddress string `json:"asset_address"`
	LastAddHeight *int64 `json:"last_add_height,omitempty"`
	LastWithdrawHeight *int64 `json:"last_withdraw_height,omitempty"`
	Units string `json:"units"`
	AssetDepositValue string `json:"asset_deposit_value"`
	AssetRedeemValue string `json:"asset_redeem_value"`
	GrowthPct string `json:"growth_pct"`
}

// NewSaver instantiates a new Saver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaver(asset string, assetAddress string, units string, assetDepositValue string, assetRedeemValue string, growthPct string) *Saver {
	this := Saver{}
	this.Asset = asset
	this.AssetAddress = assetAddress
	this.Units = units
	this.AssetDepositValue = assetDepositValue
	this.AssetRedeemValue = assetRedeemValue
	this.GrowthPct = growthPct
	return &this
}

// NewSaverWithDefaults instantiates a new Saver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaverWithDefaults() *Saver {
	this := Saver{}
	return &this
}

// GetAsset returns the Asset field value
func (o *Saver) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *Saver) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *Saver) SetAsset(v string) {
	o.Asset = v
}

// GetAssetAddress returns the AssetAddress field value
func (o *Saver) GetAssetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetAddress
}

// GetAssetAddressOk returns a tuple with the AssetAddress field value
// and a boolean to check if the value has been set.
func (o *Saver) GetAssetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetAddress, true
}

// SetAssetAddress sets field value
func (o *Saver) SetAssetAddress(v string) {
	o.AssetAddress = v
}

// GetLastAddHeight returns the LastAddHeight field value if set, zero value otherwise.
func (o *Saver) GetLastAddHeight() int64 {
	if o == nil || o.LastAddHeight == nil {
		var ret int64
		return ret
	}
	return *o.LastAddHeight
}

// GetLastAddHeightOk returns a tuple with the LastAddHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saver) GetLastAddHeightOk() (*int64, bool) {
	if o == nil || o.LastAddHeight == nil {
		return nil, false
	}
	return o.LastAddHeight, true
}

// HasLastAddHeight returns a boolean if a field has been set.
func (o *Saver) HasLastAddHeight() bool {
	if o != nil && o.LastAddHeight != nil {
		return true
	}

	return false
}

// SetLastAddHeight gets a reference to the given int64 and assigns it to the LastAddHeight field.
func (o *Saver) SetLastAddHeight(v int64) {
	o.LastAddHeight = &v
}

// GetLastWithdrawHeight returns the LastWithdrawHeight field value if set, zero value otherwise.
func (o *Saver) GetLastWithdrawHeight() int64 {
	if o == nil || o.LastWithdrawHeight == nil {
		var ret int64
		return ret
	}
	return *o.LastWithdrawHeight
}

// GetLastWithdrawHeightOk returns a tuple with the LastWithdrawHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saver) GetLastWithdrawHeightOk() (*int64, bool) {
	if o == nil || o.LastWithdrawHeight == nil {
		return nil, false
	}
	return o.LastWithdrawHeight, true
}

// HasLastWithdrawHeight returns a boolean if a field has been set.
func (o *Saver) HasLastWithdrawHeight() bool {
	if o != nil && o.LastWithdrawHeight != nil {
		return true
	}

	return false
}

// SetLastWithdrawHeight gets a reference to the given int64 and assigns it to the LastWithdrawHeight field.
func (o *Saver) SetLastWithdrawHeight(v int64) {
	o.LastWithdrawHeight = &v
}

// GetUnits returns the Units field value
func (o *Saver) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *Saver) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *Saver) SetUnits(v string) {
	o.Units = v
}

// GetAssetDepositValue returns the AssetDepositValue field value
func (o *Saver) GetAssetDepositValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetDepositValue
}

// GetAssetDepositValueOk returns a tuple with the AssetDepositValue field value
// and a boolean to check if the value has been set.
func (o *Saver) GetAssetDepositValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetDepositValue, true
}

// SetAssetDepositValue sets field value
func (o *Saver) SetAssetDepositValue(v string) {
	o.AssetDepositValue = v
}

// GetAssetRedeemValue returns the AssetRedeemValue field value
func (o *Saver) GetAssetRedeemValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetRedeemValue
}

// GetAssetRedeemValueOk returns a tuple with the AssetRedeemValue field value
// and a boolean to check if the value has been set.
func (o *Saver) GetAssetRedeemValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetRedeemValue, true
}

// SetAssetRedeemValue sets field value
func (o *Saver) SetAssetRedeemValue(v string) {
	o.AssetRedeemValue = v
}

// GetGrowthPct returns the GrowthPct field value
func (o *Saver) GetGrowthPct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrowthPct
}

// GetGrowthPctOk returns a tuple with the GrowthPct field value
// and a boolean to check if the value has been set.
func (o *Saver) GetGrowthPctOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrowthPct, true
}

// SetGrowthPct sets field value
func (o *Saver) SetGrowthPct(v string) {
	o.GrowthPct = v
}

func (o Saver) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["asset_address"] = o.AssetAddress
	}
	if o.LastAddHeight != nil {
		toSerialize["last_add_height"] = o.LastAddHeight
	}
	if o.LastWithdrawHeight != nil {
		toSerialize["last_withdraw_height"] = o.LastWithdrawHeight
	}
	if true {
		toSerialize["units"] = o.Units
	}
	if true {
		toSerialize["asset_deposit_value"] = o.AssetDepositValue
	}
	if true {
		toSerialize["asset_redeem_value"] = o.AssetRedeemValue
	}
	if true {
		toSerialize["growth_pct"] = o.GrowthPct
	}
	return json.Marshal(toSerialize)
}

type NullableSaver struct {
	value *Saver
	isSet bool
}

func (v NullableSaver) Get() *Saver {
	return v.value
}

func (v *NullableSaver) Set(val *Saver) {
	v.value = val
	v.isSet = true
}

func (v NullableSaver) IsSet() bool {
	return v.isSet
}

func (v *NullableSaver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaver(val *Saver) *NullableSaver {
	return &NullableSaver{value: val, isSet: true}
}

func (v NullableSaver) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


