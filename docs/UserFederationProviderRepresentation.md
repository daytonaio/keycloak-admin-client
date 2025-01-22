# UserFederationProviderRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**FullSyncPeriod** | Pointer to **int32** |  | [optional] 
**ChangedSyncPeriod** | Pointer to **int32** |  | [optional] 
**LastSync** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserFederationProviderRepresentation

`func NewUserFederationProviderRepresentation() *UserFederationProviderRepresentation`

NewUserFederationProviderRepresentation instantiates a new UserFederationProviderRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFederationProviderRepresentationWithDefaults

`func NewUserFederationProviderRepresentationWithDefaults() *UserFederationProviderRepresentation`

NewUserFederationProviderRepresentationWithDefaults instantiates a new UserFederationProviderRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserFederationProviderRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserFederationProviderRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserFederationProviderRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserFederationProviderRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserFederationProviderRepresentation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserFederationProviderRepresentation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserFederationProviderRepresentation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserFederationProviderRepresentation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProviderName

`func (o *UserFederationProviderRepresentation) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *UserFederationProviderRepresentation) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *UserFederationProviderRepresentation) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *UserFederationProviderRepresentation) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetConfig

`func (o *UserFederationProviderRepresentation) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UserFederationProviderRepresentation) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UserFederationProviderRepresentation) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UserFederationProviderRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPriority

`func (o *UserFederationProviderRepresentation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UserFederationProviderRepresentation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UserFederationProviderRepresentation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UserFederationProviderRepresentation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetFullSyncPeriod

`func (o *UserFederationProviderRepresentation) GetFullSyncPeriod() int32`

GetFullSyncPeriod returns the FullSyncPeriod field if non-nil, zero value otherwise.

### GetFullSyncPeriodOk

`func (o *UserFederationProviderRepresentation) GetFullSyncPeriodOk() (*int32, bool)`

GetFullSyncPeriodOk returns a tuple with the FullSyncPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncPeriod

`func (o *UserFederationProviderRepresentation) SetFullSyncPeriod(v int32)`

SetFullSyncPeriod sets FullSyncPeriod field to given value.

### HasFullSyncPeriod

`func (o *UserFederationProviderRepresentation) HasFullSyncPeriod() bool`

HasFullSyncPeriod returns a boolean if a field has been set.

### GetChangedSyncPeriod

`func (o *UserFederationProviderRepresentation) GetChangedSyncPeriod() int32`

GetChangedSyncPeriod returns the ChangedSyncPeriod field if non-nil, zero value otherwise.

### GetChangedSyncPeriodOk

`func (o *UserFederationProviderRepresentation) GetChangedSyncPeriodOk() (*int32, bool)`

GetChangedSyncPeriodOk returns a tuple with the ChangedSyncPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedSyncPeriod

`func (o *UserFederationProviderRepresentation) SetChangedSyncPeriod(v int32)`

SetChangedSyncPeriod sets ChangedSyncPeriod field to given value.

### HasChangedSyncPeriod

`func (o *UserFederationProviderRepresentation) HasChangedSyncPeriod() bool`

HasChangedSyncPeriod returns a boolean if a field has been set.

### GetLastSync

`func (o *UserFederationProviderRepresentation) GetLastSync() int32`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *UserFederationProviderRepresentation) GetLastSyncOk() (*int32, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *UserFederationProviderRepresentation) SetLastSync(v int32)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *UserFederationProviderRepresentation) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


