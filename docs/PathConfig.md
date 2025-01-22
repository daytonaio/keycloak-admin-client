# PathConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to [**[]MethodConfig**](MethodConfig.md) |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**EnforcementMode** | Pointer to [**EnforcementMode**](EnforcementMode.md) |  | [optional] 
**ClaimInformationPoint** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Invalidated** | Pointer to **bool** |  | [optional] 
**StaticPath** | Pointer to **bool** |  | [optional] 
**Static** | Pointer to **bool** |  | [optional] 

## Methods

### NewPathConfig

`func NewPathConfig() *PathConfig`

NewPathConfig instantiates a new PathConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathConfigWithDefaults

`func NewPathConfigWithDefaults() *PathConfig`

NewPathConfigWithDefaults instantiates a new PathConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PathConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PathConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PathConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PathConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PathConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PathConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PathConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PathConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPath

`func (o *PathConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethods

`func (o *PathConfig) GetMethods() []MethodConfig`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *PathConfig) GetMethodsOk() (*[]MethodConfig, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *PathConfig) SetMethods(v []MethodConfig)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *PathConfig) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetScopes

`func (o *PathConfig) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PathConfig) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PathConfig) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PathConfig) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetId

`func (o *PathConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PathConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PathConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PathConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnforcementMode

`func (o *PathConfig) GetEnforcementMode() EnforcementMode`

GetEnforcementMode returns the EnforcementMode field if non-nil, zero value otherwise.

### GetEnforcementModeOk

`func (o *PathConfig) GetEnforcementModeOk() (*EnforcementMode, bool)`

GetEnforcementModeOk returns a tuple with the EnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementMode

`func (o *PathConfig) SetEnforcementMode(v EnforcementMode)`

SetEnforcementMode sets EnforcementMode field to given value.

### HasEnforcementMode

`func (o *PathConfig) HasEnforcementMode() bool`

HasEnforcementMode returns a boolean if a field has been set.

### GetClaimInformationPoint

`func (o *PathConfig) GetClaimInformationPoint() map[string]map[string]interface{}`

GetClaimInformationPoint returns the ClaimInformationPoint field if non-nil, zero value otherwise.

### GetClaimInformationPointOk

`func (o *PathConfig) GetClaimInformationPointOk() (*map[string]map[string]interface{}, bool)`

GetClaimInformationPointOk returns a tuple with the ClaimInformationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimInformationPoint

`func (o *PathConfig) SetClaimInformationPoint(v map[string]map[string]interface{})`

SetClaimInformationPoint sets ClaimInformationPoint field to given value.

### HasClaimInformationPoint

`func (o *PathConfig) HasClaimInformationPoint() bool`

HasClaimInformationPoint returns a boolean if a field has been set.

### GetInvalidated

`func (o *PathConfig) GetInvalidated() bool`

GetInvalidated returns the Invalidated field if non-nil, zero value otherwise.

### GetInvalidatedOk

`func (o *PathConfig) GetInvalidatedOk() (*bool, bool)`

GetInvalidatedOk returns a tuple with the Invalidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidated

`func (o *PathConfig) SetInvalidated(v bool)`

SetInvalidated sets Invalidated field to given value.

### HasInvalidated

`func (o *PathConfig) HasInvalidated() bool`

HasInvalidated returns a boolean if a field has been set.

### GetStaticPath

`func (o *PathConfig) GetStaticPath() bool`

GetStaticPath returns the StaticPath field if non-nil, zero value otherwise.

### GetStaticPathOk

`func (o *PathConfig) GetStaticPathOk() (*bool, bool)`

GetStaticPathOk returns a tuple with the StaticPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticPath

`func (o *PathConfig) SetStaticPath(v bool)`

SetStaticPath sets StaticPath field to given value.

### HasStaticPath

`func (o *PathConfig) HasStaticPath() bool`

HasStaticPath returns a boolean if a field has been set.

### GetStatic

`func (o *PathConfig) GetStatic() bool`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *PathConfig) GetStaticOk() (*bool, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *PathConfig) SetStatic(v bool)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *PathConfig) HasStatic() bool`

HasStatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


