# OutboundDelayStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainingDelayBlocks** | Pointer to **int64** | the number of remaining MAYAChain blocks the outbound will be delayed | [optional] 
**RemainingDelaySeconds** | Pointer to **int64** | the estimated remaining seconds of the outbound delay before it will be sent | [optional] 
**Completed** | **bool** | returns true if no transaction outbound delay remains | 

## Methods

### NewOutboundDelayStage

`func NewOutboundDelayStage(completed bool, ) *OutboundDelayStage`

NewOutboundDelayStage instantiates a new OutboundDelayStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundDelayStageWithDefaults

`func NewOutboundDelayStageWithDefaults() *OutboundDelayStage`

NewOutboundDelayStageWithDefaults instantiates a new OutboundDelayStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainingDelayBlocks

`func (o *OutboundDelayStage) GetRemainingDelayBlocks() int64`

GetRemainingDelayBlocks returns the RemainingDelayBlocks field if non-nil, zero value otherwise.

### GetRemainingDelayBlocksOk

`func (o *OutboundDelayStage) GetRemainingDelayBlocksOk() (*int64, bool)`

GetRemainingDelayBlocksOk returns a tuple with the RemainingDelayBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDelayBlocks

`func (o *OutboundDelayStage) SetRemainingDelayBlocks(v int64)`

SetRemainingDelayBlocks sets RemainingDelayBlocks field to given value.

### HasRemainingDelayBlocks

`func (o *OutboundDelayStage) HasRemainingDelayBlocks() bool`

HasRemainingDelayBlocks returns a boolean if a field has been set.

### GetRemainingDelaySeconds

`func (o *OutboundDelayStage) GetRemainingDelaySeconds() int64`

GetRemainingDelaySeconds returns the RemainingDelaySeconds field if non-nil, zero value otherwise.

### GetRemainingDelaySecondsOk

`func (o *OutboundDelayStage) GetRemainingDelaySecondsOk() (*int64, bool)`

GetRemainingDelaySecondsOk returns a tuple with the RemainingDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDelaySeconds

`func (o *OutboundDelayStage) SetRemainingDelaySeconds(v int64)`

SetRemainingDelaySeconds sets RemainingDelaySeconds field to given value.

### HasRemainingDelaySeconds

`func (o *OutboundDelayStage) HasRemainingDelaySeconds() bool`

HasRemainingDelaySeconds returns a boolean if a field has been set.

### GetCompleted

`func (o *OutboundDelayStage) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *OutboundDelayStage) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *OutboundDelayStage) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


