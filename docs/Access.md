# Access

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** |  | [optional] 
**VerifyCaller** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccess

`func NewAccess() *Access`

NewAccess instantiates a new Access object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessWithDefaults

`func NewAccessWithDefaults() *Access`

NewAccessWithDefaults instantiates a new Access object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *Access) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Access) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Access) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Access) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetVerifyCaller

`func (o *Access) GetVerifyCaller() bool`

GetVerifyCaller returns the VerifyCaller field if non-nil, zero value otherwise.

### GetVerifyCallerOk

`func (o *Access) GetVerifyCallerOk() (*bool, bool)`

GetVerifyCallerOk returns a tuple with the VerifyCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCaller

`func (o *Access) SetVerifyCaller(v bool)`

SetVerifyCaller sets VerifyCaller field to given value.

### HasVerifyCaller

`func (o *Access) HasVerifyCaller() bool`

HasVerifyCaller returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


