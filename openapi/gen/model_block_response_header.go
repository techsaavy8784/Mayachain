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

// BlockResponseHeader struct for BlockResponseHeader
type BlockResponseHeader struct {
	Version BlockResponseHeaderVersion `json:"version"`
	ChainId string `json:"chain_id"`
	Height int64 `json:"height"`
	Time string `json:"time"`
	LastBlockId BlockResponseId `json:"last_block_id"`
	LastCommitHash string `json:"last_commit_hash"`
	DataHash string `json:"data_hash"`
	ValidatorsHash string `json:"validators_hash"`
	NextValidatorsHash string `json:"next_validators_hash"`
	ConsensusHash string `json:"consensus_hash"`
	AppHash string `json:"app_hash"`
	LastResultsHash string `json:"last_results_hash"`
	EvidenceHash string `json:"evidence_hash"`
	ProposerAddress string `json:"proposer_address"`
}

// NewBlockResponseHeader instantiates a new BlockResponseHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockResponseHeader(version BlockResponseHeaderVersion, chainId string, height int64, time string, lastBlockId BlockResponseId, lastCommitHash string, dataHash string, validatorsHash string, nextValidatorsHash string, consensusHash string, appHash string, lastResultsHash string, evidenceHash string, proposerAddress string) *BlockResponseHeader {
	this := BlockResponseHeader{}
	this.Version = version
	this.ChainId = chainId
	this.Height = height
	this.Time = time
	this.LastBlockId = lastBlockId
	this.LastCommitHash = lastCommitHash
	this.DataHash = dataHash
	this.ValidatorsHash = validatorsHash
	this.NextValidatorsHash = nextValidatorsHash
	this.ConsensusHash = consensusHash
	this.AppHash = appHash
	this.LastResultsHash = lastResultsHash
	this.EvidenceHash = evidenceHash
	this.ProposerAddress = proposerAddress
	return &this
}

// NewBlockResponseHeaderWithDefaults instantiates a new BlockResponseHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockResponseHeaderWithDefaults() *BlockResponseHeader {
	this := BlockResponseHeader{}
	return &this
}

// GetVersion returns the Version field value
func (o *BlockResponseHeader) GetVersion() BlockResponseHeaderVersion {
	if o == nil {
		var ret BlockResponseHeaderVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetVersionOk() (*BlockResponseHeaderVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *BlockResponseHeader) SetVersion(v BlockResponseHeaderVersion) {
	o.Version = v
}

// GetChainId returns the ChainId field value
func (o *BlockResponseHeader) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *BlockResponseHeader) SetChainId(v string) {
	o.ChainId = v
}

// GetHeight returns the Height field value
func (o *BlockResponseHeader) GetHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BlockResponseHeader) SetHeight(v int64) {
	o.Height = v
}

// GetTime returns the Time field value
func (o *BlockResponseHeader) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *BlockResponseHeader) SetTime(v string) {
	o.Time = v
}

// GetLastBlockId returns the LastBlockId field value
func (o *BlockResponseHeader) GetLastBlockId() BlockResponseId {
	if o == nil {
		var ret BlockResponseId
		return ret
	}

	return o.LastBlockId
}

// GetLastBlockIdOk returns a tuple with the LastBlockId field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetLastBlockIdOk() (*BlockResponseId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastBlockId, true
}

// SetLastBlockId sets field value
func (o *BlockResponseHeader) SetLastBlockId(v BlockResponseId) {
	o.LastBlockId = v
}

// GetLastCommitHash returns the LastCommitHash field value
func (o *BlockResponseHeader) GetLastCommitHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastCommitHash
}

// GetLastCommitHashOk returns a tuple with the LastCommitHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetLastCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastCommitHash, true
}

// SetLastCommitHash sets field value
func (o *BlockResponseHeader) SetLastCommitHash(v string) {
	o.LastCommitHash = v
}

// GetDataHash returns the DataHash field value
func (o *BlockResponseHeader) GetDataHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataHash
}

// GetDataHashOk returns a tuple with the DataHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetDataHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataHash, true
}

// SetDataHash sets field value
func (o *BlockResponseHeader) SetDataHash(v string) {
	o.DataHash = v
}

