# ComponentTypeRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**HelpText** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]ConfigPropertyRepresentation**](ConfigPropertyRepresentation.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewComponentTypeRepresentation

`func NewComponentTypeRepresentation() *ComponentTypeRepresentation`

NewComponentTypeRepresentation instantiates a new ComponentTypeRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentTypeRepresentationWithDefaults

`func NewComponentTypeRepresentationWithDefaults() *ComponentTypeRepresentation`

NewComponentTypeRepresentationWithDefaults instantiates a new ComponentTypeRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentTypeRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentTypeRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentTypeRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentTypeRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHelpText

`func (o *ComponentTypeRepresentation) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ComponentTypeRepresentation) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ComponentTypeRepresentation) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *ComponentTypeRepresentation) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetProperties

`func (o *ComponentTypeRepresentation) GetProperties() []ConfigPropertyRepresentation`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ComponentTypeRepresentation) GetPropertiesOk() (*[]ConfigPropertyRepresentation, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ComponentTypeRepresentation) SetProperties(v []ConfigPropertyRepresentation)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ComponentTypeRepresentation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetMetadata

`func (o *ComponentTypeRepresentation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComponentTypeRepresentation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComponentTypeRepresentation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ComponentTypeRepresentation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


