# ResourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResourceType

`func NewResourceType() *ResourceType`

NewResourceType instantiates a new ResourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeWithDefaults

`func NewResourceTypeWithDefaults() *ResourceType`

NewResourceTypeWithDefaults instantiates a new ResourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScopes

`func (o *ResourceType) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ResourceType) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ResourceType) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ResourceType) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


