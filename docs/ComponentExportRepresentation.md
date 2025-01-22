# ComponentExportRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**SubComponents** | Pointer to [**map[string][]ComponentExportRepresentation**](array.md) |  | [optional] 
**Config** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewComponentExportRepresentation

`func NewComponentExportRepresentation() *ComponentExportRepresentation`

NewComponentExportRepresentation instantiates a new ComponentExportRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentExportRepresentationWithDefaults

`func NewComponentExportRepresentationWithDefaults() *ComponentExportRepresentation`

NewComponentExportRepresentationWithDefaults instantiates a new ComponentExportRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentExportRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentExportRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentExportRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentExportRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentExportRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentExportRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentExportRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentExportRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderId

`func (o *ComponentExportRepresentation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ComponentExportRepresentation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ComponentExportRepresentation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ComponentExportRepresentation) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetSubType

`func (o *ComponentExportRepresentation) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ComponentExportRepresentation) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ComponentExportRepresentation) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ComponentExportRepresentation) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubComponents

`func (o *ComponentExportRepresentation) GetSubComponents() map[string][]ComponentExportRepresentation`

GetSubComponents returns the SubComponents field if non-nil, zero value otherwise.

### GetSubComponentsOk

`func (o *ComponentExportRepresentation) GetSubComponentsOk() (*map[string][]ComponentExportRepresentation, bool)`

GetSubComponentsOk returns a tuple with the SubComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubComponents

`func (o *ComponentExportRepresentation) SetSubComponents(v map[string][]ComponentExportRepresentation)`

SetSubComponents sets SubComponents field to given value.

### HasSubComponents

`func (o *ComponentExportRepresentation) HasSubComponents() bool`

HasSubComponents returns a boolean if a field has been set.

### GetConfig

`func (o *ComponentExportRepresentation) GetConfig() map[string][]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ComponentExportRepresentation) GetConfigOk() (*map[string][]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ComponentExportRepresentation) SetConfig(v map[string][]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ComponentExportRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


