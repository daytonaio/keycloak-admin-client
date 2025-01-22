# ClientMappingsRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**Mappings** | Pointer to [**[]RoleRepresentation**](RoleRepresentation.md) |  | [optional] 

## Methods

### NewClientMappingsRepresentation

`func NewClientMappingsRepresentation() *ClientMappingsRepresentation`

NewClientMappingsRepresentation instantiates a new ClientMappingsRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMappingsRepresentationWithDefaults

`func NewClientMappingsRepresentationWithDefaults() *ClientMappingsRepresentation`

NewClientMappingsRepresentationWithDefaults instantiates a new ClientMappingsRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientMappingsRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientMappingsRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientMappingsRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientMappingsRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClient

`func (o *ClientMappingsRepresentation) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ClientMappingsRepresentation) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ClientMappingsRepresentation) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ClientMappingsRepresentation) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetMappings

`func (o *ClientMappingsRepresentation) GetMappings() []RoleRepresentation`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ClientMappingsRepresentation) GetMappingsOk() (*[]RoleRepresentation, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ClientMappingsRepresentation) SetMappings(v []RoleRepresentation)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ClientMappingsRepresentation) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


