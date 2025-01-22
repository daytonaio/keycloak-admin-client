# ClientScopeRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**ProtocolMappers** | Pointer to [**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) |  | [optional] 

## Methods

### NewClientScopeRepresentation

`func NewClientScopeRepresentation() *ClientScopeRepresentation`

NewClientScopeRepresentation instantiates a new ClientScopeRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientScopeRepresentationWithDefaults

`func NewClientScopeRepresentationWithDefaults() *ClientScopeRepresentation`

NewClientScopeRepresentationWithDefaults instantiates a new ClientScopeRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientScopeRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientScopeRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientScopeRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientScopeRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClientScopeRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientScopeRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientScopeRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientScopeRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClientScopeRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientScopeRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientScopeRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientScopeRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocol

`func (o *ClientScopeRepresentation) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ClientScopeRepresentation) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ClientScopeRepresentation) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ClientScopeRepresentation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAttributes

`func (o *ClientScopeRepresentation) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ClientScopeRepresentation) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ClientScopeRepresentation) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ClientScopeRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProtocolMappers

`func (o *ClientScopeRepresentation) GetProtocolMappers() []ProtocolMapperRepresentation`

GetProtocolMappers returns the ProtocolMappers field if non-nil, zero value otherwise.

### GetProtocolMappersOk

`func (o *ClientScopeRepresentation) GetProtocolMappersOk() (*[]ProtocolMapperRepresentation, bool)`

GetProtocolMappersOk returns a tuple with the ProtocolMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMappers

`func (o *ClientScopeRepresentation) SetProtocolMappers(v []ProtocolMapperRepresentation)`

SetProtocolMappers sets ProtocolMappers field to given value.

### HasProtocolMappers

`func (o *ClientScopeRepresentation) HasProtocolMappers() bool`

HasProtocolMappers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


