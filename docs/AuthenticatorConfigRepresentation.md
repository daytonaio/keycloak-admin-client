# AuthenticatorConfigRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAuthenticatorConfigRepresentation

`func NewAuthenticatorConfigRepresentation() *AuthenticatorConfigRepresentation`

NewAuthenticatorConfigRepresentation instantiates a new AuthenticatorConfigRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorConfigRepresentationWithDefaults

`func NewAuthenticatorConfigRepresentationWithDefaults() *AuthenticatorConfigRepresentation`

NewAuthenticatorConfigRepresentationWithDefaults instantiates a new AuthenticatorConfigRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticatorConfigRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticatorConfigRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticatorConfigRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticatorConfigRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *AuthenticatorConfigRepresentation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AuthenticatorConfigRepresentation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AuthenticatorConfigRepresentation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AuthenticatorConfigRepresentation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetConfig

`func (o *AuthenticatorConfigRepresentation) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AuthenticatorConfigRepresentation) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AuthenticatorConfigRepresentation) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AuthenticatorConfigRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


