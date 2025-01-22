# AuthenticationFlowRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**TopLevel** | Pointer to **bool** |  | [optional] 
**BuiltIn** | Pointer to **bool** |  | [optional] 
**AuthenticationExecutions** | Pointer to [**[]AuthenticationExecutionExportRepresentation**](AuthenticationExecutionExportRepresentation.md) |  | [optional] 

## Methods

### NewAuthenticationFlowRepresentation

`func NewAuthenticationFlowRepresentation() *AuthenticationFlowRepresentation`

NewAuthenticationFlowRepresentation instantiates a new AuthenticationFlowRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationFlowRepresentationWithDefaults

`func NewAuthenticationFlowRepresentationWithDefaults() *AuthenticationFlowRepresentation`

NewAuthenticationFlowRepresentationWithDefaults instantiates a new AuthenticationFlowRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationFlowRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationFlowRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationFlowRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationFlowRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *AuthenticationFlowRepresentation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AuthenticationFlowRepresentation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AuthenticationFlowRepresentation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AuthenticationFlowRepresentation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationFlowRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationFlowRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationFlowRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationFlowRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProviderId

`func (o *AuthenticationFlowRepresentation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AuthenticationFlowRepresentation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AuthenticationFlowRepresentation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AuthenticationFlowRepresentation) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetTopLevel

`func (o *AuthenticationFlowRepresentation) GetTopLevel() bool`

GetTopLevel returns the TopLevel field if non-nil, zero value otherwise.

### GetTopLevelOk

`func (o *AuthenticationFlowRepresentation) GetTopLevelOk() (*bool, bool)`

GetTopLevelOk returns a tuple with the TopLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevel

`func (o *AuthenticationFlowRepresentation) SetTopLevel(v bool)`

SetTopLevel sets TopLevel field to given value.

### HasTopLevel

`func (o *AuthenticationFlowRepresentation) HasTopLevel() bool`

HasTopLevel returns a boolean if a field has been set.

### GetBuiltIn

`func (o *AuthenticationFlowRepresentation) GetBuiltIn() bool`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *AuthenticationFlowRepresentation) GetBuiltInOk() (*bool, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *AuthenticationFlowRepresentation) SetBuiltIn(v bool)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *AuthenticationFlowRepresentation) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetAuthenticationExecutions

`func (o *AuthenticationFlowRepresentation) GetAuthenticationExecutions() []AuthenticationExecutionExportRepresentation`

GetAuthenticationExecutions returns the AuthenticationExecutions field if non-nil, zero value otherwise.

### GetAuthenticationExecutionsOk

`func (o *AuthenticationFlowRepresentation) GetAuthenticationExecutionsOk() (*[]AuthenticationExecutionExportRepresentation, bool)`

GetAuthenticationExecutionsOk returns a tuple with the AuthenticationExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationExecutions

`func (o *AuthenticationFlowRepresentation) SetAuthenticationExecutions(v []AuthenticationExecutionExportRepresentation)`

SetAuthenticationExecutions sets AuthenticationExecutions field to given value.

### HasAuthenticationExecutions

`func (o *AuthenticationFlowRepresentation) HasAuthenticationExecutions() bool`

HasAuthenticationExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


