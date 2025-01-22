# ResourceRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uris** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to [**[]ScopeRepresentation**](ScopeRepresentation.md) |  | [optional] 
**IconUri** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ResourceOwnerRepresentation**](ResourceOwnerRepresentation.md) |  | [optional] [readonly] 
**OwnerManagedAccess** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**ScopesUma** | Pointer to [**[]ScopeRepresentation**](ScopeRepresentation.md) |  | [optional] 

## Methods

### NewResourceRepresentation

`func NewResourceRepresentation() *ResourceRepresentation`

NewResourceRepresentation instantiates a new ResourceRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRepresentationWithDefaults

`func NewResourceRepresentationWithDefaults() *ResourceRepresentation`

NewResourceRepresentationWithDefaults instantiates a new ResourceRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUris

`func (o *ResourceRepresentation) GetUris() []string`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *ResourceRepresentation) GetUrisOk() (*[]string, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *ResourceRepresentation) SetUris(v []string)`

SetUris sets Uris field to given value.

### HasUris

`func (o *ResourceRepresentation) HasUris() bool`

HasUris returns a boolean if a field has been set.

### GetType

`func (o *ResourceRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScopes

`func (o *ResourceRepresentation) GetScopes() []ScopeRepresentation`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ResourceRepresentation) GetScopesOk() (*[]ScopeRepresentation, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ResourceRepresentation) SetScopes(v []ScopeRepresentation)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ResourceRepresentation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetIconUri

`func (o *ResourceRepresentation) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *ResourceRepresentation) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *ResourceRepresentation) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *ResourceRepresentation) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.

### GetOwner

`func (o *ResourceRepresentation) GetOwner() ResourceOwnerRepresentation`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResourceRepresentation) GetOwnerOk() (*ResourceOwnerRepresentation, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResourceRepresentation) SetOwner(v ResourceOwnerRepresentation)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResourceRepresentation) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerManagedAccess

`func (o *ResourceRepresentation) GetOwnerManagedAccess() bool`

GetOwnerManagedAccess returns the OwnerManagedAccess field if non-nil, zero value otherwise.

### GetOwnerManagedAccessOk

`func (o *ResourceRepresentation) GetOwnerManagedAccessOk() (*bool, bool)`

GetOwnerManagedAccessOk returns a tuple with the OwnerManagedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerManagedAccess

`func (o *ResourceRepresentation) SetOwnerManagedAccess(v bool)`

SetOwnerManagedAccess sets OwnerManagedAccess field to given value.

### HasOwnerManagedAccess

`func (o *ResourceRepresentation) HasOwnerManagedAccess() bool`

HasOwnerManagedAccess returns a boolean if a field has been set.

### GetDisplayName

`func (o *ResourceRepresentation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResourceRepresentation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResourceRepresentation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ResourceRepresentation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAttributes

`func (o *ResourceRepresentation) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResourceRepresentation) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResourceRepresentation) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ResourceRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUri

`func (o *ResourceRepresentation) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ResourceRepresentation) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ResourceRepresentation) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ResourceRepresentation) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetScopesUma

`func (o *ResourceRepresentation) GetScopesUma() []ScopeRepresentation`

GetScopesUma returns the ScopesUma field if non-nil, zero value otherwise.

### GetScopesUmaOk

`func (o *ResourceRepresentation) GetScopesUmaOk() (*[]ScopeRepresentation, bool)`

GetScopesUmaOk returns a tuple with the ScopesUma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesUma

`func (o *ResourceRepresentation) SetScopesUma(v []ScopeRepresentation)`

SetScopesUma sets ScopesUma field to given value.

### HasScopesUma

`func (o *ResourceRepresentation) HasScopesUma() bool`

HasScopesUma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


