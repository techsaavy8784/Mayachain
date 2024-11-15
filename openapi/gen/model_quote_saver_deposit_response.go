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

// QuoteSaverDepositResponse struct for QuoteSaverDepositResponse
type QuoteSaverDepositResponse struct {
	// the inbound address for the transaction on the source chain
	InboundAddress string `json:"inbound_address"`
	// the approximate number of source chain blocks required before processing
	InboundConfirmationBlocks *int64 `json:"inbound_confirmation_blocks,omitempty"`
	// the approximate seconds for block confirmations required before processing
	InboundConfirmationSeconds *int64 `json:"inbound_confirmation_seconds,omitempty"`
	// the number of mayachain blocks the outbound will be delayed
	OutboundDelayBlocks *int64 `json:"outbound_delay_blocks,omitempty"`
	// the approximate seconds for the outbound delay before it will be sent
	OutboundDelaySeconds *int64 `json:"outbound_delay_seconds,omitempty"`
	Fees QuoteFees `json:"fees"`
	// the EVM chain router contract address
	Router *string `json:"router,omitempty"`
	// expiration timestamp in unix seconds
	Expiry int64 `json:"expiry"`
	// static warning message
	Warning string `json:"warning"`
	// chain specific quote notes
	Notes string `json:"notes"`
	// Defines the minimum transaction size for the chain in base units (sats, wei, uatom). Transactions with asset amounts lower than the dust_threshold are ignored.
	DustThreshold *string `json:"dust_threshold,omitempty"`
	// The recommended minimum inbound amount for this transaction type & inbound asset. Sending less than this amount could result in failed refunds.
	RecommendedMinAmountIn *string `json:"recommended_min_amount_in,omitempty"`
	// the recommended gas rate to use for the inbound to ensure timely confirmation
	RecommendedGasRate string `json:"recommended_gas_rate"`
	// the units of the recommended gas rate
	GasRateUnits string `json:"gas_rate_units"`
	// generated memo for the deposit
	Memo string `json:"memo"`
	// same as expected_amount_deposit, to be deprecated in favour of expected_amount_deposit
	ExpectedAmountOut *string `json:"expected_amount_out,omitempty"`
	// the amount of the target asset the user can expect to deposit after fees
	ExpectedAmountDeposit string `json:"expected_amount_deposit"`
}

// NewQuoteSaverDepositResponse instantiates a new QuoteSaverDepositResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSaverDepositResponse(inboundAddress string, fees QuoteFees, expiry int64, warning string, notes string, recommendedGasRate string, gasRateUnits string, memo string, expectedAmountDeposit string) *QuoteSaverDepositResponse {
	this := QuoteSaverDepositResponse{}
	this.InboundAddress = inboundAddress
	this.Fees = fees
	this.Expiry = expiry
	this.Warning = warning
	this.Notes = notes
	this.RecommendedGasRate = recommendedGasRate
	this.GasRateUnits = gasRateUnits
	this.Memo = memo
	this.ExpectedAmountDeposit = expectedAmountDeposit
	return &this
}

// NewQuoteSaverDepositResponseWithDefaults instantiates a new QuoteSaverDepositResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSaverDepositResponseWithDefaults() *QuoteSaverDepositResponse {
	this := QuoteSaverDepositResponse{}
	return &this
}

// GetInboundAddress returns the InboundAddress field value
func (o *QuoteSaverDepositResponse) GetInboundAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InboundAddress
}

// GetInboundAddressOk returns a tuple with the InboundAddress field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetInboundAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InboundAddress, true
}

// SetInboundAddress sets field value
func (o *QuoteSaverDepositResponse) SetInboundAddress(v string) {
	o.InboundAddress = v
}

// GetInboundConfirmationBlocks returns the InboundConfirmationBlocks field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetInboundConfirmationBlocks() int64 {
	if o == nil || o.InboundConfirmationBlocks == nil {
		var ret int64
		return ret
	}
	return *o.InboundConfirmationBlocks
}

