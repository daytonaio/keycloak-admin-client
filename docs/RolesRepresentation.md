# RolesRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | Pointer to [**[]RoleRepresentation**](RoleRepresentation.md) |  | [optional] 
**Client** | Pointer to [**map[string][]RoleRepresentation**](array.md) |  | [optional] 
**Application** | Pointer to [**map[string][]RoleRepresentation**](array.md) |  | [optional] 

## Methods

### NewRolesRepresentation

`func NewRolesRepresentation() *RolesRepresentation`

NewRolesRepresentation instantiates a new RolesRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesRepresentationWithDefaults

`func NewRolesRepresentationWithDefaults() *RolesRepresentation`

NewRolesRepresentationWithDefaults instantiates a new RolesRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *RolesRepresentation) GetRealm() []RoleRepresentation`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *RolesRepresentation) GetRealmOk() (*[]RoleRepresentation, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *RolesRepresentation) SetRealm(v []RoleRepresentation)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *RolesRepresentation) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetClient

`func (o *RolesRepresentation) GetClient() map[string][]RoleRepresentation`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *RolesRepresentation) GetClientOk() (*map[string][]RoleRepresentation, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *RolesRepresentation) SetClient(v map[string][]RoleRepresentation)`

SetClient sets Client field to given value.

### HasClient

`func (o *RolesRepresentation) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetApplication

`func (o *RolesRepresentation) GetApplication() map[string][]RoleRepresentation`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *RolesRepresentation) GetApplicationOk() (*map[string][]RoleRepresentation, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *RolesRepresentation) SetApplication(v map[string][]RoleRepresentation)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *RolesRepresentation) HasApplication() bool`

HasApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


