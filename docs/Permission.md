# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rsid** | Pointer to **string** |  | [optional] 
**Rsname** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Claims** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRsid

`func (o *Permission) GetRsid() string`

GetRsid returns the Rsid field if non-nil, zero value otherwise.

### GetRsidOk

`func (o *Permission) GetRsidOk() (*string, bool)`

GetRsidOk returns a tuple with the Rsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsid

`func (o *Permission) SetRsid(v string)`

SetRsid sets Rsid field to given value.

### HasRsid

`func (o *Permission) HasRsid() bool`

HasRsid returns a boolean if a field has been set.

### GetRsname

`func (o *Permission) GetRsname() string`

GetRsname returns the Rsname field if non-nil, zero value otherwise.

### GetRsnameOk

`func (o *Permission) GetRsnameOk() (*string, bool)`

GetRsnameOk returns a tuple with the Rsname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsname

`func (o *Permission) SetRsname(v string)`

SetRsname sets Rsname field to given value.

### HasRsname

`func (o *Permission) HasRsname() bool`

HasRsname returns a boolean if a field has been set.

### GetScopes

`func (o *Permission) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Permission) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Permission) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Permission) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetClaims

`func (o *Permission) GetClaims() map[string][]string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *Permission) GetClaimsOk() (*map[string][]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *Permission) SetClaims(v map[string][]string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *Permission) HasClaims() bool`

HasClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


