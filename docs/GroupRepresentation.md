# GroupRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**SubGroupCount** | Pointer to **int64** |  | [optional] 
**SubGroups** | Pointer to [**[]GroupRepresentation**](GroupRepresentation.md) |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 
**RealmRoles** | Pointer to **[]string** |  | [optional] 
**ClientRoles** | Pointer to **map[string][]string** |  | [optional] 
**Access** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewGroupRepresentation

`func NewGroupRepresentation() *GroupRepresentation`

NewGroupRepresentation instantiates a new GroupRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRepresentationWithDefaults

`func NewGroupRepresentationWithDefaults() *GroupRepresentation`

NewGroupRepresentationWithDefaults instantiates a new GroupRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *GroupRepresentation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GroupRepresentation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GroupRepresentation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GroupRepresentation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetParentId

`func (o *GroupRepresentation) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GroupRepresentation) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GroupRepresentation) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GroupRepresentation) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSubGroupCount

`func (o *GroupRepresentation) GetSubGroupCount() int64`

GetSubGroupCount returns the SubGroupCount field if non-nil, zero value otherwise.

### GetSubGroupCountOk

`func (o *GroupRepresentation) GetSubGroupCountOk() (*int64, bool)`

GetSubGroupCountOk returns a tuple with the SubGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroupCount

`func (o *GroupRepresentation) SetSubGroupCount(v int64)`

SetSubGroupCount sets SubGroupCount field to given value.

### HasSubGroupCount

`func (o *GroupRepresentation) HasSubGroupCount() bool`

HasSubGroupCount returns a boolean if a field has been set.

### GetSubGroups

`func (o *GroupRepresentation) GetSubGroups() []GroupRepresentation`

GetSubGroups returns the SubGroups field if non-nil, zero value otherwise.

### GetSubGroupsOk

`func (o *GroupRepresentation) GetSubGroupsOk() (*[]GroupRepresentation, bool)`

GetSubGroupsOk returns a tuple with the SubGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroups

`func (o *GroupRepresentation) SetSubGroups(v []GroupRepresentation)`

SetSubGroups sets SubGroups field to given value.

### HasSubGroups

`func (o *GroupRepresentation) HasSubGroups() bool`

HasSubGroups returns a boolean if a field has been set.

### GetAttributes

`func (o *GroupRepresentation) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GroupRepresentation) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GroupRepresentation) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GroupRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRealmRoles

`func (o *GroupRepresentation) GetRealmRoles() []string`

GetRealmRoles returns the RealmRoles field if non-nil, zero value otherwise.

### GetRealmRolesOk

`func (o *GroupRepresentation) GetRealmRolesOk() (*[]string, bool)`

GetRealmRolesOk returns a tuple with the RealmRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmRoles

`func (o *GroupRepresentation) SetRealmRoles(v []string)`

SetRealmRoles sets RealmRoles field to given value.

### HasRealmRoles

`func (o *GroupRepresentation) HasRealmRoles() bool`

HasRealmRoles returns a boolean if a field has been set.

### GetClientRoles

`func (o *GroupRepresentation) GetClientRoles() map[string][]string`

GetClientRoles returns the ClientRoles field if non-nil, zero value otherwise.

### GetClientRolesOk

`func (o *GroupRepresentation) GetClientRolesOk() (*map[string][]string, bool)`

GetClientRolesOk returns a tuple with the ClientRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRoles

`func (o *GroupRepresentation) SetClientRoles(v map[string][]string)`

SetClientRoles sets ClientRoles field to given value.

### HasClientRoles

`func (o *GroupRepresentation) HasClientRoles() bool`

HasClientRoles returns a boolean if a field has been set.

### GetAccess

`func (o *GroupRepresentation) GetAccess() map[string]bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GroupRepresentation) GetAccessOk() (*map[string]bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GroupRepresentation) SetAccess(v map[string]bool)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GroupRepresentation) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