// GetInboundConfirmationBlocksOk returns a tuple with the InboundConfirmationBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetInboundConfirmationBlocksOk() (*int64, bool) {
	if o == nil || o.InboundConfirmationBlocks == nil {
		return nil, false
	}
	return o.InboundConfirmationBlocks, true
}

// HasInboundConfirmationBlocks returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasInboundConfirmationBlocks() bool {
	if o != nil && o.InboundConfirmationBlocks != nil {
		return true
	}

	return false
}

// SetInboundConfirmationBlocks gets a reference to the given int64 and assigns it to the InboundConfirmationBlocks field.
func (o *QuoteSaverDepositResponse) SetInboundConfirmationBlocks(v int64) {
	o.InboundConfirmationBlocks = &v
}

// GetInboundConfirmationSeconds returns the InboundConfirmationSeconds field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetInboundConfirmationSeconds() int64 {
	if o == nil || o.InboundConfirmationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.InboundConfirmationSeconds
}

// GetInboundConfirmationSecondsOk returns a tuple with the InboundConfirmationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetInboundConfirmationSecondsOk() (*int64, bool) {
	if o == nil || o.InboundConfirmationSeconds == nil {
		return nil, false
	}
	return o.InboundConfirmationSeconds, true
}

// HasInboundConfirmationSeconds returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasInboundConfirmationSeconds() bool {
	if o != nil && o.InboundConfirmationSeconds != nil {
		return true
	}

	return false
}

// SetInboundConfirmationSeconds gets a reference to the given int64 and assigns it to the InboundConfirmationSeconds field.
func (o *QuoteSaverDepositResponse) SetInboundConfirmationSeconds(v int64) {
	o.InboundConfirmationSeconds = &v
}

// GetOutboundDelayBlocks returns the OutboundDelayBlocks field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetOutboundDelayBlocks() int64 {
	if o == nil || o.OutboundDelayBlocks == nil {
		var ret int64
		return ret
	}
	return *o.OutboundDelayBlocks
}

// GetOutboundDelayBlocksOk returns a tuple with the OutboundDelayBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetOutboundDelayBlocksOk() (*int64, bool) {
	if o == nil || o.OutboundDelayBlocks == nil {
		return nil, false
	}
	return o.OutboundDelayBlocks, true
}

// HasOutboundDelayBlocks returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasOutboundDelayBlocks() bool {
	if o != nil && o.OutboundDelayBlocks != nil {
		return true
	}

	return false
}

// SetOutboundDelayBlocks gets a reference to the given int64 and assigns it to the OutboundDelayBlocks field.
func (o *QuoteSaverDepositResponse) SetOutboundDelayBlocks(v int64) {
	o.OutboundDelayBlocks = &v
}

// GetOutboundDelaySeconds returns the OutboundDelaySeconds field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetOutboundDelaySeconds() int64 {
	if o == nil || o.OutboundDelaySeconds == nil {
		var ret int64
		return ret
	}
	return *o.OutboundDelaySeconds
}

// GetOutboundDelaySecondsOk returns a tuple with the OutboundDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetOutboundDelaySecondsOk() (*int64, bool) {
	if o == nil || o.OutboundDelaySeconds == nil {
		return nil, false
	}
	return o.OutboundDelaySeconds, true
}

// HasOutboundDelaySeconds returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasOutboundDelaySeconds() bool {
	if o != nil && o.OutboundDelaySeconds != nil {
		return true
	}

	return false
}

// SetOutboundDelaySeconds gets a reference to the given int64 and assigns it to the OutboundDelaySeconds field.
func (o *QuoteSaverDepositResponse) SetOutboundDelaySeconds(v int64) {
	o.OutboundDelaySeconds = &v
}

// GetFees returns the Fees field value
func (o *QuoteSaverDepositResponse) GetFees() QuoteFees {
	if o == nil {
		var ret QuoteFees
		return ret
	}

	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetFeesOk() (*QuoteFees, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fees, true
}

// SetFees sets field value
func (o *QuoteSaverDepositResponse) SetFees(v QuoteFees) {
	o.Fees = v
}

// GetRouter returns the Router field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetRouter() string {
	if o == nil || o.Router == nil {
		var ret string
		return ret
	}
	return *o.Router
}

