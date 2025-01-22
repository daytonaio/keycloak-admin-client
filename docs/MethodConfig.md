# MethodConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**ScopesEnforcementMode** | Pointer to [**ScopeEnforcementMode**](ScopeEnforcementMode.md) |  | [optional] 

## Methods

### NewMethodConfig

`func NewMethodConfig() *MethodConfig`

NewMethodConfig instantiates a new MethodConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodConfigWithDefaults

`func NewMethodConfigWithDefaults() *MethodConfig`

NewMethodConfigWithDefaults instantiates a new MethodConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *MethodConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MethodConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MethodConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MethodConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetScopes

`func (o *MethodConfig) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *MethodConfig) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *MethodConfig) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *MethodConfig) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetScopesEnforcementMode

`func (o *MethodConfig) GetScopesEnforcementMode() ScopeEnforcementMode`

GetScopesEnforcementMode returns the ScopesEnforcementMode field if non-nil, zero value otherwise.

### GetScopesEnforcementModeOk

`func (o *MethodConfig) GetScopesEnforcementModeOk() (*ScopeEnforcementMode, bool)`

GetScopesEnforcementModeOk returns a tuple with the ScopesEnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesEnforcementMode

`func (o *MethodConfig) SetScopesEnforcementMode(v ScopeEnforcementMode)`

SetScopesEnforcementMode sets ScopesEnforcementMode field to given value.

### HasScopesEnforcementMode

`func (o *MethodConfig) HasScopesEnforcementMode() bool`

HasScopesEnforcementMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