// GetValidatorsHash returns the ValidatorsHash field value
func (o *BlockResponseHeader) GetValidatorsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorsHash
}

// GetValidatorsHashOk returns a tuple with the ValidatorsHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetValidatorsHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorsHash, true
}

// SetValidatorsHash sets field value
func (o *BlockResponseHeader) SetValidatorsHash(v string) {
	o.ValidatorsHash = v
}

// GetNextValidatorsHash returns the NextValidatorsHash field value
func (o *BlockResponseHeader) GetNextValidatorsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextValidatorsHash
}

// GetNextValidatorsHashOk returns a tuple with the NextValidatorsHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetNextValidatorsHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextValidatorsHash, true
}

// SetNextValidatorsHash sets field value
func (o *BlockResponseHeader) SetNextValidatorsHash(v string) {
	o.NextValidatorsHash = v
}

// GetConsensusHash returns the ConsensusHash field value
func (o *BlockResponseHeader) GetConsensusHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsensusHash
}

// GetConsensusHashOk returns a tuple with the ConsensusHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetConsensusHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsensusHash, true
}

// SetConsensusHash sets field value
func (o *BlockResponseHeader) SetConsensusHash(v string) {
	o.ConsensusHash = v
}

// GetAppHash returns the AppHash field value
func (o *BlockResponseHeader) GetAppHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppHash
}

// GetAppHashOk returns a tuple with the AppHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetAppHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppHash, true
}

// SetAppHash sets field value
func (o *BlockResponseHeader) SetAppHash(v string) {
	o.AppHash = v
}

// GetLastResultsHash returns the LastResultsHash field value
func (o *BlockResponseHeader) GetLastResultsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastResultsHash
}

// GetLastResultsHashOk returns a tuple with the LastResultsHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetLastResultsHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastResultsHash, true
}

// SetLastResultsHash sets field value
func (o *BlockResponseHeader) SetLastResultsHash(v string) {
	o.LastResultsHash = v
}

// GetEvidenceHash returns the EvidenceHash field value
func (o *BlockResponseHeader) GetEvidenceHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvidenceHash
}

// GetEvidenceHashOk returns a tuple with the EvidenceHash field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetEvidenceHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceHash, true
}

// SetEvidenceHash sets field value
func (o *BlockResponseHeader) SetEvidenceHash(v string) {
	o.EvidenceHash = v
}

// GetProposerAddress returns the ProposerAddress field value
func (o *BlockResponseHeader) GetProposerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProposerAddress
}

// GetProposerAddressOk returns a tuple with the ProposerAddress field value
// and a boolean to check if the value has been set.
func (o *BlockResponseHeader) GetProposerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProposerAddress, true
}

// SetProposerAddress sets field value
func (o *BlockResponseHeader) SetProposerAddress(v string) {
	o.ProposerAddress = v
}

func (o BlockResponseHeader) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["chain_id"] = o.ChainId
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["last_block_id"] = o.LastBlockId
	}
	if true {
		toSerialize["last_commit_hash"] = o.LastCommitHash
	}
	if true {
		toSerialize["data_hash"] = o.DataHash
	}
	if true {
		toSerialize["validators_hash"] = o.ValidatorsHash
	}
	if true {
		toSerialize["next_validators_hash"] = o.NextValidatorsHash
	}
	if true {
		toSerialize["consensus_hash"] = o.ConsensusHash
	}
	if true {
		toSerialize["app_hash"] = o.AppHash
	}
	if true {
		toSerialize["last_results_hash"] = o.LastResultsHash
	}
	if true {
		toSerialize["evidence_hash"] = o.EvidenceHash
	}
	if true {
		toSerialize["proposer_address"] = o.ProposerAddress
	}
	return json.Marshal(toSerialize)
}

type NullableBlockResponseHeader struct {
	value *BlockResponseHeader
	isSet bool
}

func (v NullableBlockResponseHeader) Get() *BlockResponseHeader {
	return v.value
}

func (v *NullableBlockResponseHeader) Set(val *BlockResponseHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockResponseHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockResponseHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockResponseHeader(val *BlockResponseHeader) *NullableBlockResponseHeader {
	return &NullableBlockResponseHeader{value: val, isSet: true}
}

func (v NullableBlockResponseHeader) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockResponseHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


