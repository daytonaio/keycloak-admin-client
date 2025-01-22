# ScopeRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IconUri** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to [**[]PolicyRepresentation**](PolicyRepresentation.md) |  | [optional] 
**Resources** | Pointer to [**[]ResourceRepresentation**](ResourceRepresentation.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewScopeRepresentation

`func NewScopeRepresentation() *ScopeRepresentation`

NewScopeRepresentation instantiates a new ScopeRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeRepresentationWithDefaults

`func NewScopeRepresentationWithDefaults() *ScopeRepresentation`

NewScopeRepresentationWithDefaults instantiates a new ScopeRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScopeRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopeRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopeRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScopeRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScopeRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopeRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopeRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScopeRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIconUri

`func (o *ScopeRepresentation) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *ScopeRepresentation) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *ScopeRepresentation) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *ScopeRepresentation) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.

### GetPolicies

`func (o *ScopeRepresentation) GetPolicies() []PolicyRepresentation`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ScopeRepresentation) GetPoliciesOk() (*[]PolicyRepresentation, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ScopeRepresentation) SetPolicies(v []PolicyRepresentation)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ScopeRepresentation) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetResources

`func (o *ScopeRepresentation) GetResources() []ResourceRepresentation`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScopeRepresentation) GetResourcesOk() (*[]ResourceRepresentation, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ScopeRepresentation) SetResources(v []ResourceRepresentation)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ScopeRepresentation) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDisplayName

`func (o *ScopeRepresentation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScopeRepresentation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScopeRepresentation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScopeRepresentation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


