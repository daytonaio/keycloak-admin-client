# ManagementPermissionReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**ScopePermissions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementPermissionReference

`func NewManagementPermissionReference() *ManagementPermissionReference`

NewManagementPermissionReference instantiates a new ManagementPermissionReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementPermissionReferenceWithDefaults

`func NewManagementPermissionReferenceWithDefaults() *ManagementPermissionReference`

NewManagementPermissionReferenceWithDefaults instantiates a new ManagementPermissionReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ManagementPermissionReference) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManagementPermissionReference) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManagementPermissionReference) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManagementPermissionReference) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetResource

`func (o *ManagementPermissionReference) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ManagementPermissionReference) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ManagementPermissionReference) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ManagementPermissionReference) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetScopePermissions

`func (o *ManagementPermissionReference) GetScopePermissions() map[string]string`

GetScopePermissions returns the ScopePermissions field if non-nil, zero value otherwise.

### GetScopePermissionsOk

`func (o *ManagementPermissionReference) GetScopePermissionsOk() (*map[string]string, bool)`

GetScopePermissionsOk returns a tuple with the ScopePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopePermissions

`func (o *ManagementPermissionReference) SetScopePermissions(v map[string]string)`

SetScopePermissions sets ScopePermissions field to given value.

### HasScopePermissions

`func (o *ManagementPermissionReference) HasScopePermissions() bool`

HasScopePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


