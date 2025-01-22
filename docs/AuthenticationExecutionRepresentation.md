# AuthenticationExecutionRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorConfig** | Pointer to **string** |  | [optional] 
**Authenticator** | Pointer to **string** |  | [optional] 
**AuthenticatorFlow** | Pointer to **bool** |  | [optional] 
**Requirement** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**AutheticatorFlow** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**ParentFlow** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationExecutionRepresentation

`func NewAuthenticationExecutionRepresentation() *AuthenticationExecutionRepresentation`

NewAuthenticationExecutionRepresentation instantiates a new AuthenticationExecutionRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationExecutionRepresentationWithDefaults

`func NewAuthenticationExecutionRepresentationWithDefaults() *AuthenticationExecutionRepresentation`

NewAuthenticationExecutionRepresentationWithDefaults instantiates a new AuthenticationExecutionRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorConfig

`func (o *AuthenticationExecutionRepresentation) GetAuthenticatorConfig() string`

GetAuthenticatorConfig returns the AuthenticatorConfig field if non-nil, zero value otherwise.

### GetAuthenticatorConfigOk

`func (o *AuthenticationExecutionRepresentation) GetAuthenticatorConfigOk() (*string, bool)`

GetAuthenticatorConfigOk returns a tuple with the AuthenticatorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorConfig

`func (o *AuthenticationExecutionRepresentation) SetAuthenticatorConfig(v string)`

SetAuthenticatorConfig sets AuthenticatorConfig field to given value.

### HasAuthenticatorConfig

`func (o *AuthenticationExecutionRepresentation) HasAuthenticatorConfig() bool`

HasAuthenticatorConfig returns a boolean if a field has been set.

### GetAuthenticator

`func (o *AuthenticationExecutionRepresentation) GetAuthenticator() string`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *AuthenticationExecutionRepresentation) GetAuthenticatorOk() (*string, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *AuthenticationExecutionRepresentation) SetAuthenticator(v string)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *AuthenticationExecutionRepresentation) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetAuthenticatorFlow

`func (o *AuthenticationExecutionRepresentation) GetAuthenticatorFlow() bool`

GetAuthenticatorFlow returns the AuthenticatorFlow field if non-nil, zero value otherwise.

### GetAuthenticatorFlowOk

`func (o *AuthenticationExecutionRepresentation) GetAuthenticatorFlowOk() (*bool, bool)`

GetAuthenticatorFlowOk returns a tuple with the AuthenticatorFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorFlow

`func (o *AuthenticationExecutionRepresentation) SetAuthenticatorFlow(v bool)`

SetAuthenticatorFlow sets AuthenticatorFlow field to given value.

### HasAuthenticatorFlow

`func (o *AuthenticationExecutionRepresentation) HasAuthenticatorFlow() bool`

HasAuthenticatorFlow returns a boolean if a field has been set.

### GetRequirement

`func (o *AuthenticationExecutionRepresentation) GetRequirement() string`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *AuthenticationExecutionRepresentation) GetRequirementOk() (*string, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *AuthenticationExecutionRepresentation) SetRequirement(v string)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *AuthenticationExecutionRepresentation) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.

### GetPriority

`func (o *AuthenticationExecutionRepresentation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AuthenticationExecutionRepresentation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AuthenticationExecutionRepresentation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AuthenticationExecutionRepresentation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAutheticatorFlow

`func (o *AuthenticationExecutionRepresentation) GetAutheticatorFlow() bool`

GetAutheticatorFlow returns the AutheticatorFlow field if non-nil, zero value otherwise.

### GetAutheticatorFlowOk

`func (o *AuthenticationExecutionRepresentation) GetAutheticatorFlowOk() (*bool, bool)`

GetAutheticatorFlowOk returns a tuple with the AutheticatorFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutheticatorFlow

`func (o *AuthenticationExecutionRepresentation) SetAutheticatorFlow(v bool)`

SetAutheticatorFlow sets AutheticatorFlow field to given value.

### HasAutheticatorFlow

`func (o *AuthenticationExecutionRepresentation) HasAutheticatorFlow() bool`

HasAutheticatorFlow returns a boolean if a field has been set.

### GetId

`func (o *AuthenticationExecutionRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationExecutionRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationExecutionRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationExecutionRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlowId

`func (o *AuthenticationExecutionRepresentation) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AuthenticationExecutionRepresentation) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AuthenticationExecutionRepresentation) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AuthenticationExecutionRepresentation) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetParentFlow

`func (o *AuthenticationExecutionRepresentation) GetParentFlow() string`

GetParentFlow returns the ParentFlow field if non-nil, zero value otherwise.

### GetParentFlowOk

`func (o *AuthenticationExecutionRepresentation) GetParentFlowOk() (*string, bool)`

GetParentFlowOk returns a tuple with the ParentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFlow

`func (o *AuthenticationExecutionRepresentation) SetParentFlow(v string)`

SetParentFlow sets ParentFlow field to given value.

### HasParentFlow

`func (o *AuthenticationExecutionRepresentation) HasParentFlow() bool`

HasParentFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


