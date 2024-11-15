# InboundConfirmationCountedStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountingStartHeight** | Pointer to **int64** | the MAYAChain block height when confirmation counting began | [optional] 
**Chain** | Pointer to **string** | the external source chain for which confirmation counting takes place | [optional] 
**ExternalObservedHeight** | Pointer to **int64** | the block height on the external source chain when the transaction was observed | [optional] 
**ExternalConfirmationDelayHeight** | Pointer to **int64** | the block height on the external source chain when confirmation counting will be complete | [optional] 
**RemainingConfirmationSeconds** | Pointer to **int64** | the estimated remaining seconds before confirmation counting completes | [optional] 
**Completed** | **bool** | returns true if no transaction confirmation counting remains to be done | 

## Methods

### NewInboundConfirmationCountedStage

`func NewInboundConfirmationCountedStage(completed bool, ) *InboundConfirmationCountedStage`

NewInboundConfirmationCountedStage instantiates a new InboundConfirmationCountedStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundConfirmationCountedStageWithDefaults

`func NewInboundConfirmationCountedStageWithDefaults() *InboundConfirmationCountedStage`

NewInboundConfirmationCountedStageWithDefaults instantiates a new InboundConfirmationCountedStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountingStartHeight

`func (o *InboundConfirmationCountedStage) GetCountingStartHeight() int64`

GetCountingStartHeight returns the CountingStartHeight field if non-nil, zero value otherwise.

### GetCountingStartHeightOk

`func (o *InboundConfirmationCountedStage) GetCountingStartHeightOk() (*int64, bool)`

GetCountingStartHeightOk returns a tuple with the CountingStartHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountingStartHeight

`func (o *InboundConfirmationCountedStage) SetCountingStartHeight(v int64)`

SetCountingStartHeight sets CountingStartHeight field to given value.

### HasCountingStartHeight

`func (o *InboundConfirmationCountedStage) HasCountingStartHeight() bool`

HasCountingStartHeight returns a boolean if a field has been set.

### GetChain

`func (o *InboundConfirmationCountedStage) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *InboundConfirmationCountedStage) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *InboundConfirmationCountedStage) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *InboundConfirmationCountedStage) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetExternalObservedHeight

`func (o *InboundConfirmationCountedStage) GetExternalObservedHeight() int64`

GetExternalObservedHeight returns the ExternalObservedHeight field if non-nil, zero value otherwise.

### GetExternalObservedHeightOk

`func (o *InboundConfirmationCountedStage) GetExternalObservedHeightOk() (*int64, bool)`

GetExternalObservedHeightOk returns a tuple with the ExternalObservedHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalObservedHeight

`func (o *InboundConfirmationCountedStage) SetExternalObservedHeight(v int64)`

SetExternalObservedHeight sets ExternalObservedHeight field to given value.

### HasExternalObservedHeight

`func (o *InboundConfirmationCountedStage) HasExternalObservedHeight() bool`

HasExternalObservedHeight returns a boolean if a field has been set.

### GetExternalConfirmationDelayHeight

`func (o *InboundConfirmationCountedStage) GetExternalConfirmationDelayHeight() int64`

GetExternalConfirmationDelayHeight returns the ExternalConfirmationDelayHeight field if non-nil, zero value otherwise.

### GetExternalConfirmationDelayHeightOk

`func (o *InboundConfirmationCountedStage) GetExternalConfirmationDelayHeightOk() (*int64, bool)`

GetExternalConfirmationDelayHeightOk returns a tuple with the ExternalConfirmationDelayHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfirmationDelayHeight

`func (o *InboundConfirmationCountedStage) SetExternalConfirmationDelayHeight(v int64)`

SetExternalConfirmationDelayHeight sets ExternalConfirmationDelayHeight field to given value.

### HasExternalConfirmationDelayHeight

`func (o *InboundConfirmationCountedStage) HasExternalConfirmationDelayHeight() bool`

HasExternalConfirmationDelayHeight returns a boolean if a field has been set.

### GetRemainingConfirmationSeconds

`func (o *InboundConfirmationCountedStage) GetRemainingConfirmationSeconds() int64`

GetRemainingConfirmationSeconds returns the RemainingConfirmationSeconds field if non-nil, zero value otherwise.

### GetRemainingConfirmationSecondsOk

`func (o *InboundConfirmationCountedStage) GetRemainingConfirmationSecondsOk() (*int64, bool)`

GetRemainingConfirmationSecondsOk returns a tuple with the RemainingConfirmationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingConfirmationSeconds

`func (o *InboundConfirmationCountedStage) SetRemainingConfirmationSeconds(v int64)`

SetRemainingConfirmationSeconds sets RemainingConfirmationSeconds field to given value.

### HasRemainingConfirmationSeconds

`func (o *InboundConfirmationCountedStage) HasRemainingConfirmationSeconds() bool`

HasRemainingConfirmationSeconds returns a boolean if a field has been set.

### GetCompleted

`func (o *InboundConfirmationCountedStage) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InboundConfirmationCountedStage) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InboundConfirmationCountedStage) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