// GetRouterOk returns a tuple with the Router field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetRouterOk() (*string, bool) {
	if o == nil || o.Router == nil {
		return nil, false
	}
	return o.Router, true
}

// HasRouter returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasRouter() bool {
	if o != nil && o.Router != nil {
		return true
	}

	return false
}

// SetRouter gets a reference to the given string and assigns it to the Router field.
func (o *QuoteSaverDepositResponse) SetRouter(v string) {
	o.Router = &v
}

// GetExpiry returns the Expiry field value
func (o *QuoteSaverDepositResponse) GetExpiry() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetExpiryOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *QuoteSaverDepositResponse) SetExpiry(v int64) {
	o.Expiry = v
}

// GetWarning returns the Warning field value
func (o *QuoteSaverDepositResponse) GetWarning() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warning, true
}

// SetWarning sets field value
func (o *QuoteSaverDepositResponse) SetWarning(v string) {
	o.Warning = v
}

// GetNotes returns the Notes field value
func (o *QuoteSaverDepositResponse) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *QuoteSaverDepositResponse) SetNotes(v string) {
	o.Notes = v
}

// GetDustThreshold returns the DustThreshold field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetDustThreshold() string {
	if o == nil || o.DustThreshold == nil {
		var ret string
		return ret
	}
	return *o.DustThreshold
}

// GetDustThresholdOk returns a tuple with the DustThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetDustThresholdOk() (*string, bool) {
	if o == nil || o.DustThreshold == nil {
		return nil, false
	}
	return o.DustThreshold, true
}

// HasDustThreshold returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasDustThreshold() bool {
	if o != nil && o.DustThreshold != nil {
		return true
	}

	return false
}

// SetDustThreshold gets a reference to the given string and assigns it to the DustThreshold field.
func (o *QuoteSaverDepositResponse) SetDustThreshold(v string) {
	o.DustThreshold = &v
}

// GetRecommendedMinAmountIn returns the RecommendedMinAmountIn field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetRecommendedMinAmountIn() string {
	if o == nil || o.RecommendedMinAmountIn == nil {
		var ret string
		return ret
	}
	return *o.RecommendedMinAmountIn
}

// GetRecommendedMinAmountInOk returns a tuple with the RecommendedMinAmountIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetRecommendedMinAmountInOk() (*string, bool) {
	if o == nil || o.RecommendedMinAmountIn == nil {
		return nil, false
	}
	return o.RecommendedMinAmountIn, true
}

// HasRecommendedMinAmountIn returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasRecommendedMinAmountIn() bool {
	if o != nil && o.RecommendedMinAmountIn != nil {
		return true
	}

	return false
}

// SetRecommendedMinAmountIn gets a reference to the given string and assigns it to the RecommendedMinAmountIn field.
func (o *QuoteSaverDepositResponse) SetRecommendedMinAmountIn(v string) {
	o.RecommendedMinAmountIn = &v
}

// GetRecommendedGasRate returns the RecommendedGasRate field value
func (o *QuoteSaverDepositResponse) GetRecommendedGasRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendedGasRate
}

// GetRecommendedGasRateOk returns a tuple with the RecommendedGasRate field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetRecommendedGasRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedGasRate, true
}

// SetRecommendedGasRate sets field value
func (o *QuoteSaverDepositResponse) SetRecommendedGasRate(v string) {
	o.RecommendedGasRate = v
}

// GetGasRateUnits returns the GasRateUnits field value
func (o *QuoteSaverDepositResponse) GetGasRateUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasRateUnits
}

// GetGasRateUnitsOk returns a tuple with the GasRateUnits field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetGasRateUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasRateUnits, true
}

// SetGasRateUnits sets field value
func (o *QuoteSaverDepositResponse) SetGasRateUnits(v string) {
	o.GasRateUnits = v
}

// GetMemo returns the Memo field value
func (o *QuoteSaverDepositResponse) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *QuoteSaverDepositResponse) SetMemo(v string) {
	o.Memo = v
}

