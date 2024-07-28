# Mayaname

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExpireBlockHeight** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PreferredAsset** | **string** |  | 
**Aliases** | [**[]MayanameAlias**](MayanameAlias.md) |  | 

## Methods

### NewMayaname

`func NewMayaname(preferredAsset string, aliases []MayanameAlias, ) *Mayaname`

NewMayaname instantiates a new Mayaname object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMayanameWithDefaults

`func NewMayanameWithDefaults() *Mayaname`

NewMayanameWithDefaults instantiates a new Mayaname object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Mayaname) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mayaname) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mayaname) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mayaname) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpireBlockHeight

`func (o *Mayaname) GetExpireBlockHeight() int64`

GetExpireBlockHeight returns the ExpireBlockHeight field if non-nil, zero value otherwise.

### GetExpireBlockHeightOk

`func (o *Mayaname) GetExpireBlockHeightOk() (*int64, bool)`

GetExpireBlockHeightOk returns a tuple with the ExpireBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireBlockHeight

`func (o *Mayaname) SetExpireBlockHeight(v int64)`

SetExpireBlockHeight sets ExpireBlockHeight field to given value.

### HasExpireBlockHeight

`func (o *Mayaname) HasExpireBlockHeight() bool`

HasExpireBlockHeight returns a boolean if a field has been set.

### GetOwner

`func (o *Mayaname) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Mayaname) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Mayaname) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Mayaname) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPreferredAsset

`func (o *Mayaname) GetPreferredAsset() string`

GetPreferredAsset returns the PreferredAsset field if non-nil, zero value otherwise.

### GetPreferredAssetOk

`func (o *Mayaname) GetPreferredAssetOk() (*string, bool)`

GetPreferredAssetOk returns a tuple with the PreferredAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredAsset

`func (o *Mayaname) SetPreferredAsset(v string)`

SetPreferredAsset sets PreferredAsset field to given value.


### GetAliases

`func (o *Mayaname) GetAliases() []MayanameAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Mayaname) GetAliasesOk() (*[]MayanameAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Mayaname) SetAliases(v []MayanameAlias)`

SetAliases sets Aliases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


