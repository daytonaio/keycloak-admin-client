# MappingsRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealmMappings** | Pointer to [**[]RoleRepresentation**](RoleRepresentation.md) |  | [optional] 
**ClientMappings** | Pointer to [**map[string]ClientMappingsRepresentation**](ClientMappingsRepresentation.md) |  | [optional] 

## Methods

### NewMappingsRepresentation

`func NewMappingsRepresentation() *MappingsRepresentation`

NewMappingsRepresentation instantiates a new MappingsRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingsRepresentationWithDefaults

`func NewMappingsRepresentationWithDefaults() *MappingsRepresentation`

NewMappingsRepresentationWithDefaults instantiates a new MappingsRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealmMappings

`func (o *MappingsRepresentation) GetRealmMappings() []RoleRepresentation`

GetRealmMappings returns the RealmMappings field if non-nil, zero value otherwise.

### GetRealmMappingsOk

`func (o *MappingsRepresentation) GetRealmMappingsOk() (*[]RoleRepresentation, bool)`

GetRealmMappingsOk returns a tuple with the RealmMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmMappings

`func (o *MappingsRepresentation) SetRealmMappings(v []RoleRepresentation)`

SetRealmMappings sets RealmMappings field to given value.

### HasRealmMappings

`func (o *MappingsRepresentation) HasRealmMappings() bool`

HasRealmMappings returns a boolean if a field has been set.

### GetClientMappings

`func (o *MappingsRepresentation) GetClientMappings() map[string]ClientMappingsRepresentation`

GetClientMappings returns the ClientMappings field if non-nil, zero value otherwise.

### GetClientMappingsOk

`func (o *MappingsRepresentation) GetClientMappingsOk() (*map[string]ClientMappingsRepresentation, bool)`

GetClientMappingsOk returns a tuple with the ClientMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMappings

`func (o *MappingsRepresentation) SetClientMappings(v map[string]ClientMappingsRepresentation)`

SetClientMappings sets ClientMappings field to given value.

### HasClientMappings

`func (o *MappingsRepresentation) HasClientMappings() bool`

HasClientMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