// GetExpectedAmountOut returns the ExpectedAmountOut field value if set, zero value otherwise.
func (o *QuoteSaverDepositResponse) GetExpectedAmountOut() string {
	if o == nil || o.ExpectedAmountOut == nil {
		var ret string
		return ret
	}
	return *o.ExpectedAmountOut
}

// GetExpectedAmountOutOk returns a tuple with the ExpectedAmountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetExpectedAmountOutOk() (*string, bool) {
	if o == nil || o.ExpectedAmountOut == nil {
		return nil, false
	}
	return o.ExpectedAmountOut, true
}

// HasExpectedAmountOut returns a boolean if a field has been set.
func (o *QuoteSaverDepositResponse) HasExpectedAmountOut() bool {
	if o != nil && o.ExpectedAmountOut != nil {
		return true
	}

	return false
}

// SetExpectedAmountOut gets a reference to the given string and assigns it to the ExpectedAmountOut field.
func (o *QuoteSaverDepositResponse) SetExpectedAmountOut(v string) {
	o.ExpectedAmountOut = &v
}

// GetExpectedAmountDeposit returns the ExpectedAmountDeposit field value
func (o *QuoteSaverDepositResponse) GetExpectedAmountDeposit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpectedAmountDeposit
}

// GetExpectedAmountDepositOk returns a tuple with the ExpectedAmountDeposit field value
// and a boolean to check if the value has been set.
func (o *QuoteSaverDepositResponse) GetExpectedAmountDepositOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpectedAmountDeposit, true
}

// SetExpectedAmountDeposit sets field value
func (o *QuoteSaverDepositResponse) SetExpectedAmountDeposit(v string) {
	o.ExpectedAmountDeposit = v
}

func (o QuoteSaverDepositResponse) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inbound_address"] = o.InboundAddress
	}
	if o.InboundConfirmationBlocks != nil {
		toSerialize["inbound_confirmation_blocks"] = o.InboundConfirmationBlocks
	}
	if o.InboundConfirmationSeconds != nil {
		toSerialize["inbound_confirmation_seconds"] = o.InboundConfirmationSeconds
	}
	if o.OutboundDelayBlocks != nil {
		toSerialize["outbound_delay_blocks"] = o.OutboundDelayBlocks
	}
	if o.OutboundDelaySeconds != nil {
		toSerialize["outbound_delay_seconds"] = o.OutboundDelaySeconds
	}
	if true {
		toSerialize["fees"] = o.Fees
	}
	if o.Router != nil {
		toSerialize["router"] = o.Router
	}
	if true {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["warning"] = o.Warning
	}
	if true {
		toSerialize["notes"] = o.Notes
	}
	if o.DustThreshold != nil {
		toSerialize["dust_threshold"] = o.DustThreshold
	}
	if o.RecommendedMinAmountIn != nil {
		toSerialize["recommended_min_amount_in"] = o.RecommendedMinAmountIn
	}
	if true {
		toSerialize["recommended_gas_rate"] = o.RecommendedGasRate
	}
	if true {
		toSerialize["gas_rate_units"] = o.GasRateUnits
	}
	if true {
		toSerialize["memo"] = o.Memo
	}
	if o.ExpectedAmountOut != nil {
		toSerialize["expected_amount_out"] = o.ExpectedAmountOut
	}
	if true {
		toSerialize["expected_amount_deposit"] = o.ExpectedAmountDeposit
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteSaverDepositResponse struct {
	value *QuoteSaverDepositResponse
	isSet bool
}

func (v NullableQuoteSaverDepositResponse) Get() *QuoteSaverDepositResponse {
	return v.value
}

func (v *NullableQuoteSaverDepositResponse) Set(val *QuoteSaverDepositResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSaverDepositResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSaverDepositResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSaverDepositResponse(val *QuoteSaverDepositResponse) *NullableQuoteSaverDepositResponse {
	return &NullableQuoteSaverDepositResponse{value: val, isSet: true}
}

func (v NullableQuoteSaverDepositResponse) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSaverDepositResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


