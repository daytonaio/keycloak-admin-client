# ClientTypeRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**map[string]PropertyConfig**](PropertyConfig.md) |  | [optional] 

## Methods

### NewClientTypeRepresentation

`func NewClientTypeRepresentation() *ClientTypeRepresentation`

NewClientTypeRepresentation instantiates a new ClientTypeRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTypeRepresentationWithDefaults

`func NewClientTypeRepresentationWithDefaults() *ClientTypeRepresentation`

NewClientTypeRepresentationWithDefaults instantiates a new ClientTypeRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientTypeRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientTypeRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientTypeRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientTypeRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *ClientTypeRepresentation) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClientTypeRepresentation) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClientTypeRepresentation) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClientTypeRepresentation) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetParent

`func (o *ClientTypeRepresentation) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ClientTypeRepresentation) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ClientTypeRepresentation) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ClientTypeRepresentation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetConfig

`func (o *ClientTypeRepresentation) GetConfig() map[string]PropertyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClientTypeRepresentation) GetConfigOk() (*map[string]PropertyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClientTypeRepresentation) SetConfig(v map[string]PropertyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClientTypeRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


